package compile

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/conf"
	"github.com/ycl2018/gs/consts"
	"github.com/ycl2018/gs/gen"
)

var _ gen.GsVisitor = (*StackCompileVisitor)(nil)

type StackCompileVisitor struct {
	gen.BaseGsVisitor
	Conf              *conf.CompileConf
	Env               *Env
	Log               InterpreterListener
	GlobalScope       *GlobalScope
	AllFuncs          []*FunctionSymbol // 所有函数,不包含主函数
	MainFunc          *FunctionSymbol
	CurFunc           *FunctionSymbol
	LoopStack         []*ForLoop
	CurScope          Scope
	CalledDefineFuncs map[string]int
}

func NewStackCompileVisitor(globalScope *GlobalScope, log InterpreterListener, conf *conf.CompileConf) *StackCompileVisitor {
	mainFunc := globalScope.Resolve("main").(*FunctionSymbol)
	s := &StackCompileVisitor{
		Conf:              conf,
		Env:               NewEnv(conf.Env),
		Log:               log,
		GlobalScope:       globalScope,
		MainFunc:          mainFunc,
		CurFunc:           mainFunc,
		CurScope:          globalScope,
		CalledDefineFuncs: make(map[string]int),
	}
	s.BaseGsVisitor = gen.BaseGsVisitor{ParseTreeVisitor: gen.NewBaseVisitor(s)}
	return s
}

func (s *StackCompileVisitor) WriteInstr(instr *consts.StackInstr, token antlr.Token) {
	s.CurFunc.Code = append(s.CurFunc.Code, instr)
	s.CurFunc.Debugger.Table = append(s.CurFunc.Debugger.Table, consts.Info{
		Line: token.GetLine(),
	})
}

func (s *StackCompileVisitor) VisitProgram(ctx *gen.ProgramContext) interface{} {
	s.VisitChildren(ctx)
	// 补充主函数Halt指令IR
	if len(s.MainFunc.Code) == 0 || s.MainFunc.Code[len(s.MainFunc.Code)-1].OpCode != consts.InstrHalt {
		s.MainFunc.Code = append(s.MainFunc.Code, &consts.StackInstr{
			OpCode: consts.InstrHalt,
		})
	}
	return nil
}

func (s *StackCompileVisitor) VisitFunctionDefinition(ctx *gen.FunctionDefinitionContext) interface{} {
	// 'def' ID '(' (vardef (',' vardef)* )? ')' slist
	// 生成函数定义 IR
	funcName := ctx.ID(0).GetText()
	scope := s.CurScope
	funcSymbol := scope.Resolve(funcName).(*FunctionSymbol)
	s.CurFunc = funcSymbol
	// 参数列表入栈由解释器负责自动执行
	// 生成函数体IR
	s.CurScope = funcSymbol.BodyScope
	ctx.Block().Accept(s)
	// 补充返回指令IR
	if len(s.CurFunc.Code) == 0 || s.CurFunc.Code[len(s.CurFunc.Code)-1].OpCode != consts.InstrReturn {
		// 函数体没有显式返回值时，补充返回nil指令IR
		s.Write(consts.InstrNil, ctx.GetStart())
		s.Write(consts.InstrReturn, ctx.GetStart())
	}
	s.AllFuncs = append(s.AllFuncs, funcSymbol)
	s.CurFunc = s.MainFunc
	s.CurScope = funcSymbol.Scope()
	return nil
}

func (s *StackCompileVisitor) VisitGoCall(ctx *gen.GoCallContext) interface{} {
	call := ctx.Call()
	switch call := call.(type) {
	case *gen.OuterCallContext:
		allExpr := call.AllExpr()
		for _, context := range allExpr {
			context.Accept(s)
		}
		s.loadOuterFunc(call)
		// 生成调用指令IR
		goInstr := consts.NewStackInstr(consts.InstrGoOuter, len(allExpr))
		s.WriteInstr(goInstr, ctx.GetStart())
	case *gen.InnerCallContext:
		allExpr := call.AllExpr()
		for _, context := range allExpr {
			context.Accept(s)
		}
		instr, err := s.innerCallType(call, allExpr)
		if err != nil {
			s.Log.ErrorToken(ctx.GetStart(), err.Error())
			return nil
		}
		switch instr.OpCode {
		case consts.InstrCallOuter:
			goInstr := consts.NewStackInstr(consts.InstrGoOuter, len(allExpr))
			s.WriteInstr(goInstr, ctx.GetStart())
		case consts.InstrCallDefine:
			goInstr := consts.NewStackInstr(consts.InstrGoDefine, instr.Operands)
			s.WriteInstr(goInstr, ctx.GetStart())
		case consts.InstrCall:
			s.Log.ErrorToken(ctx.GetStart(), "syntax err:go expr not support run code defined function")
		default:
			panic("unreachable")
		}
	}
	return nil
}

func (s *StackCompileVisitor) VisitInnerCall(ctx *gen.InnerCallContext) interface{} {
	allExpr := ctx.AllExpr()
	for _, context := range allExpr {
		context.Accept(s)
	}
	instr, err := s.innerCallType(ctx, allExpr)
	if err != nil {
		s.Log.ErrorToken(ctx.GetStart(), err.Error())
		return nil
	}
	s.WriteInstr(instr, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) innerCallType(ctx *gen.InnerCallContext, allExpr []gen.IExprContext) (*consts.StackInstr, error) {
	funcName := ctx.ID().GetText()
	if fSymbol := s.CurScope.Resolve(funcName); fSymbol != nil {
		switch v := fSymbol.(type) {
		case *FunctionSymbol:
			if len(allExpr) != len(v.FormalArgs) {
				return nil, fmt.Errorf("func:%s need %d args,but got %d", funcName, len(v.FormalArgs), len(allExpr))
			}
			// 生成调用指令IR
			callInstr := consts.NewStackInstr(consts.InstrCall, v.Address)
			return callInstr, nil
		case *VariableSymbol:
			// 生成调用指令IR
			s.EmitLoad(v, ctx.GetStart())
			callInstr := consts.NewStackInstr(consts.InstrCallOuter, len(allExpr))
			return callInstr, nil
		default:
			panic(fmt.Sprintf("unkown symbol type:%s to call", fSymbol))
		}
	}
	if defineFn := s.Conf.DefineFuncs.GetFunc(funcName); defineFn != nil {
		numIn := defineFn.NumIn
		if len(allExpr) != numIn {
			return nil, fmt.Errorf("func:%s need %d args,but got %d", funcName, numIn, len(allExpr))
		}
		var assign antlr.Tree = ctx
		for range 3 {
			if assign == nil {
				break
			}
			assign = assign.GetParent() // innerCall -> callAtom -> atomExpr -> assign
		}
		if assign != nil {
			if aCtx, ok := assign.(*gen.AssignContext); ok && len(aCtx.AllExpr()) == 1 {
				if lValueLen := len(aCtx.AllLvalue()); lValueLen != defineFn.NumOut {
					return nil, fmt.Errorf("func:%s return %d values,but assign to %d", funcName, defineFn.NumOut, lValueLen)
				}
			}
		}
		// 生成调用指令IR
		var addr int
		if v, ok2 := s.CalledDefineFuncs[funcName]; ok2 {
			addr = v
		} else {
			addr = len(s.CalledDefineFuncs)
			s.CalledDefineFuncs[funcName] = addr
		}
		callInstr := consts.NewStackInstr(consts.InstrCallDefine, addr)
		return callInstr, nil
	}
	return nil, fmt.Errorf("undefined func/symbol: %s", funcName)
}

func (s *StackCompileVisitor) VisitOuterCall(ctx *gen.OuterCallContext) interface{} {
	// primary accessor+ '(' (expr (',' expr)* ','?)? ')'
	allExpr := ctx.AllExpr()
	for _, context := range allExpr {
		context.Accept(s)
	}
	s.loadOuterFunc(ctx)
	// 生成调用指令IR
	callInstr := consts.NewStackInstr(consts.InstrCallOuter, len(allExpr))
	s.WriteInstr(callInstr, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitIntAtom(ctx *gen.IntAtomContext) interface{} {
	valStr := ctx.INT().GetText()
	ival, _ := strconv.ParseInt(valStr, 0, 64)
	s.Write(consts.InstrIConst, ctx.GetStart(), int(ival))
	return nil
}

func (s *StackCompileVisitor) VisitFloatAtom(ctx *gen.FloatAtomContext) interface{} {
	valStr := ctx.FLOAT().GetText()
	fval, _ := strconv.ParseFloat(valStr, 64)
	// 浮点数常量地址
	symbol := defineFloatConst(fval, s.GlobalScope)
	s.Write(consts.InstrConst, ctx.GetStart(), symbol.GetAddress())
	return nil
}

func (s *StackCompileVisitor) VisitStringAtom(ctx *gen.StringAtomContext) interface{} {
	valStr := ctx.STRING().GetText()
	// 字符串常量地址
	str, err := strconv.Unquote(valStr)
	if err != nil {
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("invalid string: %s", valStr))
		return nil
	}
	symbol := defineStringConst(str, s.GlobalScope)
	s.Write(consts.InstrConst, ctx.GetStart(), symbol.GetAddress())
	return nil
}

func (s *StackCompileVisitor) VisitInstance(ctx *gen.InstanceContext) interface{} {
	// instance : 'new' ID '{' (ID ':' expr (',' ID ':' expr)* ','?)? '}' ;
	structName := ctx.ID(0).GetText()
	structSymbol := s.CurScope.Resolve(structName)
	if structSymbol == nil {
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("undefined struct: %s", structName))
		return nil
	}
	s.Write(consts.InstrStruct, ctx.GetStart(), structSymbol.GetAddress())
	// 初始化结构体字段
	allIDs := ctx.AllID()[1:]
	allExprs := ctx.AllExpr()
	for i := 0; i < len(allIDs); i++ {
		fieldName := allIDs[i].GetText()
		allExprs[i].Accept(s)
		s.Write(consts.InstrFStore, ctx.GetStart(), s.defineStringConst(fieldName))
	}
	return nil
}

func (s *StackCompileVisitor) VisitQid(ctx *gen.QidContext) interface{} {
	s.loadQid(ctx)
	return nil
}

func (s *StackCompileVisitor) VisitAddOp(ctx *gen.AddOpContext) interface{} {
	op := ctx.GetText()
	switch op {
	case "+":
		s.Write(consts.InstrAdd, ctx.GetStart())
	case "-":
		s.Write(consts.InstrSub, ctx.GetStart())
	case "|":
		s.Write(consts.InstrBitOR, ctx.GetStart())
	case "^":
		s.Write(consts.InstrXOR, ctx.GetStart())
	default:
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("unknown addOp %s", op))
	}
	return nil
}

func (s *StackCompileVisitor) VisitCompOp(ctx *gen.CompOpContext) interface{} {
	// compOp	: EQ | LT | GT | GEQ | LEQ | NEQ ;
	op := ctx.GetText()
	switch op {
	case "==":
		s.Write(consts.InstrEQ, ctx.GetStart())
	case "<":
		s.Write(consts.InstrLT, ctx.GetStart())
	case ">":
		s.Write(consts.InstrGT, ctx.GetStart())
	case ">=":
		s.Write(consts.InstrGEQ, ctx.GetStart())
	case "<=":
		s.Write(consts.InstrLEQ, ctx.GetStart())
	case "!=":
		s.Write(consts.InstrNEQ, ctx.GetStart())
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
	// lvalue (',' lvalue)* assignOp expr (',' expr)*
	var lvalues []*gen.LvalueContext
	var exprs []gen.IExprContext
	for _, child := range ctx.GetChildren() {
		if qid, ok := child.(*gen.LvalueContext); ok {
			lvalues = append(lvalues, qid)
		} else if e, ok := child.(gen.IExprContext); ok {
			exprs = append(exprs, e)
		}
	}
	// exprs  from left to right
	for _, expr := range exprs {
		expr.Accept(s)
	}
	// build tuple
	if len(exprs) > 1 {
		s.Write(consts.InstrBuildTuple, ctx.GetStart(), len(exprs))
	}
	// unpack tuple
	if len(lvalues) > 1 {
		s.Write(consts.InstrUnpack, ctx.GetStart(), len(lvalues))
	}
	// lvalues from left to right
	//  lvalue (',' lvalue)* assignOp expr (',' expr)*
	for i := 0; i < len(lvalues); i++ {
		s.storeLvalue(lvalues[i])
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
		s.Write(consts.InstrBuildTuple, ctx.GetStart(), lenExprs)
	}
	if lenExprs == 0 {
		s.Write(consts.InstrNil, ctx.GetStart())
	}
	s.Write(consts.InstrReturn, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitInitRefCall(ctx *gen.InitRefCallContext) interface{} {
	qid := ctx.Qid()
	expr := ctx.Expr()
	var operand int
	if expr != nil {
		expr.Accept(s)
		operand = 1
	}
	s.loadQid(qid)
	s.Write(consts.InstrInitRef, ctx.GetStart(), operand)
	s.storeQid(qid)
	return nil
}

func (s *StackCompileVisitor) VisitPrintXCall(ctx *gen.PrintXCallContext) interface{} {
	exprs := ctx.AllExpr()
	for _, expr := range exprs {
		expr.Accept(s)
	}
	switch ctx.GetChild(0).(antlr.TerminalNode).GetSymbol().GetTokenType() {
	case gen.GsParserPRINT:
		s.Write(consts.InstrPrint, ctx.GetStart(), len(exprs))
	case gen.GsParserPRINTLN:
		s.Write(consts.InstrPrintln, ctx.GetStart(), len(exprs))
	case gen.GsParserPRINTF:
		s.Write(consts.InstrPrintf, ctx.GetStart(), len(exprs))
	case gen.GsParserSPRINTF:
		s.Write(consts.InstrSprintf, ctx.GetStart(), len(exprs))
	default:
		s.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("unknown printX %s", ctx.GetText()))
	}
	return nil
}

func (s *StackCompileVisitor) VisitBuiltinStmt(ctx *gen.BuiltinStmtContext) interface{} {
	s.VisitChildren(ctx)
	var noNeedPop bool
	switch call := ctx.BuiltinCall().(type) {
	case *gen.PrintXCallContext:
		tokenType := call.GetChild(0).(antlr.TerminalNode).GetSymbol().GetTokenType()
		switch tokenType {
		case gen.GsParserPRINT, gen.GsParserPRINTLN, gen.GsParserPRINTF:
			noNeedPop = true
		}
	case *gen.DeleteCallContext, *gen.InitRefCallContext:
		noNeedPop = true
	}
	if !noNeedPop {
		s.Write(consts.InstrPop, ctx.GetStart())
	}
	return nil
}

func (s *StackCompileVisitor) VisitLenCall(ctx *gen.LenCallContext) interface{} {
	ctx.Expr().Accept(s)
	s.Write(consts.InstrLen, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitNewFromTypeCall(ctx *gen.NewFromTypeCallContext) interface{} {
	expr := ctx.Expr()
	if expr == nil {
		s.Log.ErrorToken(ctx.GetStart(), "newFromType call need 1 arg")
		return nil
	}
	if expr.GetChildCount() == 1 {
		str, ok := expr.GetChild(0).(*gen.StringAtomContext)
		if ok {
			typeName := str.GetText()[1 : len(str.GetText())-1]
			typ := s.Conf.TypesAvailable.GetType(typeName)
			if typ == nil {
				s.Log.ErrorToken(str.GetStart(), "newFromType call with unknown type %s", typeName)
				return nil
			}
		}
	}
	expr.Accept(s)
	s.Write(consts.InstrNewFromType, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitAppendCall(ctx *gen.AppendCallContext) interface{} {
	for _, expr := range ctx.AllExpr() {
		expr.Accept(s)
	}
	if ctx.EXPAND() != nil {
		s.Write(consts.InstrAppendExpand, ctx.GetStart())
	} else {
		s.Write(consts.InstrAppend, ctx.GetStart(), len(ctx.AllExpr()))
	}
	return nil
}

func (s *StackCompileVisitor) VisitDeleteCall(ctx *gen.DeleteCallContext) interface{} {
	for _, expr := range ctx.AllExpr() {
		expr.Accept(s)
	}
	s.Write(consts.InstrDelete, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitConvertCall(ctx *gen.ConvertCallContext) interface{} {
	ctx.Expr().Accept(s)
	var oprand reflect.Kind
	switch t := ctx.GetChild(0).(antlr.TerminalNode).GetSymbol().GetTokenType(); t {
	case gen.GsParserUINT:
		oprand = reflect.Uint
	case gen.GsParserUINT8:
		oprand = reflect.Uint8
	case gen.GsParserUINT16:
		oprand = reflect.Uint16
	case gen.GsParserUINT32:
		oprand = reflect.Uint32
	case gen.GsParserUINT64:
		oprand = reflect.Uint64
	case gen.GsParserINTS:
		oprand = reflect.Int
	case gen.GsParserINT8:
		oprand = reflect.Int8
	case gen.GsParserINT16:
		oprand = reflect.Int16
	case gen.GsParserINT32:
		oprand = reflect.Int32
	case gen.GsParserINT64:
		oprand = reflect.Int64
	case gen.GsParserFLOAT32:
		oprand = reflect.Float32
	case gen.GsParserFLOAT64:
		oprand = reflect.Float64
	case gen.GsParserBOOL:
		oprand = reflect.Bool
	case gen.GsParserSTRINGS:
		oprand = reflect.String
	default:
		s.Log.ErrorToken(ctx.GetStart(), "unsupported convert to %d", t)
		return nil
	}
	s.Write(consts.InstrConvert, ctx.GetStart(), int(oprand))
	return nil
}

func (s *StackCompileVisitor) VisitIfStatement(ctx *gen.IfStatementContext) interface{} {
	if ctx.SimpleStmt() != nil {
		ctx.SimpleStmt().Accept(s)
	}
	ctx.Expr().Accept(s)
	// 生成分支跳转IR
	brfInstr := consts.NewStackInstr(consts.InstrBRF, placeholder)
	s.WriteInstr(brfInstr, ctx.GetStart())
	ctx.Block(0).Accept(s)
	hasElse := ctx.ELSE() != nil
	if hasElse {
		var brInstr = consts.NewStackInstr(consts.InstrBR, placeholder)
		s.WriteInstr(brInstr, ctx.GetStart())
		s.FillTarget(brfInstr)
		if elseBlock := ctx.Block(1); elseBlock != nil {
			elseBlock.Accept(s)
		} else if ifStmt := ctx.IfStatement(); ifStmt != nil {
			s.VisitIfStatement(ifStmt.(*gen.IfStatementContext))
		}
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
	brfInstr := consts.NewStackInstr(consts.InstrBRF, placeholder)
	s.WriteInstr(brfInstr, ctx.GetStart())
	ctx.Block().Accept(s)
	updateTarget := len(s.CurFunc.Code)
	if forUpdate != nil {
		forUpdate.Accept(s)
	}
	s.Write(consts.InstrBR, ctx.GetStart(), brTarget)
	s.FillTarget(brfInstr)
	loop := s.PopLoopStack()
	for _, br := range loop.Breaks {
		s.FillTarget(br)
	}
	for _, br := range loop.Continues {
		br.Operands = updateTarget
	}
	return nil
}

func (s *StackCompileVisitor) VisitForRangeStmt(ctx *gen.ForRangeStmtContext) interface{} {
	// 'for' iterVar '=' 'range' expr block
	s.PushLoopStack(&ForLoop{Token: ctx.GetStart()})
	ctx.Expr().Accept(s)
	brNilInstr := consts.NewStackInstr(consts.InstrBRNil, placeholder)
	s.WriteInstr(brNilInstr, ctx.GetStart())  // for safe range
	s.Write(consts.InstrIter, ctx.GetStart()) // push iter state
	rangeNextAddr := len(s.CurFunc.Code)
	s.Write(consts.InstrIterDone, ctx.GetStart())
	brtInstr := consts.NewStackInstr(consts.InstrBRT, placeholder)
	s.WriteInstr(brtInstr, ctx.GetStart())
	scope := s.CurScope
	switch iterVar := ctx.IterVar().(type) {
	case *gen.SingleIterContext:
		s.Write(consts.InstrIterNext, iterVar.GetStart(), 1)
		iter := scope.Resolve(iterVar.ID().GetText()).(*VariableSymbol)
		s.EmitStore(iter, iterVar.GetStart())
	case *gen.DoubleIterContext:
		s.Write(consts.InstrIterNext, iterVar.GetStart(), 2) //push k,v on stack
		first := scope.Resolve(iterVar.ID(0).GetText()).(*VariableSymbol)
		second := scope.Resolve(iterVar.ID(1).GetText()).(*VariableSymbol)
		s.EmitStore(second, iterVar.GetStart())
		s.EmitStore(first, iterVar.GetStart())
	}
	ctx.Block().Accept(s)
	s.Write(consts.InstrBR, ctx.GetStart(), rangeNextAddr)
	s.FillTarget(brNilInstr, brtInstr)

	loop := s.PopLoopStack()
	for _, br := range loop.Breaks {
		s.FillTarget(br)
	}
	for _, br := range loop.Continues {
		br.Operands = rangeNextAddr
	}
	s.Write(consts.InstrPop, ctx.GetStart()) // pop iter state
	return nil
}

func (s *StackCompileVisitor) VisitForCondStmt(ctx *gen.ForCondStmtContext) interface{} {
	// 'for' expr? block
	s.PushLoopStack(&ForLoop{Token: ctx.GetStart()})
	brTarget := len(s.CurFunc.Code)
	if ctx.Expr() != nil {
		ctx.Expr().Accept(s)
	}
	brfInstr := consts.NewStackInstr(consts.InstrBRF, placeholder)
	s.WriteInstr(brfInstr, ctx.GetStart())
	ctx.Block().Accept(s)
	s.Write(consts.InstrBR, ctx.GetStart(), brTarget)
	s.FillTarget(brfInstr)
	loop := s.PopLoopStack()
	for _, br := range loop.Breaks {
		s.FillTarget(br)
	}
	for _, br := range loop.Continues {
		br.Operands = brTarget
	}
	return nil
}

func (s *StackCompileVisitor) VisitCallStmt(ctx *gen.CallStmtContext) interface{} {
	ctx.Call().Accept(s)
	// discard return values
	s.Write(consts.InstrPop, ctx.GetStart()) // pop return value
	return nil
}

func (s *StackCompileVisitor) VisitBreakStmt(ctx *gen.BreakStmtContext) interface{} {
	breakInstr := consts.NewStackInstr(consts.InstrBR, placeholder)
	s.WriteInstr(breakInstr, ctx.GetStart())
	if s.GetCurrentLoop() == nil {
		s.Log.ErrorToken(ctx.GetStart(), "break statement not in loop")
		return nil
	}
	s.GetCurrentLoop().Breaks = append(s.GetCurrentLoop().Breaks, breakInstr)
	return nil
}

func (s *StackCompileVisitor) VisitContinueStmt(ctx *gen.ContinueStmtContext) interface{} {
	continueInstr := consts.NewStackInstr(consts.InstrBR, placeholder)
	s.WriteInstr(continueInstr, ctx.GetStart())
	if s.GetCurrentLoop() == nil {
		s.Log.ErrorToken(ctx.GetStart(), "continue statement not in loop")
		return nil
	}
	s.GetCurrentLoop().Continues = append(s.GetCurrentLoop().Continues, continueInstr)
	return nil
}

func (s *StackCompileVisitor) VisitIncrDecr(ctx *gen.IncrDecrContext) interface{} {
	// qid (INCR | DECR)
	incr := ctx.INCR() != nil
	if incr {
		s.Write(consts.InstrIConst, ctx.GetStart(), 1)
	} else {
		s.Write(consts.InstrIConst, ctx.GetStart(), -1)
	}
	lvalue := ctx.Lvalue().(*gen.LvalueContext)
	s.loadLvalue(lvalue)
	s.Write(consts.InstrAdd, ctx.GetStart())
	s.storeLvalue(lvalue)
	return nil
}

func (s *StackCompileVisitor) VisitSelfAssign(ctx *gen.SelfAssignContext) interface{} {
	// qid selfAssignOp expr
	lvalue := ctx.Lvalue().(*gen.LvalueContext)
	s.loadLvalue(lvalue)
	ctx.Expr().Accept(s)
	ctx.SelfAssignOp().Accept(s)
	s.storeLvalue(lvalue)
	return nil
}

func (s *StackCompileVisitor) VisitSelfAssignOp(ctx *gen.SelfAssignOpContext) interface{} {
	// selfAssignOp
	// '+=' | '-=' | '*=' | '/=' | '%='
	switch ctx.GetText() {
	case "+=":
		s.Write(consts.InstrAdd, ctx.GetStart())
	case "-=":
		s.Write(consts.InstrSub, ctx.GetStart())
	case "*=":
		s.Write(consts.InstrMul, ctx.GetStart())
	case "/=":
		s.Write(consts.InstrDiv, ctx.GetStart())
	case "%=":
		s.Write(consts.InstrMod, ctx.GetStart())
	default:
		s.Log.ErrorToken(ctx.GetStart(), "unknown selfAssignOp: %s", ctx.GetText())
	}
	return nil
}

func (s *StackCompileVisitor) VisitLogicalOrExpr(ctx *gen.LogicalOrExprContext) interface{} {
	if ctx.GetChildCount() == 1 {
		// must be constNode
		return s.VisitChildren(ctx)
	}
	exprs := ctx.AllExpr()
	l, r := exprs[0], exprs[1]
	l.Accept(s)
	brtInstr := consts.NewStackInstr(consts.InstrBRIfT, placeholder)
	s.WriteInstr(brtInstr, ctx.GetStart())
	r.Accept(s)
	s.FillTarget(brtInstr)
	return nil
}

func (s *StackCompileVisitor) VisitLogicalAndExpr(ctx *gen.LogicalAndExprContext) interface{} {
	if ctx.GetChildCount() == 1 {
		// must be constNode
		return s.VisitChildren(ctx)
	}
	exprs := ctx.AllExpr()
	l, r := exprs[0], exprs[1]
	l.Accept(s)
	brfInstr := consts.NewStackInstr(consts.InstrBRIfF, placeholder)
	s.WriteInstr(brfInstr, ctx.GetStart())
	r.Accept(s)
	s.FillTarget(brfInstr)
	return nil
}

func (s *StackCompileVisitor) VisitComparisonExpr(ctx *gen.ComparisonExprContext) interface{} {
	allExprs, cmpOp := ctx.AllExpr(), ctx.CompOp()
	if len(allExprs) < 2 {
		return s.VisitChildren(ctx)
	}
	l, r := allExprs[0], allExprs[1]
	l.Accept(s)
	r.Accept(s)
	var constKindl, constKindr consts.ConstNodeKind = -1, -1
	if lc, ok1 := toConstNode(l); ok1 {
		constKindl = lc.Kind
	}
	if rc, ok2 := toConstNode(r); ok2 {
		constKindr = rc.Kind
	}
	if constKindl == consts.ConstNodeKindInt || constKindr == consts.ConstNodeKindInt {
		s.Write(consts.InstrCmpInt, ctx.GetStart(), cmpOp.GetChild(0).(antlr.TerminalNode).GetSymbol().GetTokenType())
		return nil
	}
	if constKindl == consts.ConstNodeKindString || constKindr == consts.ConstNodeKindString {
		s.Write(consts.InstrCmpString, ctx.GetStart(), cmpOp.GetChild(0).(antlr.TerminalNode).GetSymbol().GetTokenType())
		return nil
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

func (s *StackCompileVisitor) VisitMulExpr(ctx *gen.MulExprContext) interface{} {
	var preOp *gen.MulOpContext
	for _, tree := range ctx.GetChildren() {
		switch tree := tree.(type) {
		default:
			tree.(antlr.RuleContext).Accept(s)
			if preOp != nil {
				preOp.Accept(s)
			}
		case *gen.MulOpContext:
			preOp = tree
		}
	}
	return nil
}

func (s *StackCompileVisitor) VisitNegAtom(ctx *gen.NegAtomContext) interface{} {
	if ctx.GetChildCount() == 1 {
		ctx.GetChild(0).(*consts.ConstNode).Accept(s)
		return nil
	}
	ctx.Atom().Accept(s)
	s.Write(consts.InstrNeg, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitTrueAtom(ctx *gen.TrueAtomContext) interface{} {
	s.Write(consts.InstrTrue, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitFalseAtom(ctx *gen.FalseAtomContext) interface{} {
	s.Write(consts.InstrFalse, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitNotAtom(ctx *gen.NotAtomContext) interface{} {
	ctx.Atom().Accept(s)
	s.Write(consts.InstrNot, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitNilAtom(ctx *gen.NilAtomContext) interface{} {
	s.Write(consts.InstrNil, ctx.GetStart())
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
	s.Write(consts.InstrArray, ctx.GetStart(), len(exprs))
	return nil
}

func (s *StackCompileVisitor) VisitSliceExpr(ctx *gen.SliceExprContext) interface{} {
	// sliceExpr : start=expr? ':' end=expr? ;
	var isLeft = true
	var l, r gen.IExprContext
	for _, tree := range ctx.GetChildren() {
		switch tree := tree.(type) {
		case gen.IExprContext:
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
		s.Write(consts.InstrIConst, ctx.GetStart(), -1)
	}
	if r != nil {
		r.Accept(s)
	} else {
		s.Write(consts.InstrIConst, ctx.GetStart(), -1)
	}
	// qidValue | l | r | InstrSlice
	s.Write(consts.InstrSliceSplit, ctx.GetStart())
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
	}
	s.Write(consts.InstrDict, ctx.GetStart(), len(entries)*2)
	return nil
}

func (s *StackCompileVisitor) VisitConstKeyEntry(ctx *gen.ConstKeyEntryContext) interface{} {
	// strKeyEntry : STRING ':' expr ;
	key := ctx.GetChildren()[0].(antlr.TerminalNode)
	switch key.GetSymbol().GetTokenType() {
	case gen.GsLexerSTRING:
		s.Write(consts.InstrConst, ctx.GetStart(), s.defineStringConst(ctx.STRING().GetText()[1:len(ctx.STRING().GetText())-1]))
	case gen.GsLexerINT:
		val, err := strconv.ParseInt(key.GetText(), 0, 64)
		if err != nil {
			panic(fmt.Sprintf("can't parse %s to int", key.GetText()))
		}
		s.Write(consts.InstrIConst, ctx.GetStart(), int(val))
	case gen.GsLexerFLOAT:
		val, err := strconv.ParseFloat(key.GetText(), 64)
		if err != nil {
			panic(fmt.Sprintf("can't parse %s to float", key.GetText()))
		}
		s.Write(consts.InstrConst, ctx.GetStart(), defineFloatConst(val, s.GlobalScope).GetAddress())
	case gen.GsParserTRUE:
		s.Write(consts.InstrTrue, ctx.GetStart())
	case gen.GsParserFALSE:
		s.Write(consts.InstrFalse, ctx.GetStart())
	default:
		panic(fmt.Sprintf("unknown tokenType:%s", key.GetText()))
	}

	ctx.Expr().Accept(s)
	return nil
}

func (s *StackCompileVisitor) VisitIdKeyEntry(ctx *gen.IdKeyEntryContext) interface{} {
	// idKeyEntry : ID ':' expr ;
	lvalue := ctx.Lvalue().(*gen.LvalueContext)
	s.loadLvalue(lvalue)
	ctx.Expr().Accept(s)
	return nil
}

func (s *StackCompileVisitor) VisitMulOp(ctx *gen.MulOpContext) interface{} {
	// mulOp : MUL | DIV | MOD ;
	switch ctx.GetText() {
	case "*":
		s.Write(consts.InstrMul, ctx.GetStart())
	case "/":
		s.Write(consts.InstrDiv, ctx.GetStart())
	case "%":
		s.Write(consts.InstrMod, ctx.GetStart())
	case "<<":
		s.Write(consts.InstrLShift, ctx.GetStart())
	case ">>":
		s.Write(consts.InstrRShift, ctx.GetStart())
	case "&":
		s.Write(consts.InstrBitAND, ctx.GetStart())
	default:
		panic("unknown mul op")
	}
	return nil
}

func (s *StackCompileVisitor) VisitDerefAtom(ctx *gen.DerefAtomContext) interface{} {
	// derefAtom : '*' lvalue ;
	ctx.Lvalue().Accept(s)
	s.Write(consts.InstrDeref, ctx.GetStart())
	return nil
}

func (s *StackCompileVisitor) VisitLvalue(ctx *gen.LvalueContext) interface{} {
	// if start by */&
	if v, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		switch v.GetText() {
		case "*":
			s.Visit(ctx.Lvalue())
			s.Write(consts.InstrDeref, ctx.GetStart())
			return nil
		default:
		}
	}
	s.VisitChildren(ctx)
	return nil
}
