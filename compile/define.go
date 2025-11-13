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
	Scopes         map[antlr.ParserRuleContext]Scope
	GlobalScope    *GlobalScope
	ScopeAllocator int32
}

func (g *GsDefineVisitor) VisitAssign(ctx *gen.AssignContext) interface{} {
	// qid (',' qid)* assignOp expr (',' expr)*
	for _, qid := range ctx.AllQid() {
		varName := qid.Primary().GetText()
		if g.CurScope.Resolve(varName) == nil {
			// 优先使用全局变量
			g.CurScope.Define(NewVariableSymbol(varName, qid.GetStart()))
		}
	}
	return g.VisitChildren(ctx)
}

func (g *GsDefineVisitor) VisitForRangeStmt(ctx *gen.ForRangeStmtContext) interface{} {
	g.SaveScope(ctx, g.CurScope)
	return g.VisitChildren(ctx)
}

func (g *GsDefineVisitor) VisitSingleIter(ctx *gen.SingleIterContext) interface{} {
	g.SaveScope(ctx, g.CurScope)
	g.CurScope.Define(NewVariableSymbol(ctx.ID().GetText(), ctx.ID().GetSymbol()))
	return nil
}

func (g *GsDefineVisitor) VisitDoubleIter(ctx *gen.DoubleIterContext) interface{} {
	g.SaveScope(ctx, g.CurScope)
	for _, node := range ctx.AllID() {
		g.CurScope.Define(NewVariableSymbol(node.GetText(), node.GetSymbol()))
	}
	return nil
}

func (g *GsDefineVisitor) VisitQidAtom(ctx *gen.QidAtomContext) interface{} {
	g.SaveScope(ctx, g.CurScope)
	return g.VisitChildren(ctx)
}

func (g *GsDefineVisitor) VisitQid(ctx *gen.QidContext) interface{} {
	g.SaveScope(ctx, g.CurScope)
	if ctx.Primary().ID() != nil {
		refName := ctx.Primary().ID().GetText()
		if g.CurScope.Resolve(refName) == nil {
			g.Log.ErrorToken(ctx.GetStart(), "undefined variable: %s", refName)
		}
	}
	return g.VisitChildren(ctx)
}

func NewGsDefineVisitor(log InterpreterListener) *GsDefineVisitor {
	gScope := NewGlobalScope()
	ret := &GsDefineVisitor{
		CurScope:    gScope,
		GlobalScope: gScope,
		Scopes:      make(map[antlr.ParserRuleContext]Scope),
		Log:         log,
	}
	ret.BaseGsVisitor = &gen.BaseGsVisitor{ParseTreeVisitor: &gen.BaseVisitor{RealVisitor: ret}}
	return ret
}

func (g *GsDefineVisitor) SaveScope(ctx antlr.ParserRuleContext, scope Scope) {
	g.Scopes[ctx] = scope
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
	g.SaveScope(ctx, g.CurScope)
	funcName := ctx.ID(0).GetText()
	funcSymbol := NewFunctionSymbol(funcName, ctx.ID(0).GetSymbol())
	funcSymbol.SetScope(g.CurScope)
	// vardefs
	for i := 1; i < len(ctx.AllID()); i++ {
		arg := ctx.AllID()[i].GetText()
		funcSymbol.Define(NewVariableSymbol(arg, ctx.AllID()[i].GetSymbol()))
	}
	// block
	g.CurScope = &LocalScope{
		Symbols:        make(map[string]Symbol),
		EnclosingScope: funcSymbol,
		BaseAllocAddr:  int32(len(funcSymbol.FormalArgs)),
		ID:             g.ScopeAllocator,
	}
	g.ScopeAllocator++
	// 保存函数的local scope
	funcSymbol.BodyScope = g.CurScope.(*LocalScope)

	g.Visit(ctx.Block())
	// 这里直接回退到funcSymbol的scope
	g.CurScope = preScope
	g.CurScope.Define(funcSymbol)
	return nil
}

func (g *GsDefineVisitor) VisitCall(ctx *gen.CallContext) interface{} {
	// name=ID '(' (expr (',' expr )*)? ')'
	g.SaveScope(ctx, g.CurScope)
	return g.VisitChildren(ctx)
}

func (g *GsDefineVisitor) VisitInstance(ctx *gen.InstanceContext) interface{} {
	// 'new' sname=ID NL
	g.SaveScope(ctx, g.CurScope)
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
