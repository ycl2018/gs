package vm

import (
	"fmt"
	"math"
	"reflect"
	"slices"
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

func WithEnv(env any) Option {
	return func(i *Interpreter) {
		i.Env = env
	}
}

type Interpreter struct {
	IP           int                 // 指令地址
	Code         []consts.StackInstr // 代码
	ConstPool    []consts.Const
	MainFunc     consts.FunctionConst // Main函数入口地址
	BuildEnvType reflect.Type
	// 函数调用栈
	Calls []*StackFrame
	FP    int // 栈桢计数器

	Operands []any // 操作数栈，全局复用
	SP       int   // 操作数栈计数器

	// 虚拟全局内存
	Globals  []any
	DataSize int
	Env      any

	// ops
	enableTrace bool
	dump        bool
}

const DefaultOperandStackSize = 100

type StackFrame struct {
	ReturnAddr int                 // 返回值地址
	ReturnCode []consts.StackInstr // 之前函数的指令，用于返回时恢复
	FuncConsts *consts.FunctionConst
	Locals     []any // 参数和本地变量
}

func NewStackFrame(f *consts.FunctionConst, returnAddr int, code []consts.StackInstr) *StackFrame {
	return &StackFrame{
		ReturnAddr: returnAddr,
		ReturnCode: code,
		FuncConsts: f,
		Locals:     make([]any, f.LocalCount+f.ParamCount),
	}
}

type Code struct {
	Globals      int
	ConstPool    []consts.Const
	MainFunc     consts.FunctionConst
	BuildEnvType reflect.Type
}

func NewInterpreter(code *Code, ops ...Option) *Interpreter {
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
		BuildEnvType: code.BuildEnvType,
	}
	for _, op := range ops {
		op(i)
	}
	return i
}

func (i *Interpreter) Run() {
	i.IP = 0
	sf := NewStackFrame(&consts.FunctionConst{
		Name: "main",
		Code: i.MainFunc.Code,
	}, i.IP, i.Code)
	i.Calls = append(i.Calls, sf)
	i.FP++
	if i.enableTrace {
		fmt.Printf("\ntrace:\n")
	}
	i.cpu()

	if i.dump {
		i.Dump()
	}
	return
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

func (i *Interpreter) cpu() {
	// 取指令，并执行
	instr := i.Code[i.IP]
	for i.IP < len(i.Code) && instr.OpCode != consts.InstrHalt {
		if i.enableTrace {
			i.trace()
		}
		i.IP++ // next instruction or first operand
		switch instr.OpCode {
		case consts.InstrAdd, consts.InstrSub, consts.InstrMul, consts.InstrDiv, consts.InstrLT, consts.InstrEQ, consts.InstrLEQ, consts.InstrNEQ, consts.InstrGEQ, consts.InstrGT, consts.InstrBitOR, consts.InstrBitAND, consts.InstrXOR:
			i.Op(instr.OpCode)
		case consts.InstrOR:
			i.PushOpStack(i.PopOpStack().(bool) || i.PopOpStack().(bool))
		case consts.InstrAND:
			i.PushOpStack(i.PopOpStack().(bool) && i.PopOpStack().(bool))
		case consts.InstrPow:
			op2, op1 := i.PopOpStack(), i.PopOpStack()
			i.PushOpStack(math.Pow(toFloat64(op1), toFloat64(op2)))
		case consts.InstrNeg:
			i.PushOpStack(neg(i.PopOpStack()))
		case consts.InstrTrue:
			i.PushOpStack(true)
		case consts.InstrFalse:
			i.PushOpStack(false)
		case consts.InstrNot:
			i.PushOpStack(!i.PopOpStack().(bool))
		case consts.InstrArray:
			arrLen := instr.Operands
			arr := make([]any, arrLen)
			for i2 := int(arrLen) - 1; i2 >= 0; i2-- {
				arr[i2] = i.PopOpStack()
			}
			i.PushOpStack(arr)
		case consts.InstrIndexLoad:
			index, obj := i.PopOpStack(), i.PopOpStack()
			i.PushOpStack(i.Index(obj, index))
		case consts.InstrSliceSplit:
			end, start := i.PopOpStack(), i.PopOpStack()
			obj := i.PopOpStack()
			i.PushOpStack(i.SplitSlice(obj, start, end))
		case consts.InstrDict:
			dictLen := instr.Operands
			i.PushOpStack(i.MakeMap(dictLen))
		case consts.InstrCall:
			// 函数调用
			funcIndex := instr.Operands
			fs := i.ConstPool[funcIndex].Value.(consts.FunctionConst)
			funcStack := NewStackFrame(&fs, i.IP, i.Code)
			i.FP++
			i.Calls = append(i.Calls, funcStack)
			// 拷贝操作数到参数中
			// move args from operand stack to top frame on call stack
			for a := int(fs.ParamCount) - 1; a >= 0; a-- {
				funcStack.Locals[a] = i.PopOpStack()
			}
			i.IP, i.Code = 0, fs.Code
		case consts.InstrReturn:
			if i.FP < 0 {
				// main return
				return
			}
			curFs := i.Calls[i.FP]
			i.Calls = i.Calls[:i.FP]
			i.FP--
			i.IP, i.Code = curFs.ReturnAddr, curFs.ReturnCode
		case consts.InstrBR:
			toAddr := instr.Operands
			i.IP = toAddr
		case consts.InstrBRT:
			toAddr := instr.Operands
			if i.PopOpStack().(bool) {
				i.IP = toAddr
			}
		case consts.InstrBRF:
			toAddr := instr.Operands
			if !i.PopOpStack().(bool) {
				i.IP = toAddr
			}
		case consts.InstrBRNil:
			obj := i.Peek()
			if obj == nil {
				i.IP = instr.Operands
			}
		case consts.InstrCConst, consts.InstrIConst:
			i.PushOpStack(instr.Operands)
		case consts.InstrFConst:
			poolIndex := instr.Operands
			fConst := i.ConstPool[poolIndex].Value.(float64)
			i.PushOpStack(fConst)
		case consts.InstrSConst:
			poolIndex := instr.Operands
			fConst := i.ConstPool[poolIndex].Value.(string)
			i.PushOpStack(fConst)
		case consts.InstrSliceConst:
			poolIndex := instr.Operands
			sliceConst := i.ConstPool[poolIndex].Value.([]any)
			i.PushOpStack(sliceConst)
		case consts.InstrMapConst:
			poolIndex := instr.Operands
			mapConst := i.ConstPool[poolIndex].Value.(map[consts.ConstNode]*consts.ConstNode)
			i.PushOpStack(mapConst)
		case consts.InstrNil:
			i.PushOpStack(nil)
		case consts.InstrLoad:
			argIndex := instr.Operands
			curStack := i.Calls[i.FP]
			i.PushOpStack(curStack.Locals[argIndex])
		case consts.InstrGLoad:
			argIndex := instr.Operands
			gVal := i.Globals[argIndex]
			i.PushOpStack(gVal)
		case consts.InstrFLoad:
			// 字段加载,字段地址在操作数栈中，字段index为下个操作数
			index := instr.Operands
			i.PushOpStack(i.FieldLoad(i.ConstPool[index].Value.(string)))
		case consts.InstrStore:
			addr := instr.Operands
			i.Calls[i.FP].Locals[addr] = i.PopOpStack()
		case consts.InstrGStore:
			addr := instr.Operands
			i.Globals[addr] = i.PopOpStack()
		case consts.InstrFStore:
			index := instr.Operands
			i.FieldStore(i.ConstPool[index].Value.(string))
		case consts.InstrIndexStore:
			i.IndexStore()
		case consts.InstrPrint:
			printNums := instr.Operands
			for i2 := 0; i2 < printNums; i2++ {
				fmt.Print(i.PopOpStack())
			}
			fmt.Printf("\n")
		case consts.InstrStruct:
			// push struct
			def := i.ConstPool[instr.Operands].Value.(consts.ConstStructDef)
			s := NewStructSpace(&def)
			i.PushOpStack(s)
		case consts.InstrPop:
			i.PopOpStack()
		case consts.InstrBuildTuple:
			i.PushOpStack(i.BuildTuple(instr.Operands))
		case consts.InstrUnpack:
			t := i.PopOpStack().(consts.Tuple)
			num := instr.Operands
			if num != t.Num {
				panic(fmt.Sprintf("unpack tuple %d items to %d variables", t.Num, num))
			}
			for i2 := num - 1; i2 >= 0; i2-- {
				i.PushOpStack(t.Values[i2])
			}
		case consts.InstrIter:
			i.PushOpStack(i.Iter(i.PopOpStack()))
		case consts.InstrIterNext:
			iterNum := instr.Operands
			iter := i.PopOpStack().(consts.Iter)
			iter1, iter2 := iter.Next()
			if iterNum == 1 {
				i.PushOpStack(iter1)
			} else {
				i.PushOpStack(iter2)
			}
		case consts.InstrIterDone:
			i.PushOpStack(i.Peek().(consts.Iter).Done())
		case consts.InstrHalt:
			return
		case consts.InstrLoadEnv:
			i.PushOpStack(i.Env)
		case consts.InstrRFByIndex:
			i.PushOpStack(i.FieldByIndex(i.ConstPool[instr.Operands].Value.([]*reflect.StructField)))
		case consts.InstrRSetField:
			i.RSetField(i.ConstPool[instr.Operands].Value.([]*reflect.StructField))
		case consts.InstrRMapIndex:
			i.PushOpStack(i.MapIndex(i.PopOpStack(), i.PopOpStack()))
		case consts.InstrRIndex:
			i.PushOpStack(i.RIndex(i.PopOpStack(), i.PopOpStack()))
		case consts.InstrRSet:
			i.RSet(i.PopOpStack(), i.PopOpStack())
		case consts.InstrRSetMapIndex:
			i.RSetMapIndex(i.PopOpStack(), i.PopOpStack(), i.PopOpStack())
		default:
			panic(fmt.Sprintf("unknown opcode:%d", instr))
		}
		instr = i.Code[i.IP]
	}
}

func toInt(v any) int {
	switch v := v.(type) {
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case int:
		return int(v)
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	}
	panic(fmt.Sprintf("unexpected type %T for conversion to float64", v))
}

func toFloat64(v any) float64 {
	switch v := v.(type) {
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return float64(v)
	}
	panic(fmt.Sprintf("unexpected type %T for conversion to float64", v))
}

func neg(v any) any {
	switch v := v.(type) {
	case uint:
		return -(v)
	case uint8:
		return -(v)
	case uint16:
		return -(v)
	case uint32:
		return -(v)
	case uint64:
		return -(v)
	case int:
		return -(v)
	case int8:
		return -(v)
	case int16:
		return -(v)
	case int32:
		return -(v)
	case int64:
		return -(v)
	case float32:
		return -(v)
	case float64:
		return -(v)
	}
	panic(fmt.Sprintf("unexpected type %T for negation", v))
}

func (i *Interpreter) trace() {
	// asm code
	fmt.Println(i.Code[i.IP])
	// operand stack
	fmt.Printf("\tstack=[")
	for j := (0); j <= i.SP; j++ {
		fmt.Printf(" %v", i.Operands[j])
	}
	fmt.Print(" ]")
	// call stack
	if i.FP >= 0 {
		fmt.Printf(", calls=[")
		for j := (0); j <= i.FP; j++ {
			fmt.Printf(" " + i.Calls[j].FuncConsts.Name)
		}
		fmt.Print(" ]")
	}
	fmt.Println()
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
		fmt.Printf(" %3s", i.Code[j].OpCode.String())
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

func (i *Interpreter) Index(obj any, index any) any {
	switch obj := obj.(type) {
	case []any:
		return obj[toInt(index)]
	case map[any]any:
		return obj[index]
	case map[string]any:
		return obj[index.(string)]
	default:
		rv := assertValidObj(obj)
		switch rv.Kind() {
		case reflect.Slice, reflect.Array:
			return rv.Index(toInt(index)).Interface()
		case reflect.Map:
			return rv.MapIndex(reflect.ValueOf(index)).Interface()
		default:
			panic(fmt.Sprintf("unexpected type %T for index", obj))
		}
	}
}

func assertValidObj(obj any) reflect.Value {
	if obj == nil {
		panic("obj is nil")
	}
	rv := reflect.ValueOf(obj)
	if rv.Kind() == reflect.Invalid {
		panic("obj is invalid")
	}
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		if rv.Kind() == reflect.Invalid {
			panic("obj is invalid")
		}
	}
	return rv
}

func (i *Interpreter) SplitSlice(obj any, start any, end any) any {
	rv := assertValidObj(obj)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		return rv.Slice(toInt(start), toInt(end)).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for slice split", obj))
	}
}

func (i *Interpreter) MakeMap(dictLen int) any {
	m := make(map[any]any, dictLen)
	for i2 := 0; i2 < dictLen; i2++ {
		t := i.PopOpStack().(consts.Tuple)
		if t.Num != 2 {
			panic(fmt.Sprintf("unexpected tuple num %d for map init", t.Num))
		}
		m[t.Values[0]] = t.Values[1]
	}
	return m
}

func (i *Interpreter) Peek() any {
	return i.Operands[i.SP]
}

func (i *Interpreter) FieldLoad(field string) any {
	// build-in type
	if structSpace, ok := i.PopOpStack().(*StructSpace); ok {
		return structSpace.Fields[field]
	}
	// reflect
	obj := reflect.ValueOf(i.PopOpStack())
	assertValidObj(obj)
	switch obj.Kind() {
	case reflect.Struct:
		return obj.FieldByName(field).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for field load", field))
	}
}

func (i *Interpreter) FieldStore(field string) any {
	// build-in type
	obj := i.PopOpStack()
	if structSpace, ok := obj.(*StructSpace); ok {
		return structSpace.Fields[field]
	}
	// reflect
	objStruct := assertValidObj(obj)
	switch objStruct.Kind() {
	case reflect.Struct:
		return objStruct.FieldByName(field).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for field load", field))
	}
}

func (i *Interpreter) IndexStore() {
	val, index, obj := i.PopOpStack(), i.PopOpStack(), i.PopOpStack()
	rv := assertValidObj(obj)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		rv.Index(toInt(index)).Set(reflect.ValueOf(val))
	case reflect.Map:
		rv.MapIndex(reflect.ValueOf(index)).Set(reflect.ValueOf(val))
	default:
		panic(fmt.Sprintf("unexpected type %T for index store", obj))
	}
}

func (i *Interpreter) BuildTuple(operands int) any {
	var ret []any
	unpack := func(t consts.Tuple) {
		for i2 := t.Num - 1; i2 >= 0; i2-- {
			ret = append(ret, t.Values[i2])
		}
	}
	for i2 := operands - 1; i2 >= 0; i2-- {
		val := i.PopOpStack()
		if t, ok := val.(consts.Tuple); ok {
			unpack(t)
		} else {
			ret = append(ret, val)
		}
	}
	slices.Reverse(ret)
	return consts.Tuple{
		Values: ret,
		Num:    len(ret),
	}
}

func (i *Interpreter) Iter(obj any) any {
	// map/slices/int/array
	rv := assertValidObj(obj)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		return consts.Iter{
			Obj:  rv,
			Len:  rv.Len(),
			Kind: "slice",
		}
	case reflect.Map:
		return consts.Iter{
			Obj:  rv.MapRange(),
			Len:  rv.Len(),
			Kind: "map",
		}
	case reflect.Int:
		return consts.Iter{
			Obj:  rv,
			Len:  int(rv.Int()),
			Kind: "int",
		}
	default:
		panic(fmt.Sprintf("unexpected type %T for iter", obj))
	}

}

func (i *Interpreter) FieldByIndex(fields []*reflect.StructField) any {
	obj := i.PopOpStack()
	rv := assertValidObj(obj)
	var index []int
	for _, f := range fields {
		index = append(index, f.Index...)
	}
	value, err := rv.FieldByIndexErr(index)
	if err != nil {
		panic("null pointer")
	}
	return value.Interface()
}

func (i *Interpreter) RSetField(fields []*reflect.StructField) {
	value, rv := i.PopOpStack(), assertValidObj(i.PopOpStack())
	var index []int
	for _, f := range fields {
		index = append(index, f.Index...)
	}
	obj, err := rv.FieldByIndexErr(index)
	if err != nil {
		panic("null pointer")
	}
	obj.Set(reflect.ValueOf(value))
}

func (i *Interpreter) MapIndex(key any, m any) any {
	rv := assertValidObj(m)
	switch rv.Kind() {
	case reflect.Map:
		return rv.MapIndex(reflect.ValueOf(key)).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for map index", m))
	}
}

func (i *Interpreter) RIndex(index any, slice any) any {
	rv := assertValidObj(slice)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		return rv.Index(toInt(index)).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for slice index", slice))
	}
}

func (i *Interpreter) RSet(val any, obj any) {
	rv := assertValidObj(obj)
	if rv.CanSet() {
		rv.Set(reflect.ValueOf(val))
		return
	}
	panic(fmt.Sprintf("unexpected type %T for Rset", obj))
}

func (i *Interpreter) RSetMapIndex(k, m, val any) {
	rv := assertValidObj(m)
	switch rv.Kind() {
	case reflect.Map:
		rv.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(val))
	default:
		panic(fmt.Sprintf("unexpected type %T for map index store", m))
	}
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
