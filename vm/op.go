package vm

import (
	"github.com/ycl2018/gs/consts"
	"github.com/ycl2018/gs/gen"
)

func (i *Interpreter) Op(op consts.Instr) {
	// 弹出两个操作数，相加，push
	op2, op1 := i.PopOpStack(), i.PopOpStack()
	switch op {
	case consts.InstrAdd:
		v, err := gen.Add(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrSub:
		v, err := gen.Sub(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrMul:
		v, err := gen.Mul(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrDiv:
		v, err := gen.Div(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrLT:
		v, err := gen.Lt(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrEQ:
		v, err := gen.Eq(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrLEQ:
		v, err := gen.Lte(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrNEQ:
		v, err := gen.Neq(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrGEQ:
		v, err := gen.Gte(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrGT:
		v, err := gen.Gt(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrBitOR:
		v, err := gen.Or(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrBitAND:
		v, err := gen.And(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrXOR:
		v, err := gen.Xor(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrMod:
		v, err := gen.Mod(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	default:
		panic("unhandled default case")
	}
}
