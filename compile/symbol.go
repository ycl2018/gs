package compile

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	consts2 "github.com/ycl2018/gs/consts"
	"github.com/ycl2018/gs/vm"
)

type Symbol interface {
	GetName() string
	SetScope(scope Scope)
	Scope() Scope
	GetAddress() int
	SetAddress(address int)
}

type BaseSymbol struct {
	Name           string
	EnclosingScope Scope
	Address        int
	DefineToken    antlr.Token
}

func (s *BaseSymbol) SetScope(scope Scope) {
	s.EnclosingScope = scope
}

func (s *BaseSymbol) Scope() Scope {
	return s.EnclosingScope
}

func (s *BaseSymbol) GetAddress() int {
	return s.Address
}

func (s *BaseSymbol) SetAddress(address int) {
	s.Address = address
}

func (s *BaseSymbol) GetName() string {
	return s.Name
}

// FunctionSymbol 函数符号
type FunctionSymbol struct {
	BaseSymbol
	FormalArgs []Symbol
	BodyScope  *LocalScope
	Code       []*vm.StackInstr
	Results    int
	CodeAddr   int // 函数入口
}

func NewFunctionSymbol(funcName string, t antlr.Token) *FunctionSymbol {
	return &FunctionSymbol{
		BaseSymbol: BaseSymbol{
			Name:        funcName,
			DefineToken: t,
		},
	}
}

func (f *FunctionSymbol) Dump() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(".def %s args=%d locals=%d\n", f.Name, len(f.FormalArgs), f.LocalNums()))
	for i, instr := range f.Code {
		sb.WriteString(fmt.Sprintf("#%04d %s", i, instr.Dump()))
	}
	return sb.String()
}

func (f *FunctionSymbol) Define(symbol Symbol) {
	// 检查是否重复定义
	for _, arg := range f.FormalArgs {
		if arg.GetName() == symbol.GetName() {
			panic(fmt.Sprintf("duplicate formal arg %s in function %s line %d",
				symbol.GetName(), f.Name, f.DefineToken.GetLine()))
			return
		}
	}
	// 分配地址
	symbol.SetAddress(len(f.FormalArgs))
	f.FormalArgs = append(f.FormalArgs, symbol)
	symbol.SetScope(f)
}

func (f *FunctionSymbol) Resolve(s string) Symbol {
	for _, arg := range f.FormalArgs {
		if arg.GetName() == s {
			return arg
		}
	}
	if f.EnclosingScope != nil {
		return f.EnclosingScope.Resolve(s)
	}
	return nil
}

func (f *FunctionSymbol) ParentScope() Scope {
	return f.EnclosingScope
}

func (f *FunctionSymbol) LocalNums() int {
	locals := 0
	if f.BodyScope != nil {
		// 过滤 structSymbol
		locals = int(f.BodyScope.LocalVarAllocator)
	}
	return locals
}

// StructSymbol 结构体符号
type StructSymbol struct {
	BaseSymbol
	Fields []Symbol
}

func NewStructSymbol(name string, t antlr.Token) *StructSymbol {
	return &StructSymbol{
		BaseSymbol: BaseSymbol{
			Name:        name,
			DefineToken: t,
		},
	}
}

func (s *StructSymbol) Define(symbol Symbol) {
	var hasField bool
	for i := 0; i < len(s.Fields); i++ {
		if s.Fields[i].GetName() == symbol.GetName() {
			hasField = true
			s.Fields[i] = symbol
		}
	}
	if !hasField {
		s.Fields = append(s.Fields, symbol)
	}
	symbol.SetScope(s)
}

func (s *StructSymbol) Resolve(name string) Symbol {
	for i := 0; i < len(s.Fields); i++ {
		if s.Fields[i].GetName() == name {
			return s.Fields[i]
		}
	}
	if s.EnclosingScope != nil {
		return s.EnclosingScope.Resolve(name)
	}
	return nil
}

func (s *StructSymbol) ParentScope() Scope {
	return s.EnclosingScope
}

// VariableSymbol 变量符号
type VariableSymbol struct {
	BaseSymbol
}

func NewVariableSymbol(name string, t antlr.Token) *VariableSymbol {
	return &VariableSymbol{
		BaseSymbol: BaseSymbol{
			Name:        name,
			DefineToken: t,
		},
	}
}

type FieldSymbol struct {
	BaseSymbol
	Index  int
	Struct *StructSymbol
}

func NewFieldSymbol(ss *StructSymbol, name string, index int, t antlr.Token) *FieldSymbol {
	return &FieldSymbol{
		BaseSymbol: BaseSymbol{Name: name, DefineToken: t},
		Index:      index,
		Struct:     ss,
	}
}

type ConstSymbol struct {
	Name           string
	Address        int // 常量在全局符号表中的地址
	EnclosingScope Scope
	Kind           vm.ConstKind
	Value          any
	// Kind==ConstStruct:structName+存储字段名常量的地址
	// Kind==ConstFunc:存储args,locals的数量，以及函数体代码地址
	Fields []int
}

func (c *ConstSymbol) GetName() string {
	return c.Name
}

func (c *ConstSymbol) SetScope(scope Scope) {
	c.EnclosingScope = scope
}

func (c *ConstSymbol) GetAddress() int {
	return c.Address
}

func (c *ConstSymbol) SetAddress(address int) {
	c.Address = address
}

func (c *ConstSymbol) Scope() Scope {
	return c.EnclosingScope
}

func DumpSymbol(s Symbol, consts []Symbol) string {
	switch s := s.(type) {
	case *ConstSymbol:
		switch s.Kind {
		case vm.ConstFunc:
			return fmt.Sprintf("#%04d: func %s(args:%d, locals:%d)\n", s.Address, s.Name, s.Fields[1], s.Fields[2])
		case vm.ConstStruct:
			// 结构体常量
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("#%04d: struct %s {\n", s.Address, s.Name))
			for i, field := range s.Fields {
				if i == 0 {
					continue // skip name index
				}
				sb.WriteString(fmt.Sprintf("         %s;\n", consts[field].(*ConstSymbol).Value))
			}
			sb.WriteString("       }\n")
			return sb.String()
		case vm.ConstString:
			return fmt.Sprintf("#%04d: string \"%s\"\n", s.Address, consts[s.Address].(*ConstSymbol).Value)
		case vm.ConstFloat64:
			return fmt.Sprintf("#%04d: float32 %f\n", s.Address, consts[s.Address].(*ConstSymbol).Value)
		case vm.ConstMapInit:
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("#%04d: map[%s] {\n", s.Address, s.Name))
			m := s.Value.(*consts2.MapLiteralConst)
			for k, v := range m.Map {
				sb.WriteString(fmt.Sprintf("    %v: %v;\n", k, v))
			}
			sb.WriteString("}\n")
			return sb.String()
		case vm.ConstSliceInit:
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("#%04d: slice[%s]\n", s.Address, s.Name))
			m := s.Value.(*consts2.SliceLiteralConst)
			bytes, _ := json.Marshal(m.Value)
			sb.WriteString("    " + string(bytes) + "\n")
			return sb.String()
		default:
			return fmt.Sprintf("#%04d:\t%-11s\t%-11s\n", s.Address, s.Kind, s.Name)
		}
	case *StructSymbol, *FunctionSymbol:
		return ""
	default:
		return fmt.Sprintf("#%04d:\t%-11s\n", s.GetAddress(), s.GetName())
	}
}
