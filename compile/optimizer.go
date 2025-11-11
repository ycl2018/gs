package compile

import (
	"fmt"
	"math"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/consts"
	"github.com/ycl2018/gs/gen"
)

type INode interface {
	AddChild(child antlr.RuleContext) antlr.RuleContext
}

var _ gen.GsVisitor = (*ConstOptimizer)(nil)

// ConstOptimizer do optimize:
//
//	fold const expr
//
// fold array/dict literal to const

type ConstOptimizer struct {
	gen.BaseGsVisitor
	FoldConstExpr bool
	Log           InterpreterListener
}

func NewConstOptimizer(log InterpreterListener) *ConstOptimizer {
	c := &ConstOptimizer{
		Log: log,
	}
	c.BaseGsVisitor = gen.BaseGsVisitor{ParseTreeVisitor: gen.NewBaseVisitor(c)}
	return c
}

func (c *ConstOptimizer) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(c)
}

func (c *ConstOptimizer) VisitChildren(node antlr.RuleNode) interface{} {
	ctx := node.(antlr.ParserRuleContext)
	for _, child := range ctx.GetChildren() {
		child.(antlr.ParseTree).Accept(c)
	}
	return nil
}

func (c *ConstOptimizer) VisitArrayLiteral(ctx *gen.ArrayLiteralContext) interface{} {
	c.VisitChildren(ctx)
	exprs := ctx.AllExpr()
	if len(exprs) == 0 {
		return nil
	}
	var sliceInit []any
	for _, tree := range exprs {
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
		case consts.ConstNodeKindInt:
			sliceInit = append(sliceInit, consts.ToIntValue(v2))
		case consts.ConstNodeKindFloat:
			sliceInit = append(sliceInit, consts.ToFloatValue(v2))
		case consts.ConstNodeKindString:
			sliceInit = append(sliceInit, consts.ToStringValue(v2))
		case consts.ConstNodeKindBool:
			sliceInit = append(sliceInit, consts.ToBoolValue(v2))
		default:
			return nil
		}
	}
	name := fmt.Sprintf("%d_%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	gen.InitEmptyArrayLiteralContext(ctx)
	ctx.AddChild(consts.NewConstNode(consts.ConstNodeKindList, &consts.SliceLiteralConst{Value: sliceInit, Name: name}))
	return nil
}

func (c *ConstOptimizer) VisitDictLiteral(ctx *gen.DictLiteralContext) interface{} {
	//dictLiteral : '{' (dictEntry (',' dictEntry)* ','?)? '}' ;
	//dictEntry
	//    :  (STRING|INT|FLOAT|TRUE|FALSE) ':' expr  #constKeyEntry
	//    |   qid ':' expr             #idKeyEntry  // 支持qid作为键
	//    ;
	c.VisitChildren(ctx)
	allEntries := ctx.AllDictEntry()
	if len(allEntries) == 0 {
		return nil
	}
	var m = map[consts.ConstNode]*consts.ConstNode{}
	for _, entry := range allEntries {
		constKey, ok := entry.(*gen.ConstKeyEntryContext)
		if !ok {
			return nil
		}
		keyTerminal := constKey.GetChild(0).(antlr.TerminalNode)
		keyType := keyTerminal.GetSymbol().GetTokenType()
		valueExpr := entry.(antlr.RuleContext).GetChild(2).(*gen.ExprContext)
		constValue, ok := toConstNode(valueExpr)
		if !ok {
			return nil
		}
		switch keyType {
		case gen.GsLexerSTRING:
			str := keyTerminal.GetText()[1 : len(keyTerminal.GetText())-1]
			keyValue := consts.NewConstNode(consts.ConstNodeKindString, str)
			m[*keyValue] = constValue
		case gen.GsLexerINT:
			intValue, err := strconv.ParseInt(keyTerminal.GetText(), 0, 64)
			if err != nil {
				panic(err)
			}
			keyValue := consts.NewConstNode(consts.ConstNodeKindInt, int(intValue))
			m[*keyValue] = constValue
		case gen.GsLexerFLOAT:
			floatValue, err := strconv.ParseFloat(keyTerminal.GetText(), 64)
			if err != nil {
				panic(err)
			}
			keyValue := consts.NewConstNode(consts.ConstNodeKindFloat, floatValue)
			m[*keyValue] = constValue
		case gen.GsLexerTRUE:
			keyValue := consts.NewConstNode(consts.ConstNodeKindBool, true)
			m[*keyValue] = constValue
		case gen.GsLexerFALSE:
			keyValue := consts.NewConstNode(consts.ConstNodeKindBool, false)
			m[*keyValue] = constValue
		default:
			panic(fmt.Sprintf("unknown key: %s", keyTerminal))
		}
	}
	name := fmt.Sprintf("%d_%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	gen.InitEmptyDictLiteralContext(ctx)
	ctx.AddChild(consts.NewConstNode(consts.ConstNodeKindMap, &consts.MapLiteralConst{Map: m, Name: name}))
	return nil
}

func toConstNode(valueExpr *gen.ExprContext) (*consts.ConstNode, bool) {
	valueChildren := valueExpr.GetChild(0).(*gen.LogicalOrExprContext)
	if len(valueChildren.GetChildren()) != 1 {
		return nil, false
	}
	constValue, ok2 := valueChildren.GetChild(0).(*consts.ConstNode)
	if !ok2 {
		return nil, false
	}
	return constValue, false
}

func (c *ConstOptimizer) VisitLogicalOrExpr(ctx *gen.LogicalOrExprContext) interface{} {
	// logicalAndExpr (OR logicalAndExpr)*
	c.VisitChildren(ctx)
	var newChildren []antlr.Tree
	var applied bool
	for _, tree := range ctx.GetChildren() {
		var added bool
		if len(tree.GetChildren()) == 1 {
			if v, ok := tree.GetChild(0).(*consts.ConstNode); ok {
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
				if top.Kind == consts.ConstNodeKindBool && top.Value.(bool) == true {
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
			if v, ok := tree.GetChild(0).(*consts.ConstNode); ok {
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
				if top.Kind == consts.ConstNodeKindBool && top.Value.(bool) == false {
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
			if v, ok := tree.GetChild(0).(*consts.ConstNode); ok {
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
					if top.Kind == consts.ConstNodeKindString && pre.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(pre) == consts.ToStringValue(top))
					} else if top.Kind == consts.ConstNodeKindBool && pre.Kind == consts.ConstNodeKindBool {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToBoolValue(pre) == consts.ToBoolValue(top))
					}
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) == consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) == consts.ToIntValue(top))
					}

				case "<":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) < consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) < consts.ToIntValue(top))
					}
				case ">":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) > consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) > consts.ToIntValue(top))
					}
				case "!=":
					if top.Kind == consts.ConstNodeKindString && pre.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(pre) != consts.ToStringValue(top))
					} else if top.Kind == consts.ConstNodeKindBool && pre.Kind == consts.ConstNodeKindBool {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToBoolValue(pre) != consts.ToBoolValue(top))
					}
					if top.Kind == consts.ConstNodeKindFloat && pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) != consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) != consts.ToIntValue(top))
					}
				case ">=":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) >= consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) >= consts.ToIntValue(top))
					}
				case "<=":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) <= consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) <= consts.ToIntValue(top))
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
			if v, ok := tree.GetChild(0).(*consts.ConstNode); ok {
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
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(pre)+consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)+consts.ToIntValue(top))
					}

				case "-":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(pre)-consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)-consts.ToIntValue(top))
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
			if v, ok := tree.GetChild(0).(*consts.ConstNode); ok {
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
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)&consts.ToIntValue(top))
				case "&":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)|consts.ToIntValue(top))
				case "^":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)^consts.ToIntValue(top))
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
			if v, ok := tree.GetChild(0).(*consts.ConstNode); ok {
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
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(pre)*consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)*consts.ToIntValue(top))
					}

				case "/":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(pre)/consts.ToFloatValue(top))
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)/consts.ToIntValue(top))
					}
				case "%":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)/consts.ToIntValue(top))
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
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindFloat, f))
		case *gen.IntAtomContext:
			applied = true
			f, err := strconv.ParseInt(t.GetText(), 0, 64)
			if err != nil {
				c.Log.ErrorToken(t.GetStart(), "cannot parse float from:%s", t.GetText())
				return nil
			}
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindInt, int(f)))
		case *gen.StringAtomContext:
			applied = true
			str := t.GetText()[1 : len(t.GetText())-1]
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindString, str))
		case *gen.TrueAtomContext:
			applied = true
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindBool, true))
		case *gen.FalseAtomContext:
			applied = true
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindBool, false))
		case *gen.ParenAtomContext:
			// expr -> logicalor
			added := false
			expr := t.Expr()
			if constNode, ok := toConstNode(expr.(*gen.ExprContext)); ok {
				newChildren = append(newChildren, constNode)
				applied = true
				added = true
			}
			if !added {
				newChildren = append(newChildren, expr)
			}
		default:
			newChildren = append(newChildren, t)
		}
		if len(newChildren) > 2 {
			// 合并
			top, ok1 := newChildren[len(newChildren)-1].(*consts.ConstNode)
			pre, ok2 := newChildren[len(newChildren)-3].(*consts.ConstNode)
			if ok1 && ok2 {
				newChildren = newChildren[:len(newChildren)-3]
				newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindFloat, math.Pow(consts.ToFloatValue(pre), consts.ToFloatValue(top))))
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
