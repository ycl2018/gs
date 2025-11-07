package gen

import (
	"github.com/antlr4-go/antlr/v4"
)

var _ antlr.ParseTreeVisitor = (*BaseVisitor)(nil)

type BaseVisitor struct {
	realVisitor GsVisitor
}

func NewBaseVisitor(realVisitor GsVisitor) BaseVisitor {
	return BaseVisitor{realVisitor: realVisitor}
}

func (b BaseVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(b.realVisitor)
}

func (b BaseVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	ctx := node.(antlr.ParserRuleContext)
	if ctx.GetChildCount() == 1 {
		return b.Visit(ctx.GetChild(0).(antlr.ParseTree))
	}
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(b.realVisitor)
	}
	return nil
}

func (b BaseVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (b BaseVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}
