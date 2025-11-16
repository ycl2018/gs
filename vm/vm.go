package vm

import (
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"slices"
	"strings"
	"unsafe"

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

func WithPrintTo(writer io.Writer) Option {
	return func(i *Interpreter) {
		i.Out = writer
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
	Out      io.Writer

	// ops
	enableTrace     bool
	dump            bool
	initialStackCap int
}

const DefaultOperandStackSize = 256

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
		SP:           -1,
		Globals:      make([]any, code.Globals),
		DataSize:     code.Globals,
		Code:         code.MainFunc.Code,
		ConstPool:    code.ConstPool,
		MainFunc:     code.MainFunc,
		BuildEnvType: code.BuildEnvType,
		Out:          os.Stdout,
	}
	for _, op := range ops {
		op(i)
	}
	if i.initialStackCap <= 0 {
		i.initialStackCap = DefaultOperandStackSize
	}
	i.Operands = make([]any, i.initialStackCap)
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
	if i.SP >= i.initialStackCap {
		i.Operands = append(i.Operands, v)
		return
	}
	i.Operands[i.SP] = v
}

func FastGrowSlice(slice []any, newCapacity int) []any {
	if newCapacity <= cap(slice) {
		return slice[:newCapacity]
	}

	// 创建新 slice
	newSlice := make([]any, len(slice), newCapacity)

	// 使用 unsafe 进行内存复制（高性能，但危险）
	if len(slice) > 0 {
		src := unsafe.SliceData(slice)
		dst := unsafe.SliceData(newSlice)
		n := copy(unsafe.Slice(dst, len(slice)), unsafe.Slice(src, len(slice)))
		_ = n // 使用 n 避免编译警告
	}

	return newSlice
}

func (i *Interpreter) cpu() {
	// 取指令，并执行
	instr := i.Code[i.IP]
	for i.IP < len(i.Code) {
		if i.enableTrace {
			i.trace()
		}
		i.IP++ // next instruction or first operand
		switch instr.OpCode {
		case consts.InstrAdd, consts.InstrSub, consts.InstrMul, consts.InstrDiv, consts.InstrLT, consts.InstrEQ, consts.InstrLEQ, consts.InstrNEQ, consts.InstrGEQ, consts.InstrGT, consts.InstrMod, consts.InstrBitOR, consts.InstrBitAND, consts.InstrXOR:
			i.Op(instr.OpCode)
		case consts.InstrOR:
			i.PushOpStack(i.PopOpStack().(bool) || i.PopOpStack().(bool))
		case consts.InstrAND:
			i.PushOpStack(i.PopOpStack().(bool) && i.PopOpStack().(bool))
		case consts.InstrPow:
			op2, op1 := i.PopOpStack(), i.PopOpStack()
			i.PushOpStack(math.Pow(ToFloat64(op1), ToFloat64(op2)))
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
			var copied = make([]any, len(sliceConst))
			copy(copied, sliceConst)
			i.PushOpStack(copied) // copied
		case consts.InstrMapConst:
			poolIndex := instr.Operands
			mapConst := i.ConstPool[poolIndex].Value.(map[any]any)
			copied := make(map[any]any, len(mapConst))
			for k, v := range mapConst {
				copied[k] = v
			}
			i.PushOpStack(copied)
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
			i.IndexStore(i.PopOpStack(), i.PopOpStack(), i.PopOpStack())
		case consts.InstrPrint:
			printNums := instr.Operands
			var toPrint = make([]any, printNums)
			for i2 := printNums - 1; i2 >= 0; i2-- {
				toPrint[i2] = i.PopOpStack()
			}
			fmt.Fprint(i.Out, toPrint...)
		case consts.InstrPrintf:
			printNums := instr.Operands
			var toPrint = make([]any, printNums)
			for i2 := printNums - 1; i2 >= 0; i2-- {
				toPrint[i2] = i.PopOpStack()
			}
			fmtStr, ok := toPrint[0].(string)
			if !ok {
				panic(fmt.Sprintf("invalid type:%T: first printf args must be a string", toPrint[0]))
			}
			fmt.Fprintf(i.Out, fmtStr, toPrint[1:])
		case consts.InstrPrintln:
			printNums := instr.Operands
			var toPrint = make([]any, printNums)
			for i2 := printNums - 1; i2 >= 0; i2-- {
				toPrint[i2] = i.PopOpStack()
			}
			fmt.Fprintln(i.Out, toPrint...)
		case consts.InstrLen:
			i.PushOpStack(length(i.PopOpStack()))
		case consts.InstrAppend:
			appendNums := instr.Operands
			var appendVals = make([]any, appendNums)
			for i2 := appendNums - 1; i2 >= 0; i2-- {
				appendVals[i2] = i.PopOpStack()
			}
			slice, vals := appendVals[0], appendVals[1:]
			i.PushOpStack(appendSlice(slice, vals))
		case consts.InstrDelete:
			key := i.PopOpStack()
			m := i.PopOpStack()
			deleteMap(m, key)
		case consts.InstrCopy:
			val := i.PopOpStack()
			i.PushOpStack(copySlice(val))
		case consts.InstrToString:
			i.PushOpStack(toString(i.PopOpStack()))
		case consts.InstrConvert:
			i.PushOpStack(convert(i.PopOpStack(), reflect.Kind(instr.Operands)))
		case consts.InstrStruct:
			// push struct
			def := i.ConstPool[instr.Operands].Value.(consts.StructConst)
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
			iter := i.Peek().(*consts.Iter)
			var iter1, iter2 any
			if iterNum == 1 {
				iter1 = iter.Next1()
			} else {
				iter1, iter2 = iter.Next()
			}
			if iterNum == 1 {
				i.PushOpStack(iter1)
			} else {
				i.PushOpStack(iter1)
				i.PushOpStack(iter2)
			}
		case consts.InstrIterDone:
			c := i.Peek().(*consts.Iter)
			i.PushOpStack(c.Done())
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
		case consts.InstrRIndexStore:
			i.RIndexStore(i.PopOpStack(), i.PopOpStack(), i.PopOpStack())
		case consts.InstrRSet:
			i.RSet(i.PopOpStack(), i.PopOpStack())
		case consts.InstrRSetMapIndex:
			i.RSetMapIndex(i.PopOpStack(), i.PopOpStack(), i.PopOpStack())
		case consts.InstrDeref:
			i.PushOpStack(i.Deref(i.PopOpStack()))
		case consts.InstrNewPtrValue:
			i.PushOpStack(i.NewPtrValue(i.PopOpStack()))
		case consts.InstrHalt:
			return
		default:
			panic(fmt.Sprintf("unknown opcode:%d", instr))
		}
		instr = i.Code[i.IP]
	}
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
		fmt.Print(" ]\n")
	}
	fmt.Println()
	fmt.Print(i.Code[i.IP])
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
		return obj[ToInt(index)]
	case map[any]any:
		return obj[index]
	case map[string]any:
		return obj[index.(string)]
	default:
		rv := assertValidObj(obj)
		switch rv.Kind() {
		case reflect.Slice, reflect.Array:
			return rv.Index(ToInt(index)).Interface()
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
		start, end := ToInt(start), ToInt(end)
		if start == -1 {
			start = 0
		}
		if end == -1 {
			end = rv.Len()
		}
		return rv.Slice(start, end).Interface()
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

func (i *Interpreter) FieldStore(field string) {
	// build-in type
	obj := i.PopOpStack()
	val := i.PopOpStack()
	if structSpace, ok := obj.(*StructSpace); ok {
		structSpace.Fields[field] = val
		return
	}
	// reflect
	objStruct := assertValidObj(obj)
	switch objStruct.Kind() {
	case reflect.Struct:
		objStruct.FieldByName(field).Set(reflect.ValueOf(val))
	default:
		panic(fmt.Sprintf("unexpected type %T for field load", field))
	}
}

func (i *Interpreter) IndexStore(index, obj, val any) {
	switch obj := obj.(type) {
	case map[any]any:
		obj[index] = val
		return
	case []any:
		obj[ToInt(index)] = val
		return
	case []string:
		obj[ToInt(index)] = val.(string)
	case []int:
		obj[ToInt(index)] = ToInt(val)
	case []int64:
		obj[ToInt(index)] = ToInt64(val)
	case map[string]any:
		obj[index.(string)] = val
		return
	case map[string]bool:
		obj[index.(string)] = val.(bool)
	case map[string]int:
		obj[index.(string)] = ToInt(val)
	case map[string]string:
		obj[index.(string)] = val.(string)
	case map[int]any:
		obj[index.(int)] = val
	case map[int]bool:
		obj[index.(int)] = val.(bool)
	case map[int]string:
		obj[index.(int)] = val.(string)
	case map[int64]any:
		obj[index.(int64)] = val
	case map[int64]bool:
		obj[index.(int64)] = val.(bool)
	}
	rv := assertValidObj(obj)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		rv.Index(ToInt(index)).Set(reflect.ValueOf(val))
	case reflect.Map:
		rv.SetMapIndex(reflect.ValueOf(index), reflect.ValueOf(val))
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
		return &consts.Iter{
			Obj:  rv,
			Len:  rv.Len(),
			Kind: "slice",
		}
	case reflect.Map:
		return &consts.Iter{
			Obj:  rv.MapRange(),
			Len:  rv.Len(),
			Kind: "map",
		}
	case reflect.Int:
		return &consts.Iter{
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
	obj := assertValidObj(i.PopOpStack())
	value := i.PopOpStack()
	var index []int
	for _, f := range fields {
		index = append(index, f.Index...)
	}
	fieldObj, err := obj.FieldByIndexErr(index)
	if err != nil {
		panic("null pointer")
	}
	lastFiled := fields[len(fields)-1]
	SetField(fieldObj, lastFiled.Type, value)
}

// SetField 设置结构体字段的值，支持类型转换
func SetField(fieldObj reflect.Value, fieldType reflect.Type, value any) {
	switch fieldType.Kind() {
	case reflect.Int:
		fieldObj.Set(reflect.ValueOf(ToInt(value)))
	case reflect.Int8:
		fieldObj.Set(reflect.ValueOf(ToInt8(value)))
	case reflect.Int16:
		fieldObj.Set(reflect.ValueOf(ToInt16(value)))
	case reflect.Int32:
		fieldObj.Set(reflect.ValueOf(ToInt32(value)))
	case reflect.Int64:
		fieldObj.Set(reflect.ValueOf(ToInt64(value)))
	case reflect.Uint:
		fieldObj.Set(reflect.ValueOf(ToUint(value)))
	case reflect.Uint8:
		fieldObj.Set(reflect.ValueOf(ToUint8(value)))
	case reflect.Uint16:
		fieldObj.Set(reflect.ValueOf(ToUint16(value)))
	case reflect.Uint32:
		fieldObj.Set(reflect.ValueOf(ToUint32(value)))
	case reflect.Uint64:
		fieldObj.Set(reflect.ValueOf(ToUint64(value)))
	case reflect.Uintptr:
		fieldObj.Set(reflect.ValueOf(ToUintptr(value)))
	case reflect.Float32:
		fieldObj.Set(reflect.ValueOf(ToFloat32(value)))
	case reflect.Float64:
		fieldObj.Set(reflect.ValueOf(ToFloat64(value)))
	default:
		// 对于其他类型，直接设置值
		fieldObj.Set(reflect.ValueOf(value))
	}
}

func (i *Interpreter) MapIndex(key any, m any) any {
	switch m := m.(type) {
	case map[any]any:
		return m[key]
	case map[string]any:
		return m[key.(string)]
	case map[string]bool:
		return m[key.(string)]
	case map[string]string:
		return m[key.(string)]
	case map[int]bool:
		return m[ToInt(key)]
	case map[int]string:
		return m[ToInt(key)]
	case map[int]int:
		return m[ToInt(key)]
	case map[int]any:
		return m[ToInt(key)]
	case map[int64]bool:
		return m[ToInt64(key)]
	case map[int64]string:
		return m[ToInt64(key)]
	case map[int64]int64:
		return m[ToInt64(key)]
	case map[int64]any:
		return m[ToInt64(key)]
	}
	rv := assertValidObj(m)
	switch rv.Kind() {
	case reflect.Map:
		return rv.MapIndex(reflect.ValueOf(key)).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for map index", m))
	}
}

func (i *Interpreter) RIndex(index any, slice any) any {
	switch slice := slice.(type) {
	case []any:
		return slice[ToInt(index)]
	case []string:
		return slice[ToInt(index)]
	case []int:
		return slice[ToInt(index)]
	case []int32:
		return slice[ToInt(index)]
	case []int64:
		return slice[ToInt(index)]
	case []float32:
		return slice[ToInt(index)]
	case []float64:
		return slice[ToInt(index)]
	case []bool:
		return slice[ToInt(index)]
	}
	rv := assertValidObj(slice)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		return rv.Index(ToInt(index)).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for slice index", slice))
	}
}

func (i *Interpreter) RIndexStore(index any, slice any, value any) {
	switch slice := slice.(type) {
	case []any:
		slice[ToInt(index)] = value
	case []string:
		slice[ToInt(index)] = value.(string)
	case []int:
		slice[ToInt(index)] = ToInt(value)
	case []int32:
		slice[ToInt(index)] = ToInt32(value)
	case []int64:
		slice[ToInt(index)] = ToInt64(value)
	case []float32:
		slice[ToInt(index)] = ToFloat32(value)
	case []float64:
		slice[ToInt(index)] = ToFloat64(value)
	case []bool:
		slice[ToInt(index)] = (value).(bool)
	}
	rv := assertValidObj(slice)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		rv.Index(ToInt(index)).Set(reflect.ValueOf(value))
	default:
		panic(fmt.Sprintf("unexpected type %T for slice index", slice))
	}
}

func (i *Interpreter) RSet(val any, obj any) {
	rv := assertValidObj(obj)
	if rv.CanSet() {
		SetField(rv, rv.Type(), val)
		return
	}
	panic(fmt.Sprintf("unexpected type %T for Rset", obj))
}

func (i *Interpreter) RSetMapIndex(k, m, val any) {
	switch m := m.(type) {
	// fast path
	case map[any]any:
		m[k] = val
	case map[string]any:
		m[k.(string)] = val
	case map[string]bool:
		m[k.(string)] = val.(bool)
	case map[string]string:
		m[k.(string)] = val.(string)
	case map[int]bool:
		m[ToInt(k)] = val.(bool)
	case map[int]string:
		m[ToInt(k)] = val.(string)
	case map[int]int:
		m[ToInt(k)] = ToInt(val)
	case map[int]any:
		m[ToInt(k)] = val
	case map[int64]bool:
		m[ToInt64(val)] = val.(bool)
	case map[int64]string:
		m[ToInt64(val)] = val.(string)
	case map[int64]int64:
		m[ToInt64(k)] = ToInt64(val)
	case map[int64]any:
		m[ToInt64(k)] = val.(int64)
	}
	rv := assertValidObj(m)
	switch rv.Kind() {
	case reflect.Map:
		rv.Type().Key()
		rv.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(val))
	default:
		panic(fmt.Sprintf("unexpected type %T for map index store", m))
	}
}

func (i *Interpreter) Deref(ptr any) any {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("unexpected type %T for dereference", ptr))
	}
	return rv.Elem().Interface()
}

func (i *Interpreter) Addr(value any) any {
	rv := reflect.ValueOf(value)
	if !rv.CanAddr() {
		panic(fmt.Sprintf("can't address value:%v", value))
	}
	return rv.Addr().Interface()
}

func (i *Interpreter) NewPtrValue(val any) any {
	of := reflect.ValueOf(val)
	rv := reflect.New(of.Type())
	rv.Elem().Set(of)
	return rv.Interface()
}

type StructSpace struct {
	Name         string
	Fields       map[string]any
	Define       *consts.StructConst
	AllowDynamic bool
}

func (s *StructSpace) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("struct %s { ", s.Name))
	var first = true
	for _, name := range s.Define.Fields {
		if !first {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%s: %v", name, s.Fields[name])
		first = false
	}
	if s.AllowDynamic && len(s.Fields) > len(s.Define.Fields) {
		// 动态添加？
	OUTER:
		for k, v := range s.Fields {
			for _, n := range s.Define.Fields {
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

func NewStructSpace(structDef *consts.StructConst) *StructSpace {
	s := &StructSpace{Fields: make(map[string]any, len(structDef.Fields))}
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
