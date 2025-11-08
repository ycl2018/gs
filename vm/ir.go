package vm

import (
	"fmt"
)

type StackInstr struct {
	OpCode   int
	Operands []int
}

func NewStackInstr(opCode int, operands ...int) *StackInstr {
	return &StackInstr{
		OpCode:   opCode,
		Operands: operands,
	}
}

func (s StackInstr) Dump() string {
	if len(s.Operands) > 0 {
		return fmt.Sprintf("%-11s\t%v\n", Instr(s.OpCode), s.Operands[0])
	} else {
		return fmt.Sprintf("%-11s\t\n", Instr(s.OpCode))
	}
}

// ConstKind 常量类型，存储于全局常量池
type ConstKind uint8

const (
	ConstFloat64 ConstKind = iota
	ConstString
	ConstStruct
	ConstFunc
	ConstMapInit
	ConstSliceInit
)

func (c ConstKind) String() string {
	switch c {
	case ConstFloat64:
		return "float64"
	case ConstString:
		return "string"
	case ConstStruct:
		return "struct"
	case ConstFunc:
		return "func"
	case ConstMapInit:
		return "map"
	case ConstSliceInit:
		return "slice"
	default:
		panic(fmt.Sprintf("unknown const kind %d", c))
	}
}

type StructDecl struct {
	Name   string
	Fields []string
}
type FuncDecl struct {
	FuncName string
	NArgs    int32
	NLocals  int32
	Structs  []StructDecl
	BodyCode []StackInstr
}
