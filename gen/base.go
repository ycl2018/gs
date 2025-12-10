package gen

import (
	"github.com/antlr4-go/antlr/v4"
)

var _ antlr.ParseTreeVisitor = (*BaseVisitor)(nil)

type BaseVisitor struct {
	RealVisitor GsVisitor
}

func NewBaseVisitor(realVisitor GsVisitor) BaseVisitor {
	return BaseVisitor{RealVisitor: realVisitor}
}

func (b BaseVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(b.RealVisitor)
}

func (b BaseVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	ctx := node.(antlr.ParserRuleContext)
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(b.RealVisitor)
	}
	return nil
}

func (b BaseVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (b BaseVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}
