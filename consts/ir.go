package consts

import (
	"fmt"
)

type StackInstr struct {
	OpCode   Instr
	Operands int
}

func NewStackInstr(opCode Instr, operand int) *StackInstr {
	return &StackInstr{
		OpCode:   opCode,
		Operands: operand,
	}
}

func (s StackInstr) Dump() string {
	switch Instructions[s.OpCode].OpRandType {
	case NIL:
		return fmt.Sprintf("%-11s\t\n", Instr(s.OpCode))
	case INT:
		return fmt.Sprintf("%-11s\t%v\n", Instr(s.OpCode), s.Operands)
	case POLL:
		return fmt.Sprintf("%-11s\tconst#%v\n", Instr(s.OpCode), s.Operands)
	default:
		panic(fmt.Sprintf("unknown op rand type %d", Instructions[s.OpCode].OpRandType))
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
	ConstFieldIndex
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
	case ConstFieldIndex:
		return "fieldIndex"
	default:
		panic(fmt.Sprintf("unknown const kind %d", c))
	}
}
