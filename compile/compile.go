package compile

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/consts"
	"github.com/ycl2018/gs/gen"
	"github.com/ycl2018/gs/vm"
)

var _ gen.GsVisitor = (*StackCompileVisitor)(nil)

type StackCompileVisitor struct {
	gen.BaseGsVisitor
	Log         InterpreterListener
	Scopes      map[antlr.ParserRuleContext]Scope // 作用域树&符号表存储
	GlobalScope *GlobalScope
	AllFuncs    []*FunctionSymbol // 所有函数,不包含主函数
	MainFunc    *FunctionSymbol
	CurFunc     *FunctionSymbol
	ToBeFilled  map[*vm.StackInstr][]*vm.StackInstr // 目标分支跳转Tag -> 需要回填的分支跳转指令列表
	TagAlloc    int
	LoopStack   []*ForLoop
}

func NewStackCompileVisitor(scopes map[antlr.ParserRuleContext]Scope, globalScope *GlobalScope, log InterpreterListener) *StackCompileVisitor {
	mainFunc := NewFunctionSymbol("main", nil)
	s := &StackCompileVisitor{
		Log:         log,
		Scopes:      scopes,
		GlobalScope: globalScope,
		ToBeFilled:  map[*vm.StackInstr][]*vm.StackInstr{},
		MainFunc:    mainFunc,
		CurFunc:     mainFunc,
	}
	s.BaseGsVisitor = gen.BaseGsVisitor{ParseTreeVisitor: gen.NewBaseVisitor(s)}
	return s
}

func (s *StackCompileVisitor) WriteInstr(instr *vm.StackInstr) {
	s.CurFunc.Code = append(s.CurFunc.Code, instr)
}

func (s *StackCompileVisitor) AllocTag() int {
	s.TagAlloc--
	return int(s.TagAlloc)
}

func (s *StackCompileVisitor) VisitProgram(ctx *gen.ProgramContext) interface{} {
	s.VisitChildren(ctx)
	// 补充主函数Halt指令IR
	if len(s.MainFunc.Code) == 0 || s.MainFunc.Code[len(s.MainFunc.Code)-1].OpCode != vm.InstrHalt {
		s.MainFunc.Code = append(s.MainFunc.Code, &vm.StackInstr{
			OpCode: vm.InstrHalt,
		})
	}
	return nil
}

func (s *StackCompileVisitor) VisitFunctionDefinition(ctx *gen.FunctionDefinitionContext) interface{} {
	// 'def' ID '(' (vardef (',' vardef)* )? ')' slist
	// 生成函数定义 IR
	funcName := ctx.ID(0).GetText()
	scope := s.Scopes[ctx]
	funcSymbol := scope.Resolve(funcName).(*FunctionSymbol)
	s.CurFunc = funcSymbol
	// 参数列表入栈由解释器负责自动执行
	// 生成函数体IR
	ctx.Block().Accept(s)
	// 补充返回指令IR
	if len(s.CurFunc.Code) == 0 || s.CurFunc.Code[len(s.CurFunc.Code)-1].OpCode != vm.InstrReturn {
		// 函数体没有显式返回值时，补充返回nil指令IR
		s.Write(vm.InstrNil)
		s.Write(vm.InstrReturn)
	}
	s.AllFuncs = append(s.AllFuncs, funcSymbol)
	s.CurFunc = s.MainFunc
	return nil
}

func defineStringConst(val string, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%s", vm.ConstString, val),
		Kind:  vm.ConstString,
		Value: val,
	}
	cSymbol, _ := scope.DefineOrGetConst(constSymbol)
	return cSymbol
}

func (s *StackCompileVisitor) VisitCall(ctx *gen.CallContext) interface{} {
	for _, context := range ctx.AllExpr() {
		context.Accept(s)
	}
	funcName := ctx.ID().GetText()
	if s.Scopes[ctx].Resolve(funcName) == nil {
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("undefined func: %s", funcName))
		return nil
	}
	fSymbol := s.Scopes[ctx].Resolve(funcName).(*FunctionSymbol)
	// 生成调用指令IR
	callInstr := &vm.StackInstr{
		OpCode:   vm.InstrCall,
		Operands: []int{fSymbol.Address}, // 函数ID填充，编译为二进制时会替换为函数地址
	}
	s.WriteInstr(callInstr)
	return nil
}

func (s *StackCompileVisitor) VisitIntAtom(ctx *gen.IntAtomContext) interface{} {
	valStr := ctx.INT().GetText()
	ival, _ := strconv.ParseInt(valStr, 0, 64)
	s.Write(vm.InstrIConst, int(ival))
	return nil
}

func (s *StackCompileVisitor) VisitFloatAtom(ctx *gen.FloatAtomContext) interface{} {
	valStr := ctx.FLOAT().GetText()
	fval, _ := strconv.ParseFloat(valStr, 64)
	// 浮点数常量地址
	symbol := getFloatConst(fval, s.GlobalScope)
	s.WriteInstr(&vm.StackInstr{
		OpCode:   vm.InstrFConst,
		Operands: []int{symbol.GetAddress()}, // 浮点数常量地址

	})
	return nil
}

func getFloatConst(fval float64, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s_%f", vm.ConstFloat64, fval),
		Value: fval,
		Kind:  vm.ConstFloat64,
	}
	symbol, _ := scope.DefineOrGetConst(constSymbol)
	return symbol
}

func getSliceConst(sliceInit *consts.SliceLiteralConst, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s_%s", vm.ConstSliceInit, sliceInit.Name),
		Value: sliceInit,
		Kind:  vm.ConstSliceInit,
	}
	symbol, _ := scope.DefineOrGetConst(constSymbol)
	return symbol
}

func getMapConst(mapInit *consts.MapLiteralConst, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s_%s", vm.ConstMapInit, mapInit.Name),
		Value: mapInit,
		Kind:  vm.ConstMapInit,
	}
	symbol, _ := scope.DefineOrGetConst(constSymbol)
	return symbol
}

func (s *StackCompileVisitor) VisitStringAtom(ctx *gen.StringAtomContext) interface{} {
	valStr := ctx.STRING().GetText()
	// 字符串常量地址
	symbol := defineStringConst(valStr[1:len(valStr)-1], s.GlobalScope)
	s.WriteInstr(&vm.StackInstr{
		OpCode:   vm.InstrSConst,
		Operands: []int{symbol.GetAddress()}, // 字符串常量地址
	})
	return nil
}

func (s *StackCompileVisitor) VisitInstance(ctx *gen.InstanceContext) interface{} {
	// instance : 'new' ID '{' (ID ':' expr (',' ID ':' expr)* ','?)? '}' ;
	structName := ctx.ID(0).GetText()
	structSymbol := s.Scopes[ctx].Resolve(structName)
	if structSymbol == nil {
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("undefined struct: %s", structName))
		return nil
	}
	s.Write(vm.InstrStruct, structSymbol.GetAddress())
	// 初始化结构体字段
	allIDs := ctx.AllID()[1:]
	allExprs := ctx.AllExpr()
	for i := 0; i < len(allIDs); i++ {
		fieldName := allIDs[i].GetText()
		allExprs[i].Accept(s)
		s.Write(vm.InstrFStore, s.defineStringConst(fieldName))
	}
	return nil
}

func (s *StackCompileVisitor) VisitQid(ctx *gen.QidContext) interface{} {
	s.loadQid(ctx)
	return nil
}

func (s *StackCompileVisitor) VisitAddOp(ctx *gen.AddOpContext) interface{} {
	// addOp	: ADD | SUB ;
	op := ctx.GetText()
	if op == "+" {
		s.Write(vm.InstrAdd)
	} else if op == "-" {
		s.Write(vm.InstrSub)
	}
	return nil
}

func (s *StackCompileVisitor) VisitCompOp(ctx *gen.CompOpContext) interface{} {
	// compOp	: EQ | LT | GT | GEQ | LEQ | NEQ ;
	op := ctx.GetText()
	switch op {
	case "==":
		s.Write(vm.InstrEQ)
	case "<":
		s.Write(vm.InstrLT)
	case ">":
		s.Write(vm.InstrGT)
	case ">=":
		s.Write(vm.InstrGEQ)
	case "<=":
		s.Write(vm.InstrLEQ)
	case "!=":
		s.Write(vm.InstrNEQ)
	default:
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("unknown compOp %s", op))
	}
	return nil
}

func (s *StackCompileVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(s)
}

func (s *StackCompileVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	ctx := node.(antlr.ParserRuleContext)
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(s)
	}
	return nil
}

func (s *StackCompileVisitor) VisitAssign(ctx *gen.AssignContext) interface{} {
	var qids []*gen.QidContext
	var exprs []gen.IExprContext
	for _, child := range ctx.GetChildren() {
		if qid, ok := child.(*gen.QidContext); ok {
			qids = append(qids, qid)
		} else if e, ok := child.(*gen.ExprContext); ok {
			exprs = append(exprs, e)
		}
	}
	// exprs  from left to right
	for _, expr := range exprs {
		expr.Accept(s)
	}
	// build tuple
	if len(exprs) > 1 {
		s.Write(vm.InstrBuildTuple, len(exprs))
	}
	// unpack tuple
	if len(qids) > 1 {
		s.Write(vm.InstrUnpack, len(qids))
	}
	// qids from left to right
	//  qid (',' qid)* assignOp expr (',' expr)*
	// qid : primary ( (DOT | SAFE_DOT ) ID | (LBRACK | SAFE_LBRACK) expr ']' )* ;
	for i := len(qids) - 1; i >= 0; i-- {
		qid := ctx.Qid(i)
		s.storeQid(qid)
	}
	return nil
}

func (s *StackCompileVisitor) VisitReturnStmt(ctx *gen.ReturnStmtContext) interface{} {
	es := ctx.AllExpr()
	// exprs  from left to right
	for _, expr := range es {
		expr.Accept(s)
	}
	lenExprs := len(es)
	if lenExprs > 1 {
		s.Write(vm.InstrBuildTuple, lenExprs)
	}
	if lenExprs == 0 {
		s.Write(vm.InstrNil)
	}
	s.Write(vm.InstrReturn)
	return nil
}

func (s *StackCompileVisitor) VisitPrintStmt(ctx *gen.PrintStmtContext) interface{} {
	exprs := ctx.AllExpr()
	for _, expr := range exprs {
		expr.Accept(s)
	}
	s.Write(vm.InstrPrint, len(exprs))
	return nil
}

func (s *StackCompileVisitor) VisitIfStmt(ctx *gen.IfStmtContext) interface{} {
	// 生成分支语句IR
	// 'if' (assign ';')? expr block ('else' block)?
	if ctx.Assign() != nil {
		ctx.Assign().Accept(s)
	}
	blocks := ctx.AllBlock()
	if len(blocks) == 0 {
		return nil
	}
	ctx.Expr().Accept(s)
	// 生成分支跳转IR
	brfInstr := vm.NewStackInstr(vm.InstrBRF, placeholder)
	s.WriteInstr(brfInstr)
	blocks[0].Accept(s)
	if len(blocks) == 2 {
		var brInstr = vm.NewStackInstr(vm.InstrBR, placeholder)
		s.WriteInstr(brInstr)
		s.FillTarget(brfInstr)
		blocks[1].Accept(s)
		s.FillTarget(brInstr)
	} else {
		s.FillTarget(brfInstr)
	}

	return nil
}

func (s *StackCompileVisitor) VisitForCStyleStmt(ctx *gen.ForCStyleStmtContext) interface{} {
	// 'for' forInit? ';' expr? ';' forUpdate? block
	s.PushLoopStack(&ForLoop{Token: ctx.GetStart()})
	forInit, expr, forUpdate := ctx.ForInit(), ctx.Expr(), ctx.ForUpdate()
	if forInit != nil {
		forInit.Accept(s)
	}
	brTarget := len(s.CurFunc.Code)
	if expr != nil {
		expr.Accept(s)
	}
	brfInstr := vm.NewStackInstr(vm.InstrBRF, placeholder)
	s.WriteInstr(brfInstr)
	ctx.Block().Accept(s)
	updateTarget := len(s.CurFunc.Code)
	if forUpdate != nil {
		forUpdate.Accept(s)
	}
	s.Write(vm.InstrBR, brTarget)
	s.FillTarget(brfInstr)
	loop := s.PopLoopStack()
	for _, br := range loop.Breaks {
		s.FillTarget(br)
	}
	for _, br := range loop.Continues {
		br.Operands[0] = updateTarget
	}
	return nil
}

func (s *StackCompileVisitor) VisitForRangeStmt(ctx *gen.ForRangeStmtContext) interface{} {
	// 'for' iterVar '=' 'range' expr block
	s.PushLoopStack(&ForLoop{Token: ctx.GetStart()})
	ctx.Expr().Accept(s)
	brNilInstr := vm.NewStackInstr(vm.InstrBRNil, placeholder)
	s.WriteInstr(brNilInstr) // for safe range
	s.Write(vm.InstrIter)    // push iter state
	rangeNextAddr := len(s.CurFunc.Code)
	s.Write(vm.InstrIterDone)
	brtInstr := vm.NewStackInstr(vm.InstrBRT, placeholder)
	s.WriteInstr(brtInstr)
	scope := s.Scopes[ctx]
	switch iterVar := ctx.IterVar().(type) {
	case *gen.SingleIterContext:
		s.Write(vm.InstrIterNext, 1)
		iter := scope.Resolve(iterVar.ID().GetText()).(*VariableSymbol)
		s.EmitStore(iter)
	case *gen.DoubleIterContext:
		s.Write(vm.InstrIterNext, 2) //push k,v on stack
		second := scope.Resolve(iterVar.ID(0).GetText()).(*VariableSymbol)
		first := scope.Resolve(iterVar.ID(1).GetText()).(*VariableSymbol)
		s.EmitStore(second)
		s.EmitStore(first)
	}
	ctx.Block().Accept(s)
	s.Write(vm.InstrBR, rangeNextAddr)
	s.FillTarget(brNilInstr, brtInstr)

	loop := s.PopLoopStack()
	for _, br := range loop.Breaks {
		s.FillTarget(br)
	}
	for _, br := range loop.Continues {
		br.Operands[0] = rangeNextAddr
	}
	s.Write(vm.InstrPop, 1) // pop iter state
	return nil
}

func (s *StackCompileVisitor) VisitForCondStmt(ctx *gen.ForCondStmtContext) interface{} {
	// 'for' expr? block
	s.PushLoopStack(&ForLoop{Token: ctx.GetStart()})
	brTarget := len(s.CurFunc.Code)
	if ctx.Expr() != nil {
		ctx.Expr().Accept(s)
	}
	brfInstr := vm.NewStackInstr(vm.InstrBRF, placeholder)
	s.WriteInstr(brfInstr)
	ctx.Block().Accept(s)
	s.Write(vm.InstrBR, brTarget)
	s.FillTarget(brfInstr)
	loop := s.PopLoopStack()
	for _, br := range loop.Breaks {
		s.FillTarget(br)
	}
	for _, br := range loop.Continues {
		br.Operands[0] = brTarget
	}
	return nil
}

func (s *StackCompileVisitor) VisitCallStmt(ctx *gen.CallStmtContext) interface{} {
	ctx.Call().Accept(s)
	// discard return values
	s.Write(vm.InstrPop, 1)
	return nil
}

func (s *StackCompileVisitor) VisitBreakStmt(ctx *gen.BreakStmtContext) interface{} {
	breakInstr := vm.NewStackInstr(vm.InstrBR, placeholder)
	s.WriteInstr(breakInstr)
	if s.GetCurrentLoop() == nil {
		s.Log.ErrorToken(ctx.GetStart(), "break statement not in loop")
		return nil
	}
	s.GetCurrentLoop().Breaks = append(s.GetCurrentLoop().Breaks, breakInstr)
	return nil
}

func (s *StackCompileVisitor) VisitContinueStmt(ctx *gen.ContinueStmtContext) interface{} {
	continueInstr := vm.NewStackInstr(vm.InstrBR, placeholder)
	s.WriteInstr(continueInstr)
	if s.GetCurrentLoop() == nil {
		s.Log.ErrorToken(ctx.GetStart(), "continue statement not in loop")
		return nil
	}
	s.GetCurrentLoop().Continues = append(s.GetCurrentLoop().Continues, continueInstr)
	return nil
}

func (s *StackCompileVisitor) VisitIncrDecr(ctx *gen.IncrDecrContext) interface{} {
	// qid (INCR | DECR)
	qid := ctx.Qid()
	incr := ctx.INCR() != nil
	if incr {
		s.Write(vm.InstrIConst, 1)
	} else {
		s.Write(vm.InstrIConst, -1)
	}
	s.loadQid(qid)
	s.Write(vm.InstrAdd)
	s.storeQid(qid)
	return nil
}

func (s *StackCompileVisitor) VisitSelfAssign(ctx *gen.SelfAssignContext) interface{} {
	// qid selfAssignOp expr
	qid := ctx.Qid()
	ctx.Expr().Accept(s)
	s.loadQid(qid)
	ctx.SelfAssignOp().Accept(s)
	s.storeQid(qid)
	return nil
}

func (s *StackCompileVisitor) VisitSelfAssignOp(ctx *gen.SelfAssignOpContext) interface{} {
	// selfAssignOp
	// '+=' | '-=' | '*=' | '/=' | '%='
	switch ctx.GetText() {
	case "+=":
		s.Write(vm.InstrAdd)
	case "-=":
		s.Write(vm.InstrSub)
	case "*=":
		s.Write(vm.InstrMul)
	case "/=":
		s.Write(vm.InstrDiv)
	case "%=":
		s.Write(vm.InstrMod)
	default:
		s.Log.ErrorToken(ctx.GetStart(), "unknown selfAssignOp: %s", ctx.GetText())
	}
	return nil
}

func (s *StackCompileVisitor) VisitLogicalOrExpr(ctx *gen.LogicalOrExprContext) interface{} {
	allLogicalAndExpr := ctx.AllLogicalAndExpr()
	if len(allLogicalAndExpr) == 0 && ctx.GetChildCount() == 1 {
		// must be constNode
		return s.VisitChildren(ctx)
	}
	var brts []*vm.StackInstr
	for i, andExpr := range allLogicalAndExpr {
		andExpr.Accept(s)
		if i != 0 {
			s.Write(vm.InstrOR)
		}
		if i < len(allLogicalAndExpr)-1 {
			brt := vm.NewStackInstr(vm.InstrBRT, placeholder)
			s.WriteInstr(brt)
			brts = append(brts, brt)
		}
	}
	s.FillTarget(brts...)
	return nil
}

func (s *StackCompileVisitor) VisitLogicalAndExpr(ctx *gen.LogicalAndExprContext) interface{} {
	allComparisonExpr := ctx.AllComparisonExpr()
	brfs := make([]*vm.StackInstr, 0, len(allComparisonExpr)-1)
	for i, cmpExpr := range allComparisonExpr {
		cmpExpr.Accept(s)
		if i != 0 {
			s.Write(vm.InstrAND)
		}
		if i < len(allComparisonExpr)-1 {
			brf := vm.NewStackInstr(vm.InstrBRF, placeholder)
			s.WriteInstr(brf)
			brfs = append(brfs, brf)
		}
	}
	s.FillTarget(brfs...)
	return nil
}

func (s *StackCompileVisitor) VisitComparisonExpr(ctx *gen.ComparisonExprContext) interface{} {
	allExprs, cmpOp := ctx.AllAddExpr(), ctx.CompOp()
	for i := range allExprs {
		allExprs[i].Accept(s)
	}
	if cmpOp != nil {
		cmpOp.Accept(s)
	}
	return nil
}

func (s *StackCompileVisitor) VisitAddExpr(ctx *gen.AddExprContext) interface{} {
	var preOp *gen.AddOpContext
	for _, tree := range ctx.GetChildren() {
		switch tree := tree.(type) {
		default:
			tree.(antlr.RuleContext).Accept(s)
			if preOp != nil {
				preOp.Accept(s)
			}
		case *gen.AddOpContext:
			preOp = tree
		}
	}
	return nil
}

func (s *StackCompileVisitor) VisitBinExpr(ctx *gen.BinExprContext) interface{} {
	var preOp *gen.BitOpContext
	for _, tree := range ctx.GetChildren() {
		switch tree := tree.(type) {
		case *gen.MulExprContext:
			tree.Accept(s)
			if preOp != nil {
				preOp.Accept(s)
			}
		case *gen.BitOpContext:
			preOp = tree
		}
	}
	return nil
}

func (s *StackCompileVisitor) VisitMulExpr(ctx *gen.MulExprContext) interface{} {
	var preOp *gen.MulOpContext
	for _, tree := range ctx.GetChildren() {
		switch tree := tree.(type) {
		case *gen.PowExprContext:
			tree.Accept(s)
			if preOp != nil {
				preOp.Accept(s)
			}
		case *gen.MulOpContext:
			preOp = tree
		}
	}
	return nil
}

func (s *StackCompileVisitor) VisitPowExpr(ctx *gen.PowExprContext) interface{} {
	for i := range ctx.AllAtom() {
		ctx.Atom(i).Accept(s)
		if i != 0 {
			s.Write(vm.InstrPow)
		}
	}
	return nil
}

func (s *StackCompileVisitor) VisitNegAtom(ctx *gen.NegAtomContext) interface{} {
	ctx.Atom().Accept(s)
	s.Write(vm.InstrNeg)
	return nil
}

func (s *StackCompileVisitor) VisitTrueAtom(ctx *gen.TrueAtomContext) interface{} {
	s.Write(vm.InstrTrue)
	return nil
}

func (s *StackCompileVisitor) VisitFalseAtom(ctx *gen.FalseAtomContext) interface{} {
	s.Write(vm.InstrFalse)
	return nil
}

func (s *StackCompileVisitor) VisitNotAtom(ctx *gen.NotAtomContext) interface{} {
	ctx.Expr().Accept(s)
	s.Write(vm.InstrNot)
	return nil
}

func (s *StackCompileVisitor) VisitQidAtom(ctx *gen.QidAtomContext) interface{} {
	s.loadQid(ctx.Qid())
	return nil
}

func (s *StackCompileVisitor) VisitArrayLiteral(ctx *gen.ArrayLiteralContext) interface{} {
	// arrayLiteral : '[' (expr (',' expr)* ','?)? ']' ;
	if len(ctx.GetChildren()) == 1 {
		// must be constNode
		return s.VisitChildren(ctx)
	}
	exprs := ctx.AllExpr()
	for i := range exprs {
		ctx.Expr(i).Accept(s)
	}
	s.Write(vm.InstrArray, len(exprs))
	return nil
}

func (s *StackCompileVisitor) VisitIndexAccess(ctx *gen.IndexAccessContext) interface{} {
	// indexAccess : qid '[' (expr | sliceExpr) ']' ;
	s.loadQid(ctx.Qid())
	expr := ctx.Expr()
	if expr != nil {
		expr.Accept(s)
		s.Write(vm.InstrIndexAccess)
	} else {
		ctx.SliceExpr().Accept(s)
	}

	return nil
}

func (s *StackCompileVisitor) VisitSliceExpr(ctx *gen.SliceExprContext) interface{} {
	// sliceExpr : start=expr? ':' end=expr? ;
	var isLeft = true
	var l, r *gen.ExprContext
	for _, tree := range ctx.GetChildren() {
		switch tree := tree.(type) {
		case *gen.ExprContext:
			if isLeft {
				l = tree
			} else {
				r = tree
			}
		case antlr.TerminalNode:
			if tree.GetSymbol().GetTokenType() == gen.GsParserCOLON {
				isLeft = false
			}
		}
	}
	if l != nil {
		l.Accept(s)
	} else {
		s.Write(vm.InstrIConst, 0)
	}
	if r != nil {
		r.Accept(s)
	} else {
		s.Write(vm.InstrIConst, 0)
	}
	// qidValue | l | r | InstrSlice
	s.Write(vm.InstrSliceSplit)
	return nil
}

func (s *StackCompileVisitor) VisitDictLiteral(ctx *gen.DictLiteralContext) interface{} {
	// dictLiteral : '{' (dictEntry (',' dictEntry)* ','?)? '}' ;
	if len(ctx.GetChildren()) == 1 {
		// must be constNode
		return s.VisitChildren(ctx)
	}
	entries := ctx.AllDictEntry()
	for i := range entries {
		entries[i].Accept(s)
		// pack tuple2 (key, value)
		s.Write(vm.InstrBuildTuple, 2)
	}
	s.Write(vm.InstrDict, len(entries))
	return nil
}

func (s *StackCompileVisitor) VisitConstKeyEntry(ctx *gen.ConstKeyEntryContext) interface{} {
	// strKeyEntry : STRING ':' expr ;
	key := ctx.GetChildren()[0].(antlr.TerminalNode)
	switch key.GetSymbol().GetTokenType() {
	case gen.GsLexerSTRING:
		s.Write(vm.InstrSConst, s.defineStringConst(ctx.STRING().GetText()[1:len(ctx.STRING().GetText())-1]))
	case gen.GsLexerINT:
		val, err := strconv.ParseInt(key.GetText(), 0, 64)
		if err != nil {
			panic(fmt.Sprintf("can't parse %s to int", key.GetText()))
		}
		s.Write(vm.InstrIConst, int(val))
	case gen.GsLexerFLOAT:
		val, err := strconv.ParseFloat(key.GetText(), 64)
		if err != nil {
			panic(fmt.Sprintf("can't parse %s to float", key.GetText()))
		}
		s.Write(vm.InstrFConst, getFloatConst(val, s.GlobalScope).GetAddress())
	case gen.GsParserTRUE:
		s.Write(vm.InstrTrue)
	case gen.GsParserFALSE:
		s.Write(vm.InstrFalse)
	default:
		panic(fmt.Sprintf("unknown tokenType:%s", key.GetText()))
	}

	ctx.Expr().Accept(s)
	return nil
}

func (s *StackCompileVisitor) VisitIdKeyEntry(ctx *gen.IdKeyEntryContext) interface{} {
	// idKeyEntry : ID ':' expr ;
	s.loadQid(ctx.Qid())
	ctx.Expr().Accept(s)
	return nil
}

func (s *StackCompileVisitor) VisitBitOp(ctx *gen.BitOpContext) interface{} {
	// bitOp : BITAND | BITOR | XOR ;
	switch ctx.GetText() {
	case "&":
		s.Write(vm.InstrBitAND)
	case "|":
		s.Write(vm.InstrBitOR)
	case "^":
		s.Write(vm.InstrXOR)
	default:
		panic("unknown bit op")
	}
	return nil
}

func (s *StackCompileVisitor) VisitMulOp(ctx *gen.MulOpContext) interface{} {
	// mulOp : MUL | DIV | MOD ;
	switch ctx.GetText() {
	case "*":
		s.Write(vm.InstrMul)
	case "/":
		s.Write(vm.InstrDiv)
	case "%":
		s.Write(vm.InstrMod)
	default:
		panic("unknown mul op")
	}
	return nil
}

func (s *StackCompileVisitor) VisitPowOp(ctx *gen.PowOpContext) interface{} {
	// powOp : POW ;
	s.Write(vm.InstrPow)
	return nil
}
