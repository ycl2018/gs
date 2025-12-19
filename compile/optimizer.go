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
		constValue, ok := toConstNode(tree)
		if !ok {
			return nil
		}
		switch constValue.Kind {
		case consts.ConstNodeKindInt:
			sliceInit = append(sliceInit, consts.ToIntValue(constValue))
		case consts.ConstNodeKindFloat:
			sliceInit = append(sliceInit, consts.ToFloatValue(constValue))
		case consts.ConstNodeKindString:
			sliceInit = append(sliceInit, consts.ToStringValue(constValue))
		case consts.ConstNodeKindBool:
			sliceInit = append(sliceInit, consts.ToBoolValue(constValue))
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
		valueExpr := entry.(antlr.RuleContext).GetChild(2).(gen.IExprContext)
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

func toConstNode(expr gen.IExprContext) (*consts.ConstNode, bool) {
	if expr.GetChildCount() == 0 {
		v, ok := expr.(*consts.ConstNode)
		return v, ok
	}
	if e, ok := expr.(*gen.AtomExprContext); ok {
		return isConstAtom(e.Atom())
	}
	if expr.GetChildCount() != 1 {
		return nil, false
	}
	switch e := expr.GetChild(0).(type) {
	case *consts.ConstNode:
		return e, true
	}
	return nil, false
}

func (c *ConstOptimizer) VisitLogicalOrExpr(ctx *gen.LogicalOrExprContext) interface{} {
	return c.VisitExpr(ctx)
}

func (c *ConstOptimizer) VisitLogicalAndExpr(ctx *gen.LogicalAndExprContext) interface{} {
	return c.VisitExpr(ctx)
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
	return c.VisitExpr(ctx)
}

func (c *ConstOptimizer) VisitAddExpr(ctx *gen.AddExprContext) interface{} {
	return c.VisitExpr(ctx)
}

func (c *ConstOptimizer) VisitMulExpr(ctx *gen.MulExprContext) interface{} {
	return c.VisitExpr(ctx)
}

func (c *ConstOptimizer) VisitExpr(ctx gen.IExprContext) interface{} {
	c.VisitChildren(ctx)
	var newChildren []antlr.Tree
	var applied bool
	var constNodeNum int
	for _, n := range ctx.GetChildren() {
		atomExpr, ok := n.(*gen.AtomExprContext)
		if ok {
			tree := atomExpr.Atom()
			switch t := tree.(type) {
			case *gen.FloatAtomContext:
				f, err := strconv.ParseFloat(t.GetText(), 64)
				if err != nil {
					c.Log.ErrorToken(t.GetStart(), "cannot parse float from:%s", t.GetText())
					return nil
				}
				newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindFloat, f, ctx.GetStart()))
				applied = true
				constNodeNum++
			case *gen.IntAtomContext:
				f, err := strconv.ParseInt(t.GetText(), 0, 64)
				if err != nil {
					c.Log.ErrorToken(t.GetStart(), "cannot parse float from:%s", t.GetText())
					return nil
				}
				newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindInt, int(f), ctx.GetStart()))
				applied = true
				constNodeNum++
			case *gen.StringAtomContext:
				literal := t.GetText()
				str, err := strconv.Unquote(literal)
				if err != nil {
					c.Log.ErrorToken(t.GetStart(), fmt.Sprintf("invalid string: %s", literal))
					return nil
				}
				newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindString, str, ctx.GetStart()))
				applied = true
				constNodeNum++
			case *gen.TrueAtomContext:
				newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindBool, true, ctx.GetStart()))
				applied = true
				constNodeNum++
			case *gen.FalseAtomContext:
				newChildren = append(newChildren, consts.NewConstNode(consts.ConstNodeKindBool, false, ctx.GetStart()))
				applied = true
				constNodeNum++
			case *gen.ParenAtomContext:
				// expr -> logicalor
				expr := t.Expr()
				if constNode, ok := toConstNode(expr.(gen.IExprContext)); ok {
					newChildren = append(newChildren, constNode)
					applied = true
					constNodeNum++
					break
				}
				newChildren = append(newChildren, n)
			case *gen.ArrayAtomContext:
				if al, ok := t.GetChild(0).(*gen.ArrayLiteralContext); ok && t.GetChildCount() == 1 {
					if cn, ok2 := al.GetChild(0).(*consts.ConstNode); ok2 {
						applied = true
						constNodeNum++
						newChildren = append(newChildren, cn)
						break
					}
				}
				newChildren = append(newChildren, n)
			default:
				newChildren = append(newChildren, n)
			}
		} else {
			newChildren = append(newChildren, n)
			if expr, ok := n.(gen.IExprContext); ok {
				if _, ok := toConstNode(expr); ok {
					constNodeNum++
				}
			}

		}
		if constNodeNum > 1 {
			// 合并
			r, ok1 := toConstNode(newChildren[len(newChildren)-1].(gen.IExprContext))
			l, ok2 := toConstNode(newChildren[len(newChildren)-3].(gen.IExprContext))
			if ok1 && ok2 {
				op := newChildren[len(newChildren)-2].(antlr.ParseTree).GetText()
				var constNode *consts.ConstNode
				switch op {
				case "&&":
					for _, v := range []*consts.ConstNode{l, r} {
						val, valid := getBoolConst(v)
						if !valid {
							c.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("syntax err: support value %v in op:%s", v.Value, op))
							return nil
						}
						if valid && !val {
							constNode = v
							break
						}
					}
				case "||":
					for _, v := range []*consts.ConstNode{l, r} {
						val, valid := getBoolConst(v)
						if !valid {
							c.Log.ErrorToken(ctx.GetStart(), fmt.Sprintf("syntax err: support value %v in op:%s", v.Value, op))
							return nil
						}
						if valid && val {
							constNode = v
							break
						}
					}
				case "==":
					if r.Kind == consts.ConstNodeKindString && l.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(l) == consts.ToStringValue(r), ctx.GetStart())
					} else if r.Kind == consts.ConstNodeKindBool && l.Kind == consts.ConstNodeKindBool {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToBoolValue(l) == consts.ToBoolValue(r), ctx.GetStart())
					}
					if r.Kind == consts.ConstNodeKindFloat || l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(l) == consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(l) == consts.ToIntValue(r), ctx.GetStart())
					}
				case "<":
					if r.Kind == consts.ConstNodeKindString && l.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(l) < consts.ToStringValue(r), ctx.GetStart())
					} else if r.Kind == consts.ConstNodeKindFloat || l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(l) < consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(l) < consts.ToIntValue(r), ctx.GetStart())
					}
				case ">":
					if r.Kind == consts.ConstNodeKindString && l.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(l) > consts.ToStringValue(r), ctx.GetStart())
					} else if r.Kind == consts.ConstNodeKindFloat || l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(l) > consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(l) > consts.ToIntValue(r), ctx.GetStart())
					}
				case "!=":
					if r.Kind == consts.ConstNodeKindString && l.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(l) != consts.ToStringValue(r), ctx.GetStart())
					} else if r.Kind == consts.ConstNodeKindBool && l.Kind == consts.ConstNodeKindBool {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToBoolValue(l) != consts.ToBoolValue(r), ctx.GetStart())
					}
					if r.Kind == consts.ConstNodeKindFloat && l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(l) != consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(l) != consts.ToIntValue(r), ctx.GetStart())
					}
				case ">=":
					if r.Kind == consts.ConstNodeKindString && l.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(l) >= consts.ToStringValue(r), ctx.GetStart())
					} else if r.Kind == consts.ConstNodeKindFloat || l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(l) >= consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(l) >= consts.ToIntValue(r), ctx.GetStart())
					}
				case "<=":
					if r.Kind == consts.ConstNodeKindString && l.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToStringValue(l) <= consts.ToStringValue(r), ctx.GetStart())
					} else if r.Kind == consts.ConstNodeKindFloat || l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToFloatValue(l) <= consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindBool, consts.ToIntValue(l) <= consts.ToIntValue(r), ctx.GetStart())
					}
				case "+":
					if r.Kind == consts.ConstNodeKindString || l.Kind == consts.ConstNodeKindString {
						constNode = consts.NewConstNode(consts.ConstNodeKindString, consts.ToStringValue(l)+consts.ToStringValue(r), ctx.GetStart())
					} else if r.Kind == consts.ConstNodeKindFloat || l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(l)+consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)+consts.ToIntValue(r), ctx.GetStart())
					}
				case "-":
					if r.Kind == consts.ConstNodeKindFloat || l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(l)-consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)-consts.ToIntValue(r), ctx.GetStart())
					}
				case "*":
					if r.Kind == consts.ConstNodeKindFloat || l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(l)*consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)*consts.ToIntValue(r), ctx.GetStart())
					}
				case "&":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)&consts.ToIntValue(r), ctx.GetStart())
				case "|":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)|consts.ToIntValue(r), ctx.GetStart())
				case "^":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)^consts.ToIntValue(r), ctx.GetStart())
				case "<<":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)<<consts.ToIntValue(r), ctx.GetStart())
				case ">>":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)>>consts.ToIntValue(r), ctx.GetStart())
				case "/":
					if r.Kind == consts.ConstNodeKindFloat || l.Kind == consts.ConstNodeKindFloat {
						constNode = consts.NewConstNode(consts.ConstNodeKindFloat, consts.ToFloatValue(l)/consts.ToFloatValue(r), ctx.GetStart())
					} else {
						constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)/consts.ToIntValue(r), ctx.GetStart())
					}
				case "%":
					constNode = consts.NewConstNode(consts.ConstNodeKindInt, consts.ToIntValue(l)%consts.ToIntValue(r), ctx.GetStart())
				default:
					panic(fmt.Sprintf("unknown op: %v", op))
				}
				newChildren = newChildren[:len(newChildren)-3]
				newChildren = append(newChildren, constNode)
				applied = true
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
		t.Expr().Accept(c)
		expr := t.Expr()
		if constNode, ok := toConstNode(expr.(gen.IExprContext)); ok {
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

func isConstAtom(ctx gen.IAtomContext) (*consts.ConstNode, bool) {
	switch t := ctx.(type) {
	case *gen.FloatAtomContext:
		f, err := strconv.ParseFloat(t.GetText(), 64)
		if err != nil {
			return nil, false
		}
		return consts.NewConstNode(consts.ConstNodeKindFloat, f, ctx.GetStart()), true
	case *gen.IntAtomContext:
		f, err := strconv.ParseInt(t.GetText(), 0, 64)
		if err != nil {
			return nil, false
		}
		return consts.NewConstNode(consts.ConstNodeKindInt, int(f), ctx.GetStart()), true
	case *gen.StringAtomContext:
		return consts.NewConstNode(consts.ConstNodeKindString, t.GetText()[1:len(t.GetText())-1], ctx.GetStart()), true
	default:
		return nil, false
	}
}

func (c *ConstOptimizer) VisitInnerCall(ctx *gen.InnerCallContext) interface{} {
	c.VisitChildren(ctx)
	if c.Conf.DefineFuncs.GetFunc("in") != nil && ctx.ID().GetText() == "in" {
		arg1 := ctx.Expr(0)
		if arg1 == nil {
			return nil
		}
		atom, ok := arg1.(*gen.AtomExprContext)
		if !ok {
			return nil
		}
		arrayAtom, ok := atom.GetChild(0).(*gen.ArrayAtomContext)
		if !ok {
			return nil
		}
		arrayLiteral := arrayAtom.ArrayLiteral()
		if len(arrayLiteral.GetChildren()) == 0 {
			return nil
		}
		constNode, ok := arrayLiteral.GetChild(0).(*consts.ConstNode)
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

func getBoolConst(constNode *consts.ConstNode) (bool, bool) {
	if constNode.Kind == consts.ConstNodeKindBool {
		return constNode.Value.(bool), true
	}
	return false, false
}
