package compile

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/gen"
)

var _ gen.GsVisitor = (*GsDefineVisitor)(nil)

type GsDefineVisitor struct {
	*gen.BaseGsVisitor
	Log            InterpreterListener
	CurScope       Scope
	GlobalScope    *GlobalScope
	ScopeAllocator int32
	CurFuncSymbol  *FunctionSymbol
}

func NewGsDefineVisitor(log InterpreterListener) *GsDefineVisitor {
	gScope := NewGlobalScope()
	ret := &GsDefineVisitor{
		CurScope:    gScope,
		GlobalScope: gScope,
		Log:         log,
	}
	ret.BaseGsVisitor = &gen.BaseGsVisitor{ParseTreeVisitor: &gen.BaseVisitor{RealVisitor: ret}}
	return ret
}

func (g *GsDefineVisitor) VisitProgram(ctx *gen.ProgramContext) interface{} {
	mainFunc := NewFunctionSymbol("main", nil)
	g.GlobalScope.Define(mainFunc)
	g.CurFuncSymbol = mainFunc
	g.VisitChildren(ctx)
	return nil
}

func (g *GsDefineVisitor) VisitAssign(ctx *gen.AssignContext) interface{} {
	// lvalue (',' lvalue)* assignOp expr (',' expr)*
	for _, lvalue := range ctx.AllLvalue() {
		if qid := lvalue.Qid(); qid != nil {
			varName := qid.Primary().GetText()
			if varName == "$" {
				continue
			}
			g.CurScope.Define(NewVariableSymbol(varName, lvalue.GetStart()))
		}
	}
	return g.VisitChildren(ctx)
}

func (g *GsDefineVisitor) VisitSingleIter(ctx *gen.SingleIterContext) interface{} {
	g.CurScope.Define(NewVariableSymbol(ctx.ID().GetText(), ctx.ID().GetSymbol()))
	return nil
}

func (g *GsDefineVisitor) VisitDoubleIter(ctx *gen.DoubleIterContext) interface{} {
	for _, node := range ctx.AllID() {
		g.CurScope.Define(NewVariableSymbol(node.GetText(), node.GetSymbol()))
	}
	return nil
}

func (g *GsDefineVisitor) VisitQid(ctx *gen.QidContext) interface{} {
	if ctx.Primary().ID() != nil {
		refName := ctx.Primary().ID().GetText()
		if refName != "$" {
			if g.CurScope.Resolve(refName) == nil {
				g.Log.ErrorToken(ctx.GetStart(), "undefined variable: %s", refName)
			}
		}
	}
	return g.VisitChildren(ctx)
}

func (g *GsDefineVisitor) VisitStructDefinition(ctx *gen.StructDefinitionContext) interface{} {
	// 'type' ID 'struct' '{' ID (',' ID)* '}'
	structName := ctx.ID(0).GetText()
	structSymbol := NewStructSymbol(structName, ctx.ID(0).GetSymbol())
	structSymbol.EnclosingScope = g.CurScope
	for i := 1; i < len(ctx.AllID()); i++ {
		fieldName := ctx.AllID()[i].GetText()
		structSymbol.Define(NewFieldSymbol(structSymbol, fieldName, i-1, ctx.AllID()[i].GetSymbol()))
	}
	g.CurScope.Define(structSymbol)
	return nil
}

func (g *GsDefineVisitor) VisitFunctionDefinition(ctx *gen.FunctionDefinitionContext) interface{} {
	// 'func' ID '(' (ID (',' ID)* )? ')'  block
	preScope := g.CurScope
	preFuncSymbol := g.CurFuncSymbol

	funcName := ctx.ID(0).GetText()
	funcSymbol := NewFunctionSymbol(funcName, ctx.ID(0).GetSymbol())
	funcSymbol.SetScope(g.CurScope)
	g.CurFuncSymbol = funcSymbol
	// vardefs
	allIDs := ctx.AllID()
	for i := 1; i < len(allIDs); i++ {
		arg := allIDs[i].GetText()
		funcSymbol.Define(NewVariableSymbol(arg, allIDs[i].GetSymbol()))
	}
	// block
	localScope := &LocalScope{
		Symbols:        make(map[string]Symbol),
		GlobalDeclared: map[string]struct{}{},
		EnclosingScope: funcSymbol,
		BaseAllocAddr:  int32(len(funcSymbol.FormalArgs)),
		ID:             g.ScopeAllocator,
	}
	g.CurScope = localScope
	g.ScopeAllocator++
	// 保存函数的local scope
	funcSymbol.BodyScope = g.CurScope.(*LocalScope)

	g.Visit(ctx.Block())
	// 这里直接回退到funcSymbol的scope
	g.CurScope = preScope
	g.CurScope.Define(funcSymbol)
	g.CurFuncSymbol = preFuncSymbol
	return nil
}

func (g *GsDefineVisitor) VisitGlobalStmt(ctx *gen.GlobalStmtContext) interface{} {
	// 'global' ID
	varName := ctx.ID().GetText()
	if g.GlobalScope.Resolve(varName) == nil {
		g.GlobalScope.Define(NewVariableSymbol(varName, ctx.ID().GetSymbol()))
	}
	if l, ok := g.CurScope.(*LocalScope); ok {
		if l.Symbols[varName] != nil {
			g.Log.ErrorToken(ctx.GetStart(), "Syntax error: name '%s' is assigned to before global declaration %s", varName)
			return nil
		}
		l.GlobalDeclared[varName] = struct{}{}
	} else {
		g.Log.ErrorToken(ctx.GetStart(), "Syntax error: global statement can only write in func body")
		return nil
	}
	return nil
}

func (g *GsDefineVisitor) VisitReturnStmt(ctx *gen.ReturnStmtContext) interface{} {
	// 'return' (expr (',' expr )*)? NL
	returns := len(ctx.AllExpr())
	if g.CurFuncSymbol.ReturnNums == -1 {
		g.CurFuncSymbol.ReturnNums = returns
	} else if g.CurFuncSymbol.ReturnNums != returns {
		g.Log.ErrorToken(ctx.GetStart(), "return %d values not match pre defined %d in function %s line %d",
			returns, g.CurFuncSymbol.ReturnNums, g.CurFuncSymbol.Name, g.CurFuncSymbol.DefineToken.GetLine())
		return nil
	}

	return g.VisitChildren(ctx)
}

func (g *GsDefineVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(g)
}

func (g *GsDefineVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	ctx := node.(antlr.ParserRuleContext)
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(g)
	}
	return nil
}
