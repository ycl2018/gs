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

type INode interface {
	AddChild(child antlr.RuleContext) antlr.RuleContext
}

var _ gen.GsVisitor = (*ConstOptimizer)(nil)

// ConstOptimizer do optimize:
//  1. fold const expr
//  2. fold array/dict literal to const
//  3. in func: array -> map
type ConstOptimizer struct {
	gen.BaseGsVisitor
	Conf          *conf.CompileConf
	GlobalScope   *GlobalScope
	FoldConstExpr bool
	Log           InterpreterListener
}

func NewConstOptimizer(log InterpreterListener, conf *conf.CompileConf, scope *GlobalScope) *ConstOptimizer {
	c := &ConstOptimizer{
		Log:         log,
		Conf:        conf,
		GlobalScope: scope,
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
	start := ctx.GetStart()
	for range ctx.GetChildCount() {
		ctx.RemoveLastChild()
	}
	ctx.AddChild(consts.NewConstNode(consts.ConstNodeKindList, &consts.SliceLiteralConst{Value: sliceInit, Name: name}, start))
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
	var m = map[any]any{}
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
			m[str] = constValue.Value
		case gen.GsLexerINT:
			intValue, err := strconv.ParseInt(keyTerminal.GetText(), 0, 64)
			if err != nil {
				panic(err)
			}
			m[int(intValue)] = constValue.Value
		case gen.GsLexerFLOAT:
			floatValue, err := strconv.ParseFloat(keyTerminal.GetText(), 64)
			if err != nil {
				panic(err)
			}
			m[floatValue] = constValue.Value
		case gen.GsLexerTRUE:
			m[true] = constValue.Value
		case gen.GsLexerFALSE:
			m[false] = constValue.Value
		default:
			panic(fmt.Sprintf("unknown key: %s", keyTerminal))
		}
	}
	start := ctx.GetStart()
	name := fmt.Sprintf("%d_%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	for range ctx.GetChildCount() {
		ctx.RemoveLastChild()
	}
	ctx.AddChild(consts.NewConstNode(consts.ConstNodeKindMap, &consts.MapLiteralConst{Map: m, Name: name}, start))
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
	return constValue, true
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
		replaceChildren(ctx, newChildren)
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
		replaceChildren(ctx, newChildren)
	}
	return nil
}

func replaceChildren(ctx antlr.ParserRuleContext, newChildren []antlr.Tree) {
	for range ctx.GetChildCount() {
		ctx.RemoveLastChild()
	}
	for _, child := range newChildren {
		if v, ok := child.(antlr.RuleContext); ok {
			ctx.AddChild(v)
		} else {
			ctx.AddTokenNode(child.(*antlr.TerminalNodeImpl).GetSymbol())
		}
	}
}

func (c *ConstOptimizer) VisitComparisonExpr(ctx *gen.ComparisonExprContext) interface{} {
	c.VisitChildren(ctx)
	//bitExpr (compOp bitExpr)? ;
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
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(pre) == consts.ToStringValue(top), ctx.GetStart())
					} else if top.Kind == consts.ConstNodeKindBool && pre.Kind == consts.ConstNodeKindBool {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToBoolValue(pre) == consts.ToBoolValue(top), ctx.GetStart())
					}
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) == consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) == consts.ToIntValue(top), ctx.GetStart())
					}

				case "<":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) < consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) < consts.ToIntValue(top), ctx.GetStart())
					}
				case ">":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) > consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) > consts.ToIntValue(top), ctx.GetStart())
					}
				case "!=":
					if top.Kind == consts.ConstNodeKindString && pre.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(pre) != consts.ToStringValue(top), ctx.GetStart())
					} else if top.Kind == consts.ConstNodeKindBool && pre.Kind == consts.ConstNodeKindBool {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToBoolValue(pre) != consts.ToBoolValue(top), ctx.GetStart())
					}
					if top.Kind == consts.ConstNodeKindFloat && pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) != consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) != consts.ToIntValue(top), ctx.GetStart())
					}
				case ">=":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) >= consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) >= consts.ToIntValue(top), ctx.GetStart())
					}
				case "<=":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(pre) <= consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(pre) <= consts.ToIntValue(top), ctx.GetStart())
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
		replaceChildren(ctx, newChildren)
	}
	return nil
}

func (c *ConstOptimizer) VisitBinExpr(ctx *gen.BinExprContext) interface{} {
	// addExpr (bitOp addExpr)*
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
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)&consts.ToIntValue(top), ctx.GetStart())
				case "&":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)|consts.ToIntValue(top), ctx.GetStart())
				case "^":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)^consts.ToIntValue(top), ctx.GetStart())
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
		replaceChildren(ctx, newChildren)
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
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(pre)+consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)+consts.ToIntValue(top), ctx.GetStart())
					}

				case "-":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(pre)-consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)-consts.ToIntValue(top), ctx.GetStart())
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
		replaceChildren(ctx, newChildren)
	}
	return nil
}

func (c *ConstOptimizer) VisitMulExpr(ctx *gen.MulExprContext) interface{} {
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
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindFloat, f, ctx.GetStart()))
		case *gen.IntAtomContext:
			applied = true
			f, err := strconv.ParseInt(t.GetText(), 0, 64)
			if err != nil {
				c.Log.ErrorToken(t.GetStart(), "cannot parse float from:%s", t.GetText())
				return nil
			}
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindInt, int(f), ctx.GetStart()))
		case *gen.StringAtomContext:
			applied = true
			literal := t.GetText()
			str, err := strconv.Unquote(literal)
			if err != nil {
				c.Log.ErrorToken(t.GetStart(), fmt.Sprintf("invalid string: %s", literal))
				return nil
			}
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindString, str, ctx.GetStart()))
		case *gen.TrueAtomContext:
			applied = true
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindBool, true, ctx.GetStart()))
		case *gen.FalseAtomContext:
			applied = true
			newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindBool, false, ctx.GetStart()))
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
		case *gen.ArrayAtomContext:
			if al, ok := t.GetChild(0).(*gen.ArrayLiteralContext); ok && t.GetChildCount() == 1 {
				if cn, ok2 := al.GetChild(0).(*consts.ConstNode); ok2 {
					applied = true
					newChildren = append(newChildren, cn)
					break
				}
			}
			newChildren = append(newChildren, t)
		default:
			newChildren = append(newChildren, t)
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
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(pre)*consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)*consts.ToIntValue(top), ctx.GetStart())
					}

				case "/":
					if top.Kind == consts.ConstNodeKindFloat || pre.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(pre)/consts.ToFloatValue(top), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)/consts.ToIntValue(top), ctx.GetStart())
					}
				case "%":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(pre)%consts.ToIntValue(top), ctx.GetStart())
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
		replaceChildren(ctx, newChildren)
	}
	return nil
}

func (c *ConstOptimizer) VisitNegAtom(ctx *gen.NegAtomContext) interface{} {
	atom := ctx.Atom()
	var applied bool
	var newChildren []antlr.Tree
	switch t := atom.(type) {
	case *gen.FloatAtomContext:
		applied = true
		f, err := strconv.ParseFloat(t.GetText(), 64)
		if err != nil {
			c.Log.ErrorToken(t.GetStart(), "cannot parse float from:%s", t.GetText())
			return nil
		}
		newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindFloat, -f, ctx.GetStart()))
	case *gen.IntAtomContext:
		applied = true
		f, err := strconv.ParseInt(t.GetText(), 0, 64)
		if err != nil {
			c.Log.ErrorToken(t.GetStart(), "cannot parse float from:%s", t.GetText())
			return nil
		}
		newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindInt, -int(f), ctx.GetStart()))
	case *gen.ParenAtomContext:
		expr := t.Expr()
		if constNode, ok := toConstNode(expr.(*gen.ExprContext)); ok {
			switch constNode.Kind {
			case consts.ConstNodeKindInt:
				constNode.Value = -constNode.Value.(int)
			case consts.ConstNodeKindFloat:
				constNode.Value = -constNode.Value.(float64)
			default:
				c.Log.ErrorToken(ctx.GetStart(), "can't use '-' on value:%v", constNode.Value)
				return nil
			}
			newChildren = append(newChildren, constNode)
			applied = true
		}
	}
	if applied {
		replaceChildren(ctx, newChildren)
	}
	return nil
}

func (c *ConstOptimizer) VisitInnerCall(ctx *gen.InnerCallContext) interface{} {
	c.VisitChildren(ctx)
	if c.Conf.DefineFuncs.GetFunc("in") != nil && ctx.ID().GetText() == "in" {
		arg1 := ctx.Expr(0)
		if arg1 == nil {
			return nil
		}
		constNode, ok := toConstNode(arg1.(*gen.ExprContext))
		if !ok {
			return nil
		}
		if constNode.Kind == consts.ConstNodeKindList {
			values := constNode.Value.(*consts.SliceLiteralConst).Value
			if len(values) == 0 {
				return nil
			}
			var typ reflect.Kind
			sameType := func(t reflect.Kind) bool {
				if typ == reflect.Invalid {
					typ = t
					return true
				} else {
					return typ == t
				}
			}
			for _, v := range values {
				var curKind reflect.Kind
				switch v := v.(type) {
				case string:
					curKind = reflect.String
				case int:
					curKind = reflect.Int
				case float64:
					curKind = reflect.Float64
				case bool:
					curKind = reflect.Bool
				default:
					panic(fmt.Sprintf("unexpect type %T", v))
				}
				if !sameType(curKind) {
					return nil
				}
			}
			if typ == reflect.Invalid {
				return nil
			}
			var constVal any
			switch typ {
			case reflect.String:
				var constMap = make(map[string]struct{}, len(values))
				for _, v := range values {
					constMap[v.(string)] = struct{}{}
				}
				constVal = constMap
			case reflect.Int:
				var constMap = make(map[int]struct{}, len(values))
				for _, v := range values {
					constMap[v.(int)] = struct{}{}
				}
				constVal = constMap
			case reflect.Float64:
				var constMap = make(map[float64]struct{}, len(values))
				for _, v := range values {
					constMap[v.(float64)] = struct{}{}
				}
				constVal = constMap
			case reflect.Bool:
				var constMap = make(map[bool]struct{}, len(values))
				for _, v := range values {
					constMap[v.(bool)] = struct{}{}
				}
				constVal = constMap
			default:
				panic(fmt.Sprintf("unexpect type %s", typ))
			}
			// remove first expr
			var replaced bool
			var newChildren []antlr.Tree
			for _, tree := range ctx.GetChildren() {
				if _, ok := tree.(gen.IExprContext); ok && !replaced {
					newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindAny, constVal, ctx.GetStart()))
					replaced = true
				} else {
					newChildren = append(newChildren, tree)
				}
			}
			replaceChildren(ctx, newChildren)
		}
	}
	return nil
}
