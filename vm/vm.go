package vm

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"runtime/debug"
	"slices"
	"strings"

	"github.com/ycl2018/go-future/future"
	"github.com/ycl2018/gs/conf"
	"github.com/ycl2018/gs/consts"
	"github.com/ycl2018/gs/gen"
)

const DefaultOperandStackSize = 64
const DefaultPrintStackFrameSize = 5

type Interpreter struct {
	conf.RunConf
	Code
	FP       int
	SP       int
	IP       int
	CurCode  []consts.StackInstr
	Calls    []*StackFrame
	Operands []any
	Globals  []any
	Env      any
	Result   any
}

func NewInterpreter(code *Code, env any, conf *conf.RunConf) *Interpreter {
	i := &Interpreter{
		RunConf: *conf,
		Code:    *code,
		IP:      -1,
		FP:      -1,
		SP:      -1,
		Globals: make([]any, code.GlobalNum),
		CurCode: code.MainFunc.Code,
		Env:     env,
	}
	if i.StackSize <= 0 {
		i.StackSize = DefaultOperandStackSize
	}
	i.Operands = make([]any, i.StackSize)
	return i
}

type StackFrame struct {
	ReturnAddr int
	ReturnCode []consts.StackInstr
	FuncConsts *consts.FunctionConst
	Locals     []any
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
	GlobalNum    int
	ConstPool    []consts.Const
	DefineFuncs  []consts.DefineFunc
	MainFunc     consts.FunctionConst
	BuildEnvType reflect.Type
	RuntimeCache *consts.RuntimeCache
}

func (i *Interpreter) Run() (err error) {
	i.IP = 0
	sf := NewStackFrame(&i.MainFunc, i.IP, i.CurCode)
	i.Calls = append(i.Calls, sf)
	i.FP++
	if i.Trace {
		fmt.Printf("\ntrace:\n")
	}
	defer func() {
		if r := recover(); r != nil {
			stackTrace := debug.Stack()
			var errWriter strings.Builder
			_, _ = fmt.Fprintf(&errWriter, "panic: %v\n", r)
			i.PrintStack(&errWriter)
			err = &consts.CrashError{VmStack: stackTrace, CodeTrace: errWriter.String()}
		}
	}()
	i.cpu()
	return err
}

func (i *Interpreter) PopOpStack() any {
	v := i.Operands[i.SP]
	i.SP--
	return v
}

func (i *Interpreter) PushOpStack(v any) {
	i.SP++
	if i.SP >= i.StackSize {
		i.Operands = append(i.Operands, v)
		return
	}
	i.Operands[i.SP] = v
}

func (i *Interpreter) cpu() {
	// 取指令，并执行
	instr := i.CurCode[i.IP]
	for i.IP < len(i.CurCode) {
		if i.Trace {
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
			i.PushOpStack(math.Pow(gen.ToFloat64(op1), gen.ToFloat64(op2)))
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
			funcStack := NewStackFrame(&fs, i.IP, i.CurCode)
			i.FP++
			i.Calls = append(i.Calls, funcStack)
			// 拷贝操作数到参数中
			// move args from operand stack to top frame on call stack
			for a := int(fs.ParamCount) - 1; a >= 0; a-- {
				funcStack.Locals[a] = i.PopOpStack()
			}
			i.IP, i.CurCode = 0, fs.Code
		case consts.InstrReturn:
			if i.FP == 0 {
				// main return
				i.Result = i.PopOpStack()
				return
			}
			curFs := i.Calls[i.FP]
			i.Calls = i.Calls[:i.FP]
			i.FP--
			i.IP, i.CurCode = curFs.ReturnAddr, curFs.ReturnCode
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
		case consts.InstrConst:
			poolIndex := instr.Operands
			fConst := i.ConstPool[poolIndex].Value
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
			_, _ = fmt.Fprint(i.Out, toPrint...)
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
			_, _ = fmt.Fprintf(i.Out, fmtStr, toPrint[1:]...)
		case consts.InstrPrintln:
			printNums := instr.Operands
			var toPrint = make([]any, printNums)
			for i2 := printNums - 1; i2 >= 0; i2-- {
				toPrint[i2] = i.PopOpStack()
			}
			_, _ = fmt.Fprintln(i.Out, toPrint...)
		case consts.InstrSprintf:
			printNums := instr.Operands
			var toPrint = make([]any, printNums)
			for i2 := printNums - 1; i2 >= 0; i2-- {
				toPrint[i2] = i.PopOpStack()
			}
			i.PushOpStack(fmt.Sprintf(toPrint[0].(string), toPrint[1:]...))
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
		case consts.InstrAppendExpand:
			expandSlice := i.PopOpStack()
			toSlice := i.PopOpStack()
			i.PushOpStack(appendSliceExpand(toSlice, expandSlice))
		case consts.InstrDelete:
			key := i.PopOpStack()
			m := i.PopOpStack()
			deleteMap(m, key)
		case consts.InstrConvert:
			i.PushOpStack(convert(i.PopOpStack(), reflect.Kind(instr.Operands)))
		case consts.InstrStruct:
			// push struct
			def := i.ConstPool[instr.Operands].Value.(consts.StructConst)
			s := NewStructSpace(&def)
			i.PushOpStack(s)
		case consts.InstrPop:
			for range instr.Operands {
				i.PopOpStack()
			}
		case consts.InstrBuildTuple:
			i.PushOpStack(i.BuildTuple(instr.Operands))
		case consts.InstrUnpack:
			t := i.PopOpStack().(consts.Tuple)
			num := instr.Operands
			if num != t.Num {
				panic(fmt.Sprintf("unpack %d items to %d variables", t.Num, num))
			}
			for i2 := num - 1; i2 >= 0; i2-- {
				i.PushOpStack(t.Values[i2])
			}
		case consts.InstrIter:
			i.PushOpStack(Iter(i.PopOpStack()))
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
			i.PushOpStack(FieldByIndex(i.PopOpStack(), i.ConstPool[instr.Operands].Value.([]*reflect.StructField)))
		case consts.InstrRSetField:
			i.RSetField(i.ConstPool[instr.Operands].Value.([]*reflect.StructField))
		case consts.InstrRMapIndex:
			i.PushOpStack(MapIndex(i.PopOpStack(), i.PopOpStack()))
		case consts.InstrRIndex:
			i.PushOpStack(RIndex(i.PopOpStack(), i.PopOpStack()))
		case consts.InstrRIndexStore:
			RIndexStore(i.PopOpStack(), i.PopOpStack(), i.PopOpStack())
		case consts.InstrRSet:
			RSet(i.PopOpStack(), i.PopOpStack())
		case consts.InstrRSetMapIndex:
			RSetMapIndex(i.PopOpStack(), i.PopOpStack(), i.PopOpStack())
		case consts.InstrDeref:
			i.PushOpStack(Deref(i.PopOpStack()))
		case consts.InstrNewPtrValue:
			i.PushOpStack(NewPtrValue(i.PopOpStack()))
		case consts.InstrMLoadByName:
			methodName := i.ConstPool[instr.Operands].Value.(string)
			obj := i.PopOpStack()
			i.PushOpStack(i.loadMethod(obj, methodName))
		case consts.InstrMLoadByIndex:
			obj := i.PopOpStack()
			i.PushOpStack(loadMethodByIndex(obj, instr.Operands))
		case consts.InstrCallOuter:
			fn := i.PopOpStack()
			fnValue := reflect.ValueOf(fn)
			i.callFn(instr.Operands, fnValue)
		case consts.InstrCallDefine:
			fn := i.DefineFuncs[instr.Operands]
			if fn.Fast {
				var inArgs = make([]any, fn.NumIn)
				for j := fn.NumIn - 1; j >= 0; j-- {
					inArgs[j] = i.PopOpStack()
				}
				result := fn.Fn.(func([]any) []any)(inArgs)
				if len(result) != fn.NumOut {
					panic(fmt.Sprintf("func:%s define %d numOut but get:%d", fn.Name, fn.NumOut, len(result)))
				}
				if len(result) == 0 {
					i.PushOpStack(nil)
				} else if len(result) == 1 {
					i.PushOpStack(result[0])
				} else {
					i.PushOpStack(consts.Tuple{
						Values: result,
						Num:    len(result),
					})
				}
			} else {
				i.callFn(fn.NumIn, fn.Fn.(reflect.Value))
			}
		case consts.InstrGoOuter:
			obj := i.PopOpStack()
			fn := reflect.ValueOf(obj)
			i.goFn(fn, instr.Operands)
		case consts.InstrGoDefine:
			fn := i.DefineFuncs[instr.Operands]
			if fn.NumOut != 2 {
				panic(fmt.Sprintf("function %s signature is invalid:go can't run func that returns 2 values", fn.Name))
			}
			if fn.Fast {
				var inArgs = make([]any, fn.NumIn)
				for j := fn.NumIn - 1; j >= 0; j-- {
					inArgs[j] = i.PopOpStack()
				}
				f := future.Go(func() (any, error) {
					result := fn.Fn.(func([]any) []any)(inArgs)
					if len(result) != fn.NumOut {
						panic(fmt.Sprintf("func:%s define %d numOut but get:%d", fn.Name, fn.NumOut, len(result)))
					}
					err := result[1]
					if err == nil {
						return result[0], nil
					}
					return result[0], result[1].(error)
				})
				i.PushOpStack(f)
			} else {
				i.goFn(fn.Fn.(reflect.Value), fn.NumIn)
			}
		case consts.InstrInitRef:
			i.PushOpStack(i.initRef(i.PopOpStack(), instr.Operands))
		case consts.InstrNewFromType:
			i.PushOpStack(newFromType(i.PopOpStack()))
		case consts.InstrHalt:
			return
		default:
			panic(fmt.Sprintf("unknown opcode:%s", instr))
		}
		instr = i.CurCode[i.IP]
	}
}

// newFromType create a new object from type
// if obj is a pointer, return a new pointer to the element type
// if obj is not a pointer, return a new pointer to the type
func newFromType(obj any) any {
	if obj == nil {
		panic("newFromType: nil object")
	}
	rt := reflect.TypeOf(obj)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	rv := reflect.New(rt)
	return rv.Interface()
}

func (i *Interpreter) goFn(fn reflect.Value, numIn int) {
	if !fn.IsValid() {
		panic(fmt.Sprintf("call func is nil"))
	}
	if fn.Kind() != reflect.Func {
		panic(fmt.Sprintf("value is not func,got:%s", fn.Type().String()))
	}
	fnType := fn.Type()
	if fn.Type().NumOut() != 2 || fnType.Out(1) != reflect.TypeOf((*error)(nil)).Elem() {
		panic(fmt.Sprintf("go can't run func that returns (T,error) but got :%s", fnType.String()))
	}
	inNum := numIn
	var inArgs = make([]reflect.Value, inNum)
	for j := inNum - 1; j >= 0; j-- {
		arg := reflect.ValueOf(i.PopOpStack())
		inArgs[j] = arg
	}
	f := future.Go(func() (any, error) {
		call := fn.Call(inArgs)
		err := call[1].Interface()
		if err == nil {
			return call[0].Interface(), nil
		}
		return call[0].Interface(), call[1].Interface().(error)
	})
	i.PushOpStack(f)
}

func (i *Interpreter) callFn(inNum int, fn reflect.Value) {
	if !fn.IsValid() {
		panic(fmt.Sprintf("call func is nil"))
	}
	if fn.Kind() != reflect.Func {
		panic(fmt.Sprintf("value is not func,got:%s", fn.Type().String()))
	}
	var inArgs = make([]reflect.Value, inNum)
	for j := inNum - 1; j >= 0; j-- {
		arg := reflect.ValueOf(i.PopOpStack())
		inArgs[j] = arg
	}
	result := fn.Call(inArgs)
	if len(result) == 0 {
		i.PushOpStack(nil)
	} else if len(result) == 1 {
		i.PushOpStack(result[0].Interface())
	} else {
		var ret []any
		for _, value := range result {
			ret = append(ret, value.Interface())
		}
		i.PushOpStack(consts.Tuple{
			Values: ret,
			Num:    len(ret),
		})
	}
}

func loadMethodByIndex(obj any, index int) any {
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		panic("load method from nil object")
	}
	if rv.NumMethod() == 0 {
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		} else {
			ptrTo := reflect.New(rv.Type())
			ptrTo.Elem().Set(rv)
			rv = ptrTo
		}
	}
	return rv.Method(index).Interface()
}

func (i *Interpreter) loadMethod(obj any, methodName string) any {
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		panic("load method from nil object")
	}
	if i.MethodIndexCache {
		vt := rv.Type()
		if index, ok := i.RuntimeCache.FetchMethodIndex(vt, methodName); ok {
			if index.IsMethod {
				switch index.Convert {
				case consts.Elem:
					rv = rv.Elem()
				case consts.PtrTo:
					ptrTo := reflect.New(rv.Type())
					ptrTo.Elem().Set(rv)
					rv = ptrTo
				default:
				}
				return rv.Method(index.Index[0]).Interface()
			} else {
				if rv.Kind() == reflect.Ptr {
					rv = rv.Elem()
				}
				return rv.FieldByIndex(index.Index).Interface()
			}
		} else {
			origin := rv
			conv := consts.No
			// try to find method by name
			if rv.NumMethod() == 0 {
				if rv.Kind() == reflect.Ptr {
					rv = rv.Elem()
					conv = consts.Elem
				} else {
					ptrTo := reflect.New(rv.Type())
					ptrTo.Elem().Set(rv)
					rv = ptrTo
					conv = consts.PtrTo
				}
			}
			m, ok := rv.Type().MethodByName(methodName)
			if ok {
				method := rv.Method(m.Index)
				if method.IsValid() {
					if method.CanInterface() {
						i.RuntimeCache.SetMethodIndex(origin.Type(), methodName, consts.MethodIndex{
							Index:    []int{m.Index},
							IsMethod: true,
							Convert:  conv,
						})
						return method.Interface()
					}
					panic(fmt.Sprintf("method '%s' is not exported by type:%T", methodName, obj))
				}
			}
			// try to find field by name
			if origin.Kind() == reflect.Ptr {
				rv = origin.Elem()
			}
			fieldByName, ok := rv.Type().FieldByName(methodName)
			if ok {
				field := rv.FieldByIndex(fieldByName.Index)
				if field.IsValid() && field.Kind() == reflect.Func {
					i.RuntimeCache.SetMethodIndex(origin.Type(), methodName, consts.MethodIndex{
						Index:    fieldByName.Index,
						IsMethod: false,
					})
					return field.Interface()
				}
			}
			panic(fmt.Sprintf("no such method/field '%s' by type:%T", methodName, obj))
		}
	}
	if rv.NumMethod() == 0 {
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		} else {
			ptrTo := reflect.New(rv.Type())
			ptrTo.Elem().Set(rv)
			rv = ptrTo
		}
	}
	method := rv.MethodByName(methodName)
	if method.IsValid() {
		return method.Interface()
	}
	rv = assertValidObj(obj)
	field := rv.FieldByName(methodName)
	if field.IsValid() && field.Kind() == reflect.Func && field.CanInterface() {
		return field.Interface()
	}
	panic(fmt.Sprintf("no such method/field '%s' by type:%T", methodName, obj))
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
	for j := 0; j <= i.SP; j++ {
		fmt.Printf(" %v", i.Operands[j])
	}
	fmt.Print(" ]")
	// call stack
	if i.FP >= 0 {
		fmt.Printf(", calls=[")
		for j := 0; j <= i.FP; j++ {
			fmt.Printf(" " + i.Calls[j].FuncConsts.Name)
		}
		fmt.Print(" ]\n")
	}
	fmt.Printf("->%04d: %s\n", i.IP, i.CurCode[i.IP])
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
	for j := 0; j < len(i.CurCode); j++ {
		if j%8 == 0 && j != 0 {
			fmt.Println()
		}
		if j%8 == 0 {
			fmt.Printf("%04d:", j)
		}
		fmt.Printf(" %3s", i.CurCode[j].OpCode.String())
	}
	fmt.Println()
}

func (i *Interpreter) dumpDataMemory() {
	fmt.Println("Data memory:")
	for j, a := range i.Globals {
		if a != nil {
			fmt.Printf("%04d: %v %s\n", j, a, reflect.TypeOf(a).Name())
		} else {
			fmt.Printf("%04d: nil \n", j)
		}
	}
	fmt.Println()
}

func (i *Interpreter) Index(obj any, index any) any {
	switch obj := obj.(type) {
	case []any:
		return obj[gen.ToInt(index)]
	case map[any]any:
		return obj[index]
	case map[string]any:
		return obj[index.(string)]
	default:
		rv := assertValidObj(obj)
		switch rv.Kind() {
		case reflect.Slice, reflect.Array:
			return rv.Index(gen.ToInt(index)).Interface()
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
	if rv.Kind() == reflect.Array && !rv.CanAddr() {
		newValue := reflect.New(rv.Type()).Elem()
		newValue.Set(rv)
		rv = newValue
	}
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		l, r := gen.ToInt(start), gen.ToInt(end)
		if l == -1 {
			l = 0
		}
		if r == -1 {
			r = rv.Len()
		}
		return rv.Slice(l, r).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for slice split", obj))
	}
}

func (i *Interpreter) MakeMap(dictLen int) any {
	mapLen := dictLen / 2
	m := make(map[any]any, mapLen)
	for i2 := 0; i2 < mapLen; i2++ {
		v := i.PopOpStack()
		k := i.PopOpStack()
		m[k] = v
	}
	return m
}

func (i *Interpreter) Peek() any {
	return i.Operands[i.SP]
}

func (i *Interpreter) FieldLoad(field string) any {
	// build-in type
	obj := i.PopOpStack()
	if structSpace, ok := obj.(*StructSpace); ok {
		return structSpace.Fields[field]
	}
	validObj := assertValidObj(obj)
	switch validObj.Kind() {
	case reflect.Struct:
		// reflect
		if i.FieldIndexCache {
			vt := validObj.Type()
			if indexes, ok := i.RuntimeCache.FetchFieldIndex(vt, field); ok {
				return validObj.FieldByIndex(indexes).Interface()
			}
			if index, ok := i.RuntimeCache.FetchMethodIndex(vt, field); ok {
				return validObj.Method(index.Index[0]).Interface()
			}
			if fieldStruct, ok := vt.FieldByName(field); ok {
				if !fieldStruct.IsExported() {
					panic(fmt.Sprintf("field '%s' is not exported by type:%T", field, obj))
				}
				i.RuntimeCache.SetFieldIndex(vt, field, fieldStruct.Index)
				return validObj.FieldByIndex(fieldStruct.Index).Interface()
			}
			method := i.loadMethod(obj, field)
			return method
		} else {
			fieldByName := validObj.FieldByName(field)
			if fieldByName.IsValid() {
				if fieldByName.CanInterface() {
					return fieldByName.Interface()
				} else {
					panic(fmt.Sprintf("field '%s' is not exported by type:%T", field, obj))
				}
			}
			method := i.loadMethod(obj, field)
			return method
		}
	default:
		panic(fmt.Sprintf("unexpected type %T for load %s", obj, field))
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
	if i.FieldIndexCache {
		vt := objStruct.Type()
		if indexes, ok := i.RuntimeCache.FetchFieldIndex(vt, field); ok {
			fieldObj := objStruct.FieldByIndex(indexes)
			SetField(fieldObj, fieldObj.Type(), val)
			return
		}
		if fieldStruct, ok := vt.FieldByName(field); ok {
			if !fieldStruct.IsExported() {
				panic(fmt.Sprintf("field '%s' is not exported by type:%T", field, obj))
			}
			i.RuntimeCache.SetFieldIndex(vt, field, fieldStruct.Index)
			fieldObj := objStruct.FieldByIndex(fieldStruct.Index)
			SetField(fieldObj, fieldObj.Type(), val)
			return
		}
		panic(fmt.Sprintf("field '%s' is not find/exported by type:%T", field, obj))
	}
	fieldObj := objStruct.FieldByName(field)
	if !fieldObj.IsValid() {
		panic(fmt.Sprintf("field '%s' is not find/exported by type:%T", field, obj))
	}
	switch objStruct.Kind() {
	case reflect.Struct:
		SetField(fieldObj, fieldObj.Type(), val)
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
		obj[gen.ToInt(index)] = val
		return
	case []string:
		obj[gen.ToInt(index)] = val.(string)
	case []int:
		obj[gen.ToInt(index)] = gen.ToInt(val)
	case []int64:
		obj[gen.ToInt(index)] = gen.ToInt64(val)
	case map[string]any:
		obj[index.(string)] = val
		return
	case map[string]bool:
		obj[index.(string)] = val.(bool)
	case map[string]int:
		obj[index.(string)] = gen.ToInt(val)
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
		//rv.Index(ToInt(index)).Set(reflect.ValueOf(val))
		SetField(rv.Index(gen.ToInt(index)), rv.Type().Elem(), val)
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

func Iter(obj any) any {
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

func FieldByIndex(obj any, fields []*reflect.StructField) any {
	rv := assertValidObj(obj)
	var index []int
	for _, f := range fields {
		index = append(index, f.Index...)
	}
	value, err := rv.FieldByIndexErr(index)
	if err != nil {
		panic("nil pointer")
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
		panic("nil pointer")
	}
	lastFiled := fields[len(fields)-1]
	SetField(fieldObj, lastFiled.Type, value)
}

// SetField 设置结构体字段的值，支持类型转换
func SetField(fieldObj reflect.Value, fieldType reflect.Type, value any) {
	switch fieldType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldObj.SetInt(gen.ToInt64(value))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fieldObj.SetUint(gen.ToUint64(value))
	case reflect.Float32, reflect.Float64:
		fieldObj.SetFloat(gen.ToFloat64(value))
	case reflect.String:
		fieldObj.SetString(value.(string))
	case reflect.Pointer, reflect.Interface, reflect.Map, reflect.Slice:
		if value == nil {
			fieldObj.SetZero()
			return
		}
		fieldObj.Set(reflect.ValueOf(value))
	default:
		fieldObj.Set(reflect.ValueOf(value))
	}
}

func MapIndex(key any, m any) any {
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
		return m[gen.ToInt(key)]
	case map[int]string:
		return m[gen.ToInt(key)]
	case map[int]int:
		return m[gen.ToInt(key)]
	case map[int]any:
		return m[gen.ToInt(key)]
	case map[int64]bool:
		return m[gen.ToInt64(key)]
	case map[int64]string:
		return m[gen.ToInt64(key)]
	case map[int64]int64:
		return m[gen.ToInt64(key)]
	case map[int64]any:
		return m[gen.ToInt64(key)]
	}
	rv := assertValidObj(m)
	switch rv.Kind() {
	case reflect.Map:
		return rv.MapIndex(reflect.ValueOf(key)).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for map index", m))
	}
}

func RIndex(index any, slice any) any {
	switch slice := slice.(type) {
	case []any:
		return slice[gen.ToInt(index)]
	case []string:
		return slice[gen.ToInt(index)]
	case []int:
		return slice[gen.ToInt(index)]
	case []int32:
		return slice[gen.ToInt(index)]
	case []int64:
		return slice[gen.ToInt(index)]
	case []float32:
		return slice[gen.ToInt(index)]
	case []float64:
		return slice[gen.ToInt(index)]
	case []bool:
		return slice[gen.ToInt(index)]
	}
	rv := assertValidObj(slice)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		return rv.Index(gen.ToInt(index)).Interface()
	default:
		panic(fmt.Sprintf("unexpected type %T for slice index", slice))
	}
}

func RIndexStore(index any, slice any, value any) {
	switch slice := slice.(type) {
	case []any:
		slice[gen.ToInt(index)] = value
	case []string:
		slice[gen.ToInt(index)] = value.(string)
	case []int:
		slice[gen.ToInt(index)] = gen.ToInt(value)
	case []int32:
		slice[gen.ToInt(index)] = gen.ToInt32(value)
	case []int64:
		slice[gen.ToInt(index)] = gen.ToInt64(value)
	case []float32:
		slice[gen.ToInt(index)] = gen.ToFloat32(value)
	case []float64:
		slice[gen.ToInt(index)] = gen.ToFloat64(value)
	case []bool:
		slice[gen.ToInt(index)] = (value).(bool)
	}
	rv := assertValidObj(slice)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		SetField(rv.Index(gen.ToInt(index)), rv.Type().Elem(), value)
	default:
		panic(fmt.Sprintf("unexpected type %T for slice index", slice))
	}
}

func RSet(val any, obj any) {
	rv := assertValidObj(obj)
	if rv.CanSet() {
		SetField(rv, rv.Type(), val)
		return
	}
	panic(fmt.Sprintf("unexpected type %T for Rset", obj))
}

func RSetMapIndex(k, m, val any) {
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
		m[gen.ToInt(k)] = val.(bool)
	case map[int]string:
		m[gen.ToInt(k)] = val.(string)
	case map[int]int:
		m[gen.ToInt(k)] = gen.ToInt(val)
	case map[int]any:
		m[gen.ToInt(k)] = val
	case map[int64]bool:
		m[gen.ToInt64(val)] = val.(bool)
	case map[int64]string:
		m[gen.ToInt64(val)] = val.(string)
	case map[int64]int64:
		m[gen.ToInt64(k)] = gen.ToInt64(val)
	case map[int64]any:
		m[gen.ToInt64(k)] = val.(int64)
	}
	rv := assertValidObj(m)
	switch rv.Kind() {
	case reflect.Map:
		rv.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(val))
	default:
		panic(fmt.Sprintf("unexpected type %T for map index store", m))
	}
}

func Deref(ptr any) any {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("unexpected type %T for dereference", ptr))
	}
	return rv.Elem().Interface()
}

func NewPtrValue(val any) any {
	of := reflect.ValueOf(val)
	rv := reflect.New(of.Type())
	rv.Elem().Set(of)
	return rv.Interface()
}

func (i *Interpreter) PrintStack(writer io.Writer) {
	var printFrameSize int
	_, _ = fmt.Fprintf(writer, "stack trace:\n")
	for j := len(i.Calls) - 1; j >= 0; j-- {
		// 打印调用栈
		ip := i.IP - 1
		// func:<funcName>,args:(arg1,arg2,...) line:<line>
		call := i.Calls[j]
		_, _ = fmt.Fprintf(writer, "\tat "+call.FuncConsts.Name+" args:(")
		// args
		for k := range call.FuncConsts.ParamCount {
			if k < len(call.Locals) {
				if k > 0 {
					_, _ = fmt.Fprintf(writer, ", ")
				}
				_, _ = fmt.Fprintf(writer, "%v", call.Locals[k])
			}
		}
		_, _ = fmt.Fprintf(writer, ") line:%d\n", call.FuncConsts.Debugger.Table[ip].Line)
		printFrameSize++
		if printFrameSize >= DefaultPrintStackFrameSize {
			break
		}
	}
}

// initRef init ptr/map/slice with length
// when operand == 1, pop length from op stack
func (i *Interpreter) initRef(obj any, operand int) any {
	if obj == nil {
		panic("init ref: got nil interface{}")
	}
	rv := reflect.ValueOf(obj)
	switch rv.Kind() {
	case reflect.Ptr:
		if !rv.IsNil() {
			panic("init ref: ref obj is not empty/nil")
		}
		if operand == 1 {
			panic("init ref: can't init pointer with length")
		}
		return reflect.New(rv.Type().Elem()).Interface()
	case reflect.Map:
		if !rv.IsNil() {
			panic("init ref: ref obj is not empty/nil")
		}
		var length int
		if operand == 1 {
			length = gen.ToInt(i.PopOpStack())
		}
		return reflect.MakeMapWithSize(rv.Type(), length).Interface()
	case reflect.Slice:
		if !rv.IsNil() {
			panic("init ref: ref obj is not empty/nil")
		}
		var length int
		if operand == 1 {
			length = gen.ToInt(i.PopOpStack())
		}
		return reflect.MakeSlice(rv.Type(), length, length).Interface()
	default:
		panic(fmt.Sprintf("init ref: ref obj is not pointer/map/slice type: %T", obj))
	}
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
		_, _ = fmt.Fprintf(&sb, "%s: %v", name, s.Fields[name])
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
			_, _ = fmt.Fprintf(&sb, "%s: %v\n", k, v)
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
