package vm

import "github.com/ycl2018/gs/consts"

func (i *Interpreter) Op(op consts.Instr) {
	// 弹出两个操作数，相加，push
	op2, op1 := i.PopOpStack(), i.PopOpStack()
	switch op {
	case consts.InstrAdd:
		v, err := Add(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrSub:
		v, err := Sub(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrMul:
		v, err := Mul(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrDiv:
		v, err := Div(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrLT:
		v, err := Lt(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrEQ:
		v, err := Eq(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrLEQ:
		v, err := Lte(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrNEQ:
		v, err := Neq(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrGEQ:
		v, err := Gte(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrGT:
		v, err := Gt(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrBitOR:
		v, err := Or(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrBitAND:
		v, err := And(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	case consts.InstrXOR:
		v, err := Xor(op1, op2)
		if err != nil {
			panic(err)
		}
		i.PushOpStack(v)
	default:
		panic("unhandled default case")
	}
}
