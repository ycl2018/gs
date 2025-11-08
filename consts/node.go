package consts

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type ConstKind int

const (
	ConstKindInt = iota
	ConstKindFloat
	ConstKindString
	ConstKindMap
	ConstKindList
	ConstKindBool
)

func ToBoolValue(c *ConstNode) bool {
	switch c.Kind {
	case ConstKindBool:
		return c.Value.(bool)
	default:
		panic(fmt.Sprintf("ToBoolValue: unexpected type %T", c))
	}
}

func ToStringValue(c *ConstNode) string {
	switch c.Kind {
	case ConstKindString:
		return c.Value.(string)
	default:
		panic(fmt.Sprintf("ToStringValue: can't const constant kind: %v to string", c.Kind))
	}
}

func ToFloatValue(c *ConstNode) float64 {
	switch c.Kind {
	case ConstKindInt:
		return float64(c.Value.(int))
	case ConstKindFloat:
		return c.Value.(float64)
	default:
		panic(fmt.Sprintf("ToFloatValue: can't get float value for constant kind: %v", c.Kind))
	}
}

func ToIntValue(c *ConstNode) int {
	switch c.Kind {
	case ConstKindInt:
		return c.Value.(int)
	case ConstKindFloat:
		return int(c.Value.(float64))
	default:
		panic(fmt.Sprintf("ToIntValue: can't get integer value for constant kind: %v", c.Kind))
	}
}

func NewConstNode(kind ConstKind, val any) *ConstNode {
	return &ConstNode{
		Kind:  kind,
		Value: val,
	}
}

type ConstNode struct {
	Kind  ConstKind
	Value any
}

func (c *ConstNode) GetSourceInterval() antlr.Interval {
	//TODO implement me
	panic("implement me")
}

// can't visit
func (c *ConstNode) Accept(Visitor antlr.ParseTreeVisitor) interface{} {
	if v, ok := Visitor.(interface {
		VisitConstNode(v *ConstNode) interface{}
	}); ok {
		return v.VisitConstNode(c)
	}
	return nil
}

func (c *ConstNode) GetText() string {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) ToStringTree(strings []string, recognizer antlr.Recognizer) string {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) GetRuleContext() antlr.RuleContext {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) GetInvokingState() int {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) SetInvokingState(i int) {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) GetRuleIndex() int {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) GetAltNumber() int {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) SetAltNumber(altNumber int) {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) String(strings []string, context antlr.RuleContext) string {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) GetParent() antlr.Tree {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) SetParent(tree antlr.Tree) {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) GetPayload() interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) GetChild(i int) antlr.Tree {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) GetChildCount() int {
	//TODO implement me
	panic("implement me")
}

func (c *ConstNode) GetChildren() []antlr.Tree {
	//TODO implement me
	panic("implement me")
}
