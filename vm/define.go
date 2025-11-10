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
	InstrArray      // create array
	InstrIndexLoad  // array[index]/map[key]，need reflect by vm
	InstrSliceSplit // array[start:end]
	InstrDict       // create dict
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
	InstrSliceConst
	InstrMapConst
	InstrNil // push nil
	InstrLoad

	InstrGLoad      // global load
	InstrFLoad      // filed load by name,need reflect by vm[*StructSpace or Any type]
	InstrStore      // local store
	InstrGStore     // global store
	InstrFStore     // field store
	InstrIndexStore // slice/map store
	InstrPrint
	InstrStruct     // push struct on stack
	InstrPop        // pop stack
	InstrBuildTuple // build tuple

	InstrUnpack   // unpack tuple
	InstrIterNext // push iterVal
	InstrIter     // push iter state
	InstrIterDone // check iter done
	InstrHalt

	InstrLoadEnv
	InstrRV           // reflect value
	InstrRElem        // reflect value's element
	InstrRFByIndex    // reflect load field by index
	InstrRMapIndex    // reflect map index
	InstrRIndex       // reflect slice index
	InstrRSet         // reflect set value
	InstrInterface    // reflect to interface
	InstrRSetMapIndex // reflect set map index
)

// 基于栈的指令集
var Instructions = []*Instruction{
	nil,
	// 1-10
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
	// 11
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
	// 21-30
	{"xor", nil},
	{"array", []int{INT}},
	{"index_load", nil},
	{"slice_split", nil},
	{"dict", []int{INT}},
	{"call", []int{POLL}},
	{"ret", nil},
	{"br", []int{INT}},
	{"brt", []int{INT}},
	{"brf", []int{INT}},
	// 31-40
	{"br_nil", []int{INT}},
	{"br_not_nil", nil},
	{"cconst", []int{INT}},
	{"iconst", []int{INT}},
	{"fconst", []int{POLL}},
	{"sconst", []int{POLL}},
	{"slice_const", []int{POLL}},
	{"map_const", []int{POLL}},
	{"nil", []int{}},
	{"load", []int{INT}},
	// 41-50
	{"gload", []int{INT}},
	{"fload", []int{POLL}},
	{"store", []int{INT}},
	{"gstore", []int{INT}},
	{"fstore", []int{POLL}},
	{"index_store", nil},  // slice: index value; map: key value
	{"print", []int{INT}}, // print n values
	{"struct", []int{POLL}},
	{"pop", []int{INT}}, // pop n values
	{"build_tuple", []int{INT}},
	// 51-60
	{"unpack", []int{INT}},
	{"iter_next", []int{INT}},
	{"iter", []int{INT}},
	{"iter_done", nil},
	{"halt", nil},

	{"load_env",nil},
	{"rv",nil},
	{"relem",nil},
	{"rf_by_index",nil},
	{"rmap_index",nil},
	{"rindex",nil},
	{"rset",nil},
	{"interface",nil},
	{"rset_map_index",nil},
}
