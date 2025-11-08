package compile

import (
	"fmt"
	"math"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/consts"
	"github.com/ycl2018/gs/gen"
)

type Parenter interface {
	AddChild(child antlr.RuleContext) antlr.RuleContext
}

var _ gen.GsVisitor = (*ConstOptimizer)(nil)

type ConstOptimizer struct {
	gen.BaseGsVisitor
	FoldConstExpr bool
	Log           InterpreterListener
}

func NewConstOptimizer(log InterpreterListener) *ConstOptimizer {
	c := &ConstOptimizer{
		Log: log,
	}
	c.BaseGsVisitor = gen.BaseGsVisitor{gen.NewBaseVisitor(c)}
	return c
}

func (g *ConstOptimizer) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(g)
}

func (g *ConstOptimizer) VisitChildren(node antlr.RuleNode) interface{} {
	ctx := node.(antlr.ParserRuleContext)
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(g)
	}
	return nil
}

func (c *ConstOptimizer) VisitArrayLiteral(ctx *gen.ArrayLiteralContext) interface{} {
	c.VisitChildren(ctx)
	var sliceInit []any
	for _, tree := range ctx.AllExpr() {
		if len(tree.GetChildren()) != 1 {
			return nil
		}
		v := tree.(*gen.ExprContext).GetChild(0)
		if v.GetChildCount() != 1 {
			return nil
		}
		v2, ok2 := v.GetChild(0).(*consts.ConstNode)
		if !ok2 {
			return nil
		}
		switch v2.Kind {
		case consts.ConstKindInt:
			sliceInit = append(sliceInit, consts.ToIntValue(v2))
		case consts.ConstKindFloat:
			sliceInit = append(sliceInit, consts.ToFloatValue(v2))
		case consts.ConstKindString:
			sliceInit = append(sliceInit, consts.ToStringValue(v2))
		case consts.ConstKindBool:
			sliceInit = append(sliceInit, consts.ToBoolValue(v2))
		default:
			return nil
		}
	}
	name := fmt.Sprintf("%d_%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	gen.InitEmptyArrayLiteralContext(ctx)
	ctx.AddChild(consts.NewConstNode(consts.ConstKindList, &consts.SliceInitConst{Value: sliceInit, Name: name}))
	return nil
}

func (c *ConstOptimizer) VisitDictLiteral(ctx *gen.DictLiteralContext) interface{} {
	//dictLiteral : '{' (dictEntry (',' dictEntry)* ','?)? '}' ;
	//dictEntry
	//    :  (STRING|INT|FLOAT|TRUE|FALSE) ':' expr  #constKeyEntry
	//    |   qid ':' expr             #idKeyEntry  // 支持qid作为键
	//    ;
	c.VisitChildren(ctx)
	var m = map[*consts.ConstNode]*consts.ConstNode{}
	for _, entry := range ctx.AllDictEntry() {
		constKey, ok := entry.(*gen.ConstKeyEntryContext)
		if !ok {
			return nil
		}
		keyTerminal := constKey.GetChild(0).(antlr.TerminalNode)
		keyType := keyTerminal.GetSymbol().GetTokenType()
		valueExpr := entry.(antlr.RuleContext).GetChild(2).(*gen.ExprContext)
		valueChildren := valueExpr.GetChild(0).(*gen.LogicalOrExprContext)
		if len(valueChildren.GetChildren()) != 1 {
			return nil
		}
		constValue, ok2 := valueChildren.GetChild(0).(*consts.ConstNode)
		if !ok2 {
			return nil
		}
		switch keyType {
		case gen.GsLexerSTRING:
			str := keyTerminal.GetText()[1 : len(keyTerminal.GetText())-1]
			keyValue := consts.NewConstNode(consts.ConstKindString, str)
			m[keyValue] = constValue
		case gen.GsLexerINT:
			intValue, err := strconv.ParseInt(keyTerminal.GetText(), 0, 64)
			if err != nil {
				panic(err)
			}
			keyValue := consts.NewConstNode(consts.ConstKindInt, int(intValue))
			m[keyValue] = constValue
		case gen.GsLexerFLOAT:
			floatValue, err := strconv.ParseFloat(keyTerminal.GetText(), 64)
			if err != nil {
				panic(err)
			}
			keyValue := consts.NewConstNode(consts.ConstKindFloat, floatValue)
			m[keyValue] = constValue
		case gen.GsLexerTRUE:
			keyValue := consts.NewConstNode(consts.ConstKindBool, true)
			m[keyValue] = constValue
		case gen.GsLexerFALSE:
			keyValue := consts.NewConstNode(consts.ConstKindBool, false)
			m[keyValue] = constValue
		default:
			panic(fmt.Sprintf("unknown key: %s", keyTerminal))
		}
	}
	name := fmt.Sprintf("%d_%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	gen.InitEmptyDictLiteralContext(ctx)
	ctx.AddChild(consts.NewConstNode(consts.ConstKindMap, &consts.MapInitConst{Map: m, Name: name}))
	return nil
}

func (c *ConstOptimizer) VisitLogicalOrExpr(ctx *gen.LogicalOrExprContext) interface{} {
	// logicalAndExpr (OR logicalAndExpr)*
	c.VisitChildren(ctx)
	var newChildren []antlr.Tree
	var applied bool
	for _, tree := range ctx.GetChildren() {
		var added bool
		if len(tree.GetChildren()) == 1 {
			if v, ok := tree.GetChildren()[0].(*consts.ConstNode); ok {
				newChildren = append(newChildren, v)
				applied = true
				added = true
			}
		}
		if !added {
			newChildren = append(newChildren, tree)
		}
		if len(newChildren) > 0 {
			top, ok := newChildren[len(newChildren)-1].(*consts.ConstNode)
			if ok {
				if top.Kind == consts.ConstKindBool && top.Value.(bool) == true {
					newChildren = []antlr.Tree{top}
					c.FoldConstExpr = true
					break
				}
			}
		}
	}
	if applied {
		gen.InitEmptyLogicalOrExprContext(ctx)
		for _, child := range newChildren {
			ctx.AddChild(child.(antlr.RuleContext))
		}
	}
	return nil
}

func (c *ConstOptimizer) VisitLogicalAndExpr(ctx *gen.LogicalAndExprContext) interface{} {
	// comparisonExpr (AND comparisonExpr)* ;
	c.VisitChildren(ctx)
	var newChildren []antlr.Tree
	var applied bool
	for _, tree := range ctx.GetChildren() {
		var added bool
		if len(tree.GetChildren()) == 1 {
			if v, ok := tree.GetChildren()[0].(*consts.ConstNode); ok {
				newChildren = append(newChildren, v)
				applied = true
				added = true
			}
		}
		if !added {
			newChildren = append(newChildren, tree)
		}
		if len(newChildren) > 0 {
			top, ok := newChildren[len(newChildren)-1].(*consts.ConstNode)
			if ok {
				if top.Kind == consts.ConstKindBool && top.Value.(bool) == false {
					newChildren = []antlr.Tree{top}
					c.FoldConstExpr = true
					break
				}
			}
		}
	}
	if applied {
		gen.InitEmptyLogicalAndExprContext(ctx)
		for _, child := range newChildren {
			ctx.AddChild(child.(antlr.RuleContext))
		}
	}
	return nil
}

func (c *ConstOptimizer) VisitComparisonExpr(ctx *gen.ComparisonExprContext) interface{} {
	c.VisitChildren(ctx)
	//addExpr (compOp addExpr)? ;
	var newChildren []antlr.Tree
	var applied bool
	for _, tree := range ctx.GetChildren() {
		var added bool
		if len(tree.GetChildren()) == 1 {
			if v, ok := tree.GetChildren()[0].(*consts.ConstNode); ok {
				newChildren = append(newChildren, v)
				applied = true
				added = true
			}
		}
		if !added {
			newChildren = append(newChildren, tree)
		}
		if len(newChildren) > 2 {
			// 合并，只处理相邻的常数
			top, ok1 := newChildren[len(newChildren)-1].(*consts.ConstNode)
			pre, ok2 := newChildren[len(newChildren)-3].(*consts.ConstNode)
			if ok1 && ok2 {
				op := newChildren[len(newChildren)-2].(*gen.CompOpContext).GetText()
				var constNode *consts.ConstNode
				// EQ | LT | GT | NEQ | GEQ | LEQ
				switch op {
				case "==":
					if top.Kind == consts.ConstKindString && pre.Kind == consts.ConstKindString {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToStringValue(pre) == consts.ToStringValue(top))
					} else if top.Kind == consts.ConstKindBool && pre.Kind == consts.ConstKindBool {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToBoolValue(pre) == consts.ToBoolValue(top))
					}
					if top.Kind == consts.ConstKindFloat || pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToFloatValue(pre) == consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToIntValue(pre) == consts.ToIntValue(top))
					}

				case "<":
					if top.Kind == consts.ConstKindFloat || pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToFloatValue(pre) < consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToIntValue(pre) < consts.ToIntValue(top))
					}
				case ">":
					if top.Kind == consts.ConstKindFloat || pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToFloatValue(pre) > consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToIntValue(pre) > consts.ToIntValue(top))
					}
				case "!=":
					if top.Kind == consts.ConstKindString && pre.Kind == consts.ConstKindString {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToStringValue(pre) != consts.ToStringValue(top))
					} else if top.Kind == consts.ConstKindBool && pre.Kind == consts.ConstKindBool {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToBoolValue(pre) != consts.ToBoolValue(top))
					}
					if top.Kind == consts.ConstKindFloat && pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToFloatValue(pre) != consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToIntValue(pre) != consts.ToIntValue(top))
					}
				case ">=":
					if top.Kind == consts.ConstKindFloat || pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToFloatValue(pre) >= consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToIntValue(pre) >= consts.ToIntValue(top))
					}
				case "<=":
					if top.Kind == consts.ConstKindFloat || pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToFloatValue(pre) <= consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindBool, consts.ToIntValue(pre) <= consts.ToIntValue(top))
					}
				default:
					panic(fmt.Sprintf("unknown op: %v", op))
				}
				newChildren = newChildren[:len(newChildren)-3]
				newChildren = append(newChildren, constNode)
				c.FoldConstExpr = true
			}
		}
	}
	if applied {
		gen.InitEmptyComparisonExprContext(ctx)
		for _, child := range newChildren {
			ctx.AddChild(child.(antlr.RuleContext))
		}
	}
	return nil
}

func (c *ConstOptimizer) VisitAddExpr(ctx *gen.AddExprContext) interface{} {
	//
	c.VisitChildren(ctx)
	var newChildren []antlr.Tree
	var applied bool
	for _, tree := range ctx.GetChildren() {
		var added bool
		if len(tree.GetChildren()) == 1 {
			if v, ok := tree.GetChildren()[0].(*consts.ConstNode); ok {
				newChildren = append(newChildren, v)
				applied = true
				added = true
			}
		}
		if !added {
			newChildren = append(newChildren, tree)
		}
		if len(newChildren) > 2 {
			// 合并，只处理相邻的常数
			top, ok1 := newChildren[len(newChildren)-1].(*consts.ConstNode)
			pre, ok2 := newChildren[len(newChildren)-3].(*consts.ConstNode)
			if ok1 && ok2 {
				op := newChildren[len(newChildren)-2].(*gen.AddOpContext).GetText()
				var constNode *consts.ConstNode
				switch op {
				case "+":
					if top.Kind == consts.ConstKindFloat || pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindFloat, consts.ToFloatValue(pre)+consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindInt, consts.ToIntValue(pre)+consts.ToIntValue(top))
					}

				case "-":
					if top.Kind == consts.ConstKindFloat || pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindFloat, consts.ToFloatValue(pre)-consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindInt, consts.ToIntValue(pre)-consts.ToIntValue(top))
					}
				default:
					panic(fmt.Sprintf("unknown op: %v", op))
				}
				newChildren = newChildren[:len(newChildren)-3]
				newChildren = append(newChildren, constNode)
				c.FoldConstExpr = true
			}
		}
	}
	if applied {
		gen.InitEmptyAddExprContext(ctx)
		for _, child := range newChildren {
			ctx.AddChild(child.(antlr.RuleContext))
		}
	}
	return nil
}

func (c *ConstOptimizer) VisitBinExpr(ctx *gen.BinExprContext) interface{} {
	// mulExpr (bitOp mulExpr)*
	c.VisitChildren(ctx)
	var newChildren []antlr.Tree
	var applied bool
	for _, tree := range ctx.GetChildren() {
		var added bool
		if len(tree.GetChildren()) == 1 {
			if v, ok := tree.GetChildren()[0].(*consts.ConstNode); ok {
				newChildren = append(newChildren, v)
				applied = true
				added = true
			}
		}
		if !added {
			newChildren = append(newChildren, tree)
		}
		if len(newChildren) > 2 {
			// 合并
			top, ok1 := newChildren[len(newChildren)-1].(*consts.ConstNode)
			pre, ok2 := newChildren[len(newChildren)-3].(*consts.ConstNode)
			if ok1 && ok2 {
				op := newChildren[len(newChildren)-2].(*gen.BitOpContext).GetText()
				var constNode *consts.ConstNode
				switch op {
				case "|":

					constNode = consts.NewConstNode(consts.ConstKindInt, consts.ToIntValue(pre)&consts.ToIntValue(top))
				case "&":
					constNode = consts.NewConstNode(consts.ConstKindInt, consts.ToIntValue(pre)|consts.ToIntValue(top))
				case "^":
					constNode = consts.NewConstNode(consts.ConstKindInt, consts.ToIntValue(pre)^consts.ToIntValue(top))
				default:
					panic(fmt.Sprintf("unknown op: %v", op))
				}
				newChildren = newChildren[:len(newChildren)-3]
				newChildren = append(newChildren, constNode)
				c.FoldConstExpr = true
			}
		}
	}
	if applied {
		gen.InitEmptyBinExprContext(ctx)
		for _, child := range newChildren {
			ctx.AddChild(child.(antlr.RuleContext))
		}
	}
	return nil
}

func (c *ConstOptimizer) VisitMulExpr(ctx *gen.MulExprContext) interface{} {
	c.VisitChildren(ctx)
	var newChildren []antlr.Tree
	var applied bool
	for _, tree := range ctx.GetChildren() {
		var added bool
		if len(tree.GetChildren()) == 1 {
			if v, ok := tree.GetChildren()[0].(*consts.ConstNode); ok {
				newChildren = append(newChildren, v)
				applied = true
				added = true
			}
		}
		if !added {
			newChildren = append(newChildren, tree)
		}
		if len(newChildren) > 2 {
			// 合并
			top, ok1 := newChildren[len(newChildren)-1].(*consts.ConstNode)
			pre, ok2 := newChildren[len(newChildren)-3].(*consts.ConstNode)
			if ok1 && ok2 {
				op := newChildren[len(newChildren)-2].(*gen.MulOpContext).GetText()
				var constNode *consts.ConstNode
				switch op {
				case "*":
					if top.Kind == consts.ConstKindFloat || pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindFloat, consts.ToFloatValue(pre)*consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindInt, consts.ToIntValue(pre)*consts.ToIntValue(top))
					}

				case "/":
					if top.Kind == consts.ConstKindFloat || pre.Kind == consts.ConstKindFloat {
						constNode = consts.NewConstNode(consts.ConstKindFloat, consts.ToFloatValue(pre)/consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstKindInt, consts.ToIntValue(pre)/consts.ToIntValue(top))
					}
				case "%":
					constNode = consts.NewConstNode(consts.ConstKindInt, consts.ToIntValue(pre)/consts.ToIntValue(top))
				default:
					panic(fmt.Sprintf("unknown op: %v", op))
				}
				newChildren = newChildren[:len(newChildren)-3]
				newChildren = append(newChildren, constNode)
				c.FoldConstExpr = true
			}
		}
	}

	if applied {
		gen.InitEmptyMulExprContext(ctx)
		for _, child := range newChildren {
			ctx.AddChild(child.(antlr.RuleContext))
		}
	}
	return nil
}

func (c *ConstOptimizer) VisitPowExpr(ctx *gen.PowExprContext) interface{} {
	c.VisitChildren(ctx)
	var newChildren []antlr.Tree
	var applied bool
	for _, tree := range ctx.GetChildren() {
		switch t := tree.(type) {
		case *gen.FloatAtomContext:
			applied = true
			f, err := strconv.ParseFloat(t.GetText(), 64)
			if err != nil {
				c.Log.ErrorToken(t.GetStart(), "cannot parse float from:%s", t.GetText())
				return nil
			}
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstKindFloat, f))
		case *gen.IntAtomContext:
			applied = true
			f, err := strconv.ParseInt(t.GetText(), 0, 64)
			if err != nil {
				c.Log.ErrorToken(t.GetStart(), "cannot parse float from:%s", t.GetText())
				return nil
			}
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstKindInt, int(f)))
		case *gen.StringAtomContext:
			applied = true
			str := t.GetText()[1 : len(t.GetText())-1]
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstKindString, str))
		case *gen.TrueAtomContext:
			applied = true
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstKindBool, true))
		case *gen.FalseAtomContext:
			applied = true
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstKindBool, false))
		default:
			newChildren = append(newChildren, t)
		}
		if len(newChildren) > 2 {
			// 合并
			top, ok1 := newChildren[len(newChildren)-1].(*consts.ConstNode)
			pre, ok2 := newChildren[len(newChildren)-3].(*consts.ConstNode)
			if ok1 && ok2 {
				newChildren = newChildren[:len(newChildren)-3]
				newChildren = append(newChildren, consts.NewConstNode(consts.ConstKindFloat, math.Pow(consts.ToFloatValue(pre), consts.ToFloatValue(top))))
			}
			c.FoldConstExpr = true
		}
	}
	if applied {
		gen.InitEmptyPowExprContext(ctx)
		for _, child := range newChildren {
			ctx.AddChild(child.(antlr.RuleContext))
		}
	}
	return nil
}
