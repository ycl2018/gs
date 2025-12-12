package compile

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	consts2 "github.com/ycl2018/gs/consts"
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
	Code       []*consts2.StackInstr
	Debugger   consts2.Debugger
	ReturnNums int
}

func NewFunctionSymbol(funcName string, t antlr.Token) *FunctionSymbol {
	return &FunctionSymbol{
		BaseSymbol: BaseSymbol{
			Name:        funcName,
			DefineToken: t,
		},
		ReturnNums: -1,
	}
}

func (f *FunctionSymbol) Dump() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(".def %s args=%d locals=%d\n", f.Name, len(f.FormalArgs), f.LocalNums()))
	for i, instr := range f.Code {
		sb.WriteString(fmt.Sprintf("#%04d %s\n", i, instr.Dump()))
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

func (f *FunctionSymbol) getParamName() []string {
	var paramNames []string
	for _, arg := range f.FormalArgs {
		paramNames = append(paramNames, arg.GetName())
	}
	return paramNames
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
	Kind           consts2.ConstKind
	Value          any
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
		case consts2.ConstFunc:
			val := s.Value.(consts2.FunctionConst)
			return fmt.Sprintf("#%04d: func %s(args:%d, locals:%d)\n", s.Address, val.Name, val.ParamCount, val.LocalCount)
		case consts2.ConstStruct:
			// 结构体常量
			val := s.Value.(consts2.StructConst)
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("#%04d: struct %s {\n", s.Address, val.Name))
			for _, field := range val.Fields {
				sb.WriteString(fmt.Sprintf("         %s;\n", field))
			}
			sb.WriteString("       }\n")
			return sb.String()
		case consts2.ConstString:
			return fmt.Sprintf("#%04d: string \"%s\"\n", s.Address, consts[s.Address].(*ConstSymbol).Value)
		case consts2.ConstFloat64:
			return fmt.Sprintf("#%04d: float32 %f\n", s.Address, consts[s.Address].(*ConstSymbol).Value)
		case consts2.ConstMapInit:
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("#%04d: map[%s] {\n", s.Address, s.Name))
			m := s.Value.(map[any]any)
			for k, v := range m {
				sb.WriteString(fmt.Sprintf("    %v: %v;\n", k, v))
			}
			sb.WriteString("}\n")
			return sb.String()
		case consts2.ConstSliceInit:
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("#%04d: slice[%s]\n", s.Address, s.Name))
			m := s.Value.([]any)
			bytes, _ := json.Marshal(m)
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

func defineFloatConst(fval float64, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%f", consts2.ConstFloat64, fval),
		Value: fval,
		Kind:  consts2.ConstFloat64,
	}
	symbol, _ := scope.DefineOrGetConst(constSymbol)
	return symbol
}

func defineSliceConst(sliceInit *consts2.SliceLiteralConst, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%s", consts2.ConstSliceInit, sliceInit.Name),
		Value: sliceInit.Value,
		Kind:  consts2.ConstSliceInit,
	}
	symbol, _ := scope.DefineOrGetConst(constSymbol)
	return symbol
}

func defineMapConst(mapInit *consts2.MapLiteralConst, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%s", consts2.ConstMapInit, mapInit.Name),
		Value: mapInit.Map,
		Kind:  consts2.ConstMapInit,
	}
	symbol, _ := scope.DefineOrGetConst(constSymbol)
	return symbol
}

func defineStringConst(val string, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%s", consts2.ConstString, val),
		Kind:  consts2.ConstString,
		Value: val,
	}
	cSymbol, _ := scope.DefineOrGetConst(constSymbol)
	return cSymbol
}

func defineFieldIndexConst(id string, fieldIndex []*reflect.StructField, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%s", consts2.ConstFieldIndex, id),
		Kind:  consts2.ConstFieldIndex,
		Value: fieldIndex,
	}
	cSymbol, _ := scope.DefineOrGetConst(constSymbol)
	return cSymbol
}

func defineFuncConst(name string, paramCount, localCount int, scope *GlobalScope) Symbol {
	val := consts2.FunctionConst{
		Name:       name,
		ParamCount: paramCount,
		LocalCount: localCount,
	}
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%s", consts2.ConstFunc, name),
		Kind:  consts2.ConstFunc,
		Value: val,
	}
	cSymbol, _ := scope.DefineOrGetConst(constSymbol)
	return cSymbol
}

func defineStructConst(constName, name string, fields []string, scope *GlobalScope) Symbol {
	val := consts2.StructConst{
		Name:   name,
		Fields: fields,
	}
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%s", consts2.ConstStruct, constName),
		Kind:  consts2.ConstStruct,
		Value: val,
	}
	cSymbol, _ := scope.DefineOrGetConst(constSymbol)
	return cSymbol
}

func defineAnyConst(name string, val any, scope *GlobalScope) Symbol {
	constSymbol := &ConstSymbol{
		Name:  fmt.Sprintf("%s::%s", consts2.ConstAny, name),
		Kind:  consts2.ConstAny,
		Value: val,
	}
	cSymbol, _ := scope.DefineOrGetConst(constSymbol)
	return cSymbol
}
