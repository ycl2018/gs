package vm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

const (
	SupportVersion uint16 = 1
	MagicNumber    int32  = 0x25504446 // %PIE
)

type DisAssembler struct {
	Code         []byte
	ConstPool    []Const
	Instructions []*Instruction
	GlobalNums   uint16
	ip           int
	MainFuncIp   int32
}

func NewDisAssembler(instrs []*Instruction) *DisAssembler {
	return &DisAssembler{
		Instructions: instrs,
	}
}

func (d *DisAssembler) DisAssemble(code []byte) error {
	d.ip = 0
	d.Code = code
	// 校验魔数
	if d.GetInt(0) != MagicNumber {
		return errors.New("invalid magic number")
	}
	d.ip += 4
	// 校验版本号
	if ver := d.GetUInt16(int(d.ip)); ver > SupportVersion {
		return fmt.Errorf("unsupport code version:code is compiled by version:%d,current support:%d", ver, SupportVersion)
	}
	d.ip += 2
	// 读取常量池
	d.disAssembleConstPool()
	// 全局变量数
	d.GlobalNums = d.GetUInt16(d.ip)
	d.ip += 2
	d.MainFuncIp = int32(d.ip)
	return nil
}

func (d *DisAssembler) disAssembleConstPool() {
	constPoolSize := d.GetUInt16(d.ip)
	d.ip += 2
	d.ConstPool = make([]Const, 0, constPoolSize)
	for i := 0; i < int(constPoolSize); i++ {
		kind := ConstKind(d.Code[d.ip])
		d.ip++
		switch kind {
		case ConstFloat64:
			fVal := d.GetFloat32(d.ip)
			d.ip += 4
			d.ConstPool = append(d.ConstPool, Const{
				Kind:  ConstFloat64,
				Value: fVal,
			})
		case ConstString:
			strLen := d.GetUInt16(d.ip)
			d.ip += 2
			// 使用指针强转
			str := unsafe.String(&d.Code[d.ip], int(strLen))
			d.ip += int(strLen)
			d.ConstPool = append(d.ConstPool, Const{
				Kind:  ConstString,
				Value: str,
			})
		case ConstStruct:
			// 读取结构体常量
			var structDef ConstStructDef
			nameIndex := d.GetUInt16(d.ip)
			d.ip += 2
			memberCount := d.GetUInt16(d.ip)
			structDef.Name = d.ConstPool[nameIndex].Value.(string)
			d.ip += 2
			for i := 0; i < int(memberCount); i++ {
				memberIndex := d.GetUInt16(d.ip)
				d.ip += 2
				memberName := d.ConstPool[memberIndex].Value.(string)
				structDef.MemberNames = append(structDef.MemberNames, memberName)
			}
			d.ConstPool = append(d.ConstPool, Const{Kind: ConstStruct, Value: structDef})
		case ConstFunc:
			funcNameIndex := d.GetUInt16(d.ip)
			d.ip += 2
			args := d.Code[d.ip]
			d.ip++
			locals := d.GetUInt16(d.ip)
			d.ip += 2
			addr := d.GetInt(d.ip)
			d.ip += 4
			funcSym := FunctionSymbol{
				Name:   d.ConstPool[funcNameIndex].Value.(string),
				Args:   args,
				Locals: locals,
				Addr:   addr,
			}
			d.ConstPool = append(d.ConstPool, Const{Value: funcSym, Kind: ConstFunc})
		default:
			panic(fmt.Sprintf("unknown const kind %d", kind))
		}
	}
}

func (d *DisAssembler) GetInt(ip int) int32 {
	var num int32 = 0
	for i := 0; i < 4; i++ {
		// 先转为uint32避免符号扩展，移位后再转为int32
		b := int32(uint32(d.Code[ip+i]) << (8 * (3 - i)))
		num |= b
	}
	return num
}

func (d *DisAssembler) GetUInt16(ip int) uint16 {
	var num = uint16(0)
	for i := 0; i < 2; i++ {
		b := uint16(uint16(d.Code[ip+i]) << (8 * (1 - i)))
		num = num | b
	}
	return num
}

func (d *DisAssembler) GetFloat32(ip int) float32 {
	bits := d.GetInt(ip)
	return *(*float32)(unsafe.Pointer(&bits))
}

func (d *DisAssembler) Dump() (string, error) {
	var sb strings.Builder
	sb.WriteString("Disassembly:")
	sb.WriteString(fmt.Sprintf("\n.globals:%d", d.GlobalNums))
	// 常量池
	sb.WriteString(fmt.Sprintf("\n\nConstant Pool:%d", len(d.ConstPool)))
	s, err := d.DumpConstPool()
	if err != nil {
		return s, err
	}
	sb.WriteString(s)
	// 代码段
	sb.WriteString("\n\nCode:")
	var i int = int(d.MainFuncIp)
	for i < len(d.Code) {
		sb.WriteString("\n")
		next, str, err := d.DisAssembleInstruction(i)
		if err != nil {
			return "", err
		}
		sb.WriteString(str)
		i = next
	}
	return sb.String(), nil
}

func (d *DisAssembler) DumpConstPool() (string, error) {
	sb := strings.Builder{}
	for i, constVal := range d.ConstPool {
		sb.WriteString("\n")
		switch constVal.Kind {
		case ConstFloat64:
			val := constVal.Value.(float32)
			sb.WriteString(fmt.Sprintf("#%04d: float32 %f", i, val))
		case ConstString:
			val := constVal.Value.(string)
			sb.WriteString(fmt.Sprintf("#%04d: string \"%s\"", i, val))
		case ConstStruct:
			val := constVal.Value.(ConstStructDef)
			sb.WriteString(fmt.Sprintf("#%04d: struct %s {", i, val.Name))
			for _, memberName := range val.MemberNames {
				sb.WriteString(fmt.Sprintf("\n         %s;", memberName))
			}
			sb.WriteString("\n       }")
		case ConstFunc:
			funcSym := constVal.Value.(FunctionSymbol)
			sb.WriteString(fmt.Sprintf("#%04d: func %s(args:%d, locals:%d) @%d", i, funcSym.Name, funcSym.Args, funcSym.Locals, funcSym.Addr))
		default:
			return "", fmt.Errorf("unknown const kind %d in disassemble", constVal.Kind)
		}
	}
	return sb.String(), nil
}

func (d *DisAssembler) DisAssembleInstruction(i int) (int, string, error) {
	var sb strings.Builder
	instr := d.Instructions[int(d.Code[i])]
	if instr == nil {
		return 0, "", fmt.Errorf("unknown instruction:%d in code[%d]", d.Code[i], i)
	}
	instrName := instr.Name
	sb.WriteString(fmt.Sprintf("%04d:\t%-11s", i, instrName))
	i++
	if len(instr.OpRandType) == 0 {
		sb.WriteString("    ")
		return i, sb.String(), nil
	}
	var opRand []string
	for _, t := range instr.OpRandType {
		opRandNum := d.GetInt(i)
		i += 4
		switch t {
		case INT:
			opStr := strconv.Itoa(int(opRandNum))
			opRand = append(opRand, opStr)
		case POLL:
			if int(opRandNum) >= len(d.ConstPool) {
				panic(fmt.Sprintf("%d out of range", opRandNum))
			}
			switch d.ConstPool[opRandNum].Kind {
			case ConstFloat64:
				val := d.ConstPool[opRandNum].Value.(float32)
				opRand = append(opRand, fmt.Sprintf("#%d:%f", opRandNum, val))
			case ConstString:
				val := d.ConstPool[opRandNum].Value.(string)
				opRand = append(opRand, fmt.Sprintf("#%d:\"%s\"", opRandNum, val))
			case ConstStruct:
				val := d.ConstPool[opRandNum].Value.(ConstStructDef)
				opRand = append(opRand, fmt.Sprintf("#%d:struct %s", opRandNum, val.Name))
			case ConstFunc:
				funcSym := d.ConstPool[opRandNum].Value.(FunctionSymbol)
				opRand = append(opRand, fmt.Sprintf("#%d:%s()@%d", opRandNum, funcSym.Name, funcSym.Addr))
			default:
				return 0, "", fmt.Errorf("unknown const kind %d in POLL operand", d.ConstPool[opRandNum].Kind)
			}
		default:
			return 0, "", fmt.Errorf("unknown operand type %d", t)
		}
	}

	for i, s := range opRand {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(s)
	}
	return i, sb.String(), nil
}
