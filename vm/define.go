package vm

const (
	INT = iota
	POLL
)

type Instr int

func (i Instr) String() string {
	return Instructions[i].Name
}

const (
	InstrAdd = iota + 1
	InstrSub
	InstrMul
	InstrDiv
	InstrMod
	InstrLT
	InstrGT
	InstrGEQ
	InstrLEQ
	InstrNEQ
	InstrEQ
	InstrOR
	InstrAND
	InstrPow
	InstrNeg
	InstrTrue
	InstrFalse
	InstrNot
	InstrBitAND
	InstrBitOR
	InstrXOR
	InstrArray       // create array
	InstrIndexAccess // array[index]/map[key]
	InstrSliceSplit  // array[start:end]
	InstrDict        // create dict

	InstrCall
	InstrReturn
	InstrBR  // branch
	InstrBRT // branch if true
	InstrBRF // branch if false
	InstrBRNil
	InstrBRNotNil
	InstrCConst // push constant
	InstrIConst
	InstrFConst
	InstrSConst

	InstrNil // push nil

	InstrLoad
	InstrGLoad      // global load
	InstrFLoad      // filed load
	InstrStore      // local store
	InstrGStore     // global store
	InstrFStore     // field store
	InstrIndexStore // slice/map store
	InstrPrint
	InstrStruct     // push struct on stack
	InstrPop        // pop stack
	InstrBuildTuple // build tuple
	InstrUnpack     // unpack tuple
	InstrIterNext   // push iterVal
	InstrIter       // push iter state
	InstrIterDone   // check iter done

	InstrHalt
)

// 基于栈的指令集
var Instructions = []*Instruction{
	nil,
	{"add", nil},
	{"sub", nil},
	{"mul", nil},
	{"div", nil},
	{"mod", nil},
	{"lt", nil},
	{"gt", nil},
	{"geq", nil},
	{"leq", nil},
	{"neq", nil},
	{"eq", nil},
	{"or", nil},
	{"and", nil},
	{"pow", nil},
	{"neg", nil},
	{"true", nil},
	{"false", nil},
	{"not", nil},
	{"bit_and", nil},
	{"bit_or", nil},
	{"xor", nil},
	{"array", []int{INT}},
	{"index_access", []int{INT}},
	{"slice_split", []int{INT}},
	{"dict", []int{INT}},
	{"call", []int{POLL}},
	{"ret", nil},
	{"br", []int{INT}},
	{"brt", []int{INT}},
	{"brf", []int{INT}},
	{"br_nil", []int{INT}},
	{"br_not_nil", nil},
	{"cconst", []int{INT}},
	{"iconst", []int{INT}},
	{"fconst", []int{POLL}},
	{"sconst", []int{POLL}},
	{"nil", []int{}},
	{"load", []int{INT}},
	{"gload", []int{INT}},
	{"fload", []int{POLL}},
	{"store", []int{INT}},
	{"gstore", []int{INT}},
	{"fstore", []int{POLL}},
	{"index_store", nil},  // slice: index value; map: key value
	{"print", []int{INT}}, // print n values
	{"struct", []int{INT}},
	{"pop", []int{INT}}, // pop n values
	{"build_tuple", []int{INT}},
	{"unpack", []int{INT}},
	{"iter_next", []int{INT}},
	{"iter", []int{INT}},
	{"iter_done", nil},
	{"halt", nil},
}
