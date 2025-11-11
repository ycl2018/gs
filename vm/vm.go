package vm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ycl2018/gs/consts"
)

type Option func(interpreter *Interpreter)

func WithEnableTrace() Option {
	return func(i *Interpreter) {
		i.enableTrace = true
	}
}

func WithEnableDump() Option {
	return func(i *Interpreter) {
		i.dump = true
	}
}

type Interpreter struct {
	IP           int32               // 指令地址
	Code         []consts.StackInstr // 代码
	ConstPool    []consts.Const
	MainFunc     consts.FunctionConst   // Main函数入口地址
	AllFuncs     []consts.FunctionConst // 所有函数入口地址
	BuildEnvType reflect.Type
	// 函数调用栈
	Calls []*StackFrame
	FP    int32 // 栈桢计数器

	Operands []any // 操作数栈，全局复用
	SP       int32 // 操作数栈计数器

	// 虚拟全局内存
	Globals  []any
	DataSize int

	// ops
	enableTrace bool
	dump        bool
}

const DefaultOperandStackSize = 100

type StackFrame struct {
	ReturnAddr int32               // 返回值地址
	ReturnCode []consts.StackInstr // 之前函数的指令，用于返回时恢复
	FuncConsts *consts.FunctionConst
	Locals     []any // 参数和本地变量
}

func NewStackFrame(f *consts.FunctionConst, returnAddr int32) *StackFrame {
	return &StackFrame{
		ReturnAddr: returnAddr,
		ReturnCode: []consts.StackInstr{},
		FuncConsts: f,
		Locals:     make([]any, f.LocalCount+f.ParamCount),
	}
}

type Code struct {
	Globals      int
	ConstPool    []consts.Const
	MainFunc     consts.FunctionConst
	AllFuncs     []consts.FunctionConst
	BuildEnvType reflect.Type
}

func NewInterpreter(code Code, ops ...Option) *Interpreter {
	// 编译
	i := &Interpreter{
		IP:           -1,
		FP:           -1,
		Operands:     make([]any, DefaultOperandStackSize),
		SP:           -1,
		Globals:      make([]any, code.Globals),
		DataSize:     code.Globals,
		Code:         code.MainFunc.Code,
		ConstPool:    code.ConstPool,
		MainFunc:     code.MainFunc,
		AllFuncs:     code.AllFuncs,
		BuildEnvType: code.BuildEnvType,
	}
	for _, op := range ops {
		op(i)
	}
	return i
}

func (i *Interpreter) Run(env any) error {
	sf := NewStackFrame(&consts.FunctionConst{
		Name: "main",
		Code: i.MainFunc.Code,
	}, i.IP)
	i.Calls = append(i.Calls, sf)
	i.FP++
	i.IP = 0
	if i.enableTrace {
		fmt.Printf("\ntrace:\n")
	}
	i.cpu()

	if i.dump {
		i.Dump()
	}
	return nil
}

func (i *Interpreter) PopOpStack() any {
	v := i.Operands[i.SP]
	i.SP--
	return v
}

func (i *Interpreter) PushOpStack(v any) {
	i.SP++
	i.Operands[i.SP] = v
}

func (i *Interpreter) NextOperand() int32 {
	//val := asm.GetInt(int(i.IP), i.Code)

	i.IP += 4
	return 0
}

func (i *Interpreter) cpu() {
	// 取指令，并执行
	instr := i.Code[i.IP]
	for i.IP < int32(len(i.Code)) && instr.OpCode != consts.InstrHalt {
		if i.enableTrace {
			i.trace()
		}
		i.IP++ // next instruction or first operand
		switch instr.OpCode {
		case consts.InstrAdd:
			// support auto type convert
			// int32 + float32 = int32
			// int32 + string = string
			// float32 + string = string
			op2, op1 := i.PopOpStack(), i.PopOpStack()
			op1Type, op2Type := reflect.TypeOf(op1), reflect.TypeOf(op2)
			if op1Type.Kind() == reflect.String || op2Type.Kind() == reflect.String {
				i.PushOpStack(toString(op1) + toString(op2))
			} else if op1Type.Kind() == reflect.Float32 && op2Type.Kind() == reflect.Float32 {
				i.PushOpStack(toFloat32(op2) + toFloat32(op1))
			} else {
				i.PushOpStack(toInt32(op2) + toInt32(op1))
			}
		case consts.InstrSub, consts.InstrMul, consts.InstrDiv, consts.InstrLT, consts.InstrEQ, consts.InstrLEQ, consts.InstrNEQ, consts.InstrGEQ, consts.InstrGT:
			// 弹出两个操作数，相加，push
			op2, op1 := i.PopOpStack(), i.PopOpStack()
			t2, t1 := reflect.TypeOf(op2), reflect.TypeOf(op1)
			if t1.Kind() == reflect.Float32 || t2.Kind() == reflect.Float32 {
				op1, op2 = toFloat32(op1), toFloat32(op2)
				t1, t2 = reflect.TypeOf(op1), reflect.TypeOf(op2)
			}
			switch instr.OpCode {
			case consts.InstrSub:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) - toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) - toFloat32(op2))
				}
			case consts.InstrMul:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) * toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) * toFloat32(op2))
				}
			case consts.InstrDiv:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) / toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) / toFloat32(op2))
				}
			case consts.InstrLT:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) < toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) < toFloat32(op2))
				}
			case consts.InstrEQ:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) == toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) == toFloat32(op2))
				}
			case consts.InstrLEQ:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) <= toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) <= toFloat32(op2))
				}
			case consts.InstrNEQ:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) != toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) != toFloat32(op2))
				}
			case consts.InstrGEQ:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) >= toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) >= toFloat32(op2))
				}
			case consts.InstrGT:
				if t1.Kind() == reflect.Int32 {
					i.PushOpStack(toInt32(op1) > toInt32(op2))
				} else {
					i.PushOpStack(toFloat32(op1) > toFloat32(op2))
				}
			default:
				panic("unhandled default case")
			}

		case consts.InstrCall:
			// 函数调用
			funcIndex := i.NextOperand()
			fs := i.ConstPool[funcIndex].Value.(consts.FunctionSymbol)
			funcStack := NewStackFrame(&fs, i.IP)
			i.FP++
			i.Calls = append(i.Calls, funcStack)
			// 拷贝操作数到参数中
			// move args from operand stack to top frame on call stack
			for a := int(fs.Args) - 1; a >= 0; a-- {
				funcStack.Locals[a] = i.PopOpStack()
			}
			i.IP = fs.Addr
		case consts.InstrReturn:
			curFs := i.Calls[i.FP]
			i.Calls = i.Calls[:i.FP]
			i.FP--
			i.IP = curFs.ReturnAddr
		case consts.InstrBR:
			toAddr := i.NextOperand()
			i.IP = toAddr
		case consts.InstrBRT:
			toAddr := i.NextOperand()
			if i.PopOpStack().(bool) {
				i.IP = toAddr
			}
		case consts.InstrBRF:
			toAddr := i.NextOperand()
			if !i.PopOpStack().(bool) {
				i.IP = toAddr
			}
		case consts.InstrCConst, consts.InstrIConst:
			val := i.NextOperand()
			i.PushOpStack(val)
		case consts.InstrFConst:
			poolIndex := i.NextOperand()
			fConst := i.ConstPool[poolIndex].Value.(float32)
			i.PushOpStack(fConst)
		case consts.InstrSConst:
			poolIndex := i.NextOperand()
			fConst := i.ConstPool[poolIndex].Value.(string)
			i.PushOpStack(fConst)
		case consts.InstrLoad:
			argIndex := i.NextOperand()
			curStack := i.Calls[i.FP]
			i.PushOpStack(curStack.Locals[argIndex])
		case consts.InstrGLoad:
			argIndex := i.NextOperand()
			gVal := i.Globals[argIndex]
			i.PushOpStack(gVal)
		case consts.InstrFLoad:
			// 字段加载,字段地址在操作数栈中，字段index为下个操作数
			s := i.PopOpStack().(*StructSpace)
			index := i.NextOperand()
			fieldName := i.ConstPool[index].Value.(string)
			i.PushOpStack(s.Field(fieldName))
		case consts.InstrStore:
			addr := i.NextOperand()
			i.Calls[i.FP].Locals[addr] = i.PopOpStack()
		case consts.InstrGStore:
			addr := i.NextOperand()
			i.Globals[addr] = i.PopOpStack()
		case consts.InstrFStore:
			s := i.PopOpStack().(*StructSpace)
			index := i.NextOperand()
			fieldName := i.ConstPool[index].Value.(string)
			s.Set(fieldName, i.PopOpStack())
		case consts.InstrPrint:
			val := i.PopOpStack()
			switch v := val.(type) {
			case *StructSpace:
				fmt.Println(v.String())
			default:
				fmt.Println(v)
			}
		case consts.InstrStruct:
			// push struct
			def := i.ConstPool[int(i.NextOperand())].Value.(consts.ConstStructDef)
			s := NewStructSpace(&def)
			i.PushOpStack(s)
		case consts.InstrNil:
			i.PushOpStack(nil)
		case consts.InstrPop:
			i.PopOpStack()
		case consts.InstrHalt:
			return
		default:
			panic(fmt.Sprintf("unknown opcode:%d", instr))
		}
		instr = i.Code[i.IP]
	}
}

func toInt32(x any) int32 {
	switch x.(type) {
	case int32:
		return x.(int32)
	case float32:
		return int32(x.(float32))
	default:
		return 0
	}
}

func toString(x any) string {
	switch x.(type) {
	case string:
		return x.(string)
	case int32:
		return fmt.Sprintf("%d", x.(int32))
	case float32:
		return fmt.Sprintf("%f", x.(float32))
	default:
		return ""
	}
}

func toFloat32(x any) float32 {
	switch x.(type) {
	case float32:
		return x.(float32)
	case int32:
		return float32(x.(int32))
	default:
		return 0
	}
}

func (i *Interpreter) trace() {
	// asm code
	//_, str, err := i.disAssembler.DisAssembleInstruction(int(i.IP))
	//if err != nil {
	//	fmt.Printf("disAssemble instruction err:%v on ip:%d", err, i.IP)
	//	return
	//}
	//fmt.Println(str)
	//// operand stack
	//fmt.Printf("\tstack=[")
	//for j := int32(0); j <= i.SP; j++ {
	//	fmt.Printf(" %v", i.Operands[j])
	//}
	//fmt.Print(" ]")
	//// call stack
	//if i.FP >= 0 {
	//	fmt.Printf(", calls=[")
	//	for j := int32(0); j <= i.FP; j++ {
	//		fmt.Printf(" " + i.Calls[j].FuncSymbol.Name)
	//	}
	//	fmt.Print(" ]")
	//}
	//fmt.Println()
}

func (i *Interpreter) Dump() {
	if len(i.ConstPool) > 0 {
		i.dumpConstPool()
	}
	if len(i.Globals) > 0 {
		i.dumpDataMemory()
	}
	i.dumpCodeMemory()
}

func (i *Interpreter) dumpConstPool() {
	//fmt.Println("Constant Pool:")
	//dumped, _ := i.disAssembler.DumpConstPool()
	//fmt.Print(dumped)
	//fmt.Println()
}

func (i *Interpreter) dumpCodeMemory() {
	fmt.Println("Code memory:")
	for j := 0; j < len(i.Code); j++ {
		if j%8 == 0 && j != 0 {
			fmt.Println()
		}
		if j%8 == 0 {
			fmt.Printf("%04d:", j)
		}
		fmt.Printf(" %3d", (int)(i.Code[j]))
	}
	fmt.Println()
}

func (i *Interpreter) dumpDataMemory() {
	fmt.Println("Data memory:")
	for j, a := range i.Globals {
		if a != nil {
			fmt.Printf("%04d: %v %s\n", j, a, reflect.TypeOf(a).Name())
		} else {
			fmt.Printf("%04d: null \n", j)
		}
	}
	fmt.Println()
}

type StructSpace struct {
	Name         string
	Fields       map[string]any
	Define       *consts.ConstStructDef
	AllowDynamic bool
}

func (s *StructSpace) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("struct %s { ", s.Name))
	var first = true
	for _, name := range s.Define.MemberNames {
		if !first {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%s: %v", name, s.Fields[name])
		first = false
	}
	if s.AllowDynamic && len(s.Fields) > len(s.Define.MemberNames) {
		// 动态添加？
	OUTER:
		for k, v := range s.Fields {
			for _, n := range s.Define.MemberNames {
				if n == k {
					continue OUTER
				}
			}
			if !first {
				sb.WriteString(", ")
			}
			first = false
			fmt.Fprintf(&sb, "%s: %v\n", k, v)
		}
	}
	sb.WriteString(" }")
	return sb.String()
}

func NewStructSpace(structDef *consts.ConstStructDef) *StructSpace {
	s := &StructSpace{Fields: make(map[string]any, len(structDef.MemberNames))}
	s.Define = structDef
	s.Name = structDef.Name
	return s
}

func (s *StructSpace) Field(name string) any {
	return s.Fields[name]
}

func (s *StructSpace) Set(fieldName string, val any) {
	if s.AllowDynamic {
		s.Fields[fieldName] = val
	} else {
		for name := range s.Fields {
			if name == fieldName {
				s.Fields[name] = val
				return
			}
		}
		panic(fmt.Sprintf("field %s not exists in struct:%s", fieldName, s.Name))
	}
}
