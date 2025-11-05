package compile

import (
	"fmt"
	"sort"
	"unsafe"

	"github.com/ycl2018/gs/vm"
)

// 文件头：
// 4字节：%Gs
// 2字节：版本号
// 2字节：常量池大小
// 常量池
// 2字节：全局变量数量
// 代码段
type Assembler struct {
	GScope            *GlobalScope
	MainFunc          *FunctionSymbol
	AllFunc           []*FunctionSymbol
	ToBeFilled        map[*vm.StackInstr][]*vm.StackInstr // 目标分支跳转Tag -> 需要回填的分支跳转指令列表
	FuncToAddr        map[string]int32                    // 函数符号 -> 待回填的函数入口地址位置
	ToBeFilledInstrIP map[*vm.StackInstr]int32            // 目标分支跳转操作数在 code 中的位置
	FilledTag         map[*vm.StackInstr]int32            // 已回填的目标分支跳转 ip 位置
	Code              []byte
	IP                int32
}

func NewAssembler(gScope *GlobalScope, mainFunc *FunctionSymbol, allFunc []*FunctionSymbol, toBeFilled map[*vm.StackInstr][]*vm.StackInstr) *Assembler {
	return &Assembler{
		GScope:            gScope,
		MainFunc:          mainFunc,
		AllFunc:           allFunc,
		ToBeFilled:        toBeFilled,
		FuncToAddr:        map[string]int32{},
		ToBeFilledInstrIP: map[*vm.StackInstr]int32{},
		FilledTag:         map[*vm.StackInstr]int32{},
	}
}

func (a *Assembler) Assemble() []byte {
	// 文件头
	a.writeInt32(MagicNumber)
	a.writeUInt16(Version)
	// 常量池
	a.writeConstPool()
	// 全局变量数量
	a.writeUInt16(a.GScope.LocalVarAllocator)
	// 代码段
	// 函数体
	a.writeFuncBody(a.MainFunc)
	a.writeAllFuncBody(a.AllFunc)
	// 回填目标分支跳转操作数
	a.fillAllTargetInstrIP()
	// 返回代码
	return a.Code
}

func (a *Assembler) writeConstPool() {
	// 常量池大小
	a.writeUInt16(uint16(len(a.GScope.Consts)))
	// 常量池内容
	var constPool []*ConstSymbol
	for _, constVal := range a.GScope.Consts {
		constPool = append(constPool, constVal)
	}
	// 排序
	sort.Slice(constPool, func(i, j int) bool {
		return constPool[i].Address < constPool[j].Address
	})
	// 写入常量池
	for _, constVal := range constPool {
		a.appendByte(byte(constVal.Kind))
		switch constVal.Kind {
		case vm.ConstFloat64:
			a.writeFloat32(constVal.Value.(float32)) // [tag(1byte)][4bytes float32]
		case vm.ConstString:
			a.writeString(constVal.Value.(string)) // [tag(1byte)][len(2byte)][n bytes UTF-8]
		case vm.ConstStruct: // [tag(1byte)][name_index(2byte)][member_count(2byte)][member1_index(2byte)][member2_index(2byte)]...
			a.writeStruct(constVal)
		case vm.ConstFunc:
			a.writeFunc(constVal)
		default:
			panic(fmt.Sprintf("unknown const kind %d", constVal.Kind))
		}
	}

}

func (a *Assembler) appendByte(bt byte) {
	a.Code = append(a.Code, bt)
	a.IP++
}

func (a *Assembler) writeInt32(v int32) {
	a.appendByte(byte((v >> (8 * 3)) & 0xFF))
	a.appendByte(byte((v >> (8 * 2)) & 0xFF))
	a.appendByte(byte((v >> (8 * 1)) & 0xFF))
	a.appendByte(byte(v & 0xFF))
}

func (a *Assembler) writeFloat32(v float32) {
	raw := *(*int32)(unsafe.Pointer(&v))
	a.writeInt32(raw)
}

// [tag(1byte)][len(2byte)][n bytes UTF-8]
func (a *Assembler) writeString(v string) {
	a.writeUInt16(uint16(len(v)))
	for i := 0; i < len(v); i++ {
		a.appendByte(v[i])
	}
}

func (a *Assembler) writeUInt16(v uint16) {
	a.appendByte(byte((v >> (8 * 1)) & 0xFF))
	a.appendByte(byte(v & 0xFF))
}

// [tag(1byte)][name_index(2byte)][member_count(2byte)][member1_index(2byte)][member2_index(2byte)]...
func (a *Assembler) writeStruct(val *ConstSymbol) {
	a.writeUInt16(uint16(val.Fields[0]))       // name_index
	a.writeUInt16(uint16(len(val.Fields) - 1)) // member_count
	for i := 0; i < len(val.Fields)-1; i++ {
		a.writeUInt16(uint16(val.Fields[i+1])) // member_index
	}
}

// [tag(1byte)][func_name_index(2byte)][args(1byte)][locals(2byte)][addr(4byte)]
func (a *Assembler) writeFunc(val *ConstSymbol) {
	a.writeUInt16(uint16(val.Fields[0]))
	a.appendByte(byte(val.Fields[1]))
	a.writeUInt16(uint16(val.Fields[2]))
	// 记录函数入口地址，后续回填
	a.FuncToAddr[val.Name] = a.IP
	a.writeInt32(int32(val.Fields[3]))

}

func (a *Assembler) writeFuncBody(f *FunctionSymbol) {
	// 写入主函数体
	for _, instr := range f.Code {
		a.writeInstr(instr)
	}
}

func (a *Assembler) writeAllFuncBody(allFunc []*FunctionSymbol) {
	for _, f := range allFunc {
		fName := f.Name
		// 回填函数入口地址
		index := a.FuncToAddr[fName]
		if index <= 0 {
			panic(fmt.Sprintf("invalid func addr %d", index))
		}
		a.fillInt32(index, a.IP)
		a.writeFuncBody(f)
	}
}

func (a *Assembler) writeInstr(instr *vm.StackInstr) {
	// 识别跳转指令，记录跳转目标地址，后续回填
	switch instr.OpCode {
	case vm.InstrBR, vm.InstrBRF, vm.InstrBRT:
		a.ToBeFilledInstrIP[instr] = a.IP + 1
	default:
		if instr.OpCode < 0 {
			// 常量池索引
			a.FilledTag[instr] = a.IP
			return // 跳过 tag 虚拟指令
		}
	}
	a.appendByte(byte(instr.OpCode))
	for _, operand := range instr.Operands {
		a.writeInt32(int32(operand))
	}
}

// fillInt32 在 code 中 index 位置填充 operand
func (a *Assembler) fillInt32(index int32, operand int32) {
	a.Code[index] = byte((operand >> (8 * 3)) & 0xFF)
	a.Code[index+1] = byte((operand >> (8 * 2)) & 0xFF)
	a.Code[index+2] = byte((operand >> (8 * 1)) & 0xFF)
	a.Code[index+3] = byte(operand & 0xFF)
}

func (a *Assembler) fillAllTargetInstrIP() {
	for target, instrs := range a.ToBeFilled {
		tagIndex := a.FilledTag[target]
		for _, instr := range instrs {
			index := a.ToBeFilledInstrIP[instr]
			if index <= 0 {
				panic(fmt.Sprintf("invalid index %d", index))
			}
			a.fillInt32(index, tagIndex)
		}
	}
}
