package consts

const (
	NIL = iota
	INT
	POLL
)

type Instr int

func (i Instr) String() string {
	return Instructions[i].Name
}

const (
	InstrAdd Instr = iota + 1
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
	InstrCConst // push constant
	InstrIConst
	InstrFConst
	InstrSConst
	InstrSliceConst
	InstrMapConst
	InstrNil // push nil
	InstrLoad
	InstrGLoad // global load

	InstrFLoad      // filed load by name,need reflect by vm[*StructSpace or Any type]
	InstrStore      // local store
	InstrGStore     // global store
	InstrFStore     // field store
	InstrIndexStore // slice/map store
	InstrStruct     // push struct on stack
	InstrPop        // pop stack
	InstrBuildTuple // build tuple
	InstrUnpack     // unpack tuple

	InstrIterNext // push iterVal
	InstrIter     // push iter state
	InstrIterDone // check iter done
	InstrHalt

	InstrLoadEnv
	InstrRFByIndex // reflect load field by index
	InstrRSetField // reflect set field value
	InstrRMapIndex // reflect map index
	InstrRIndex
	InstrRIndexStore  // slice index store
	InstrRSet         // reflect set value
	InstrRSetMapIndex // reflect set map index

	InstrDeref
	InstrNewPtrValue // create pointer reflect value
	// builtin call
	InstrPrint
	InstrPrintf
	InstrPrintln
	InstrSprintf
	InstrLen
	InstrAppend
	InstrAppendExpand
	InstrDelete
	InstrCopy
	InstrToString
	InstrConvert
)

// 基于栈的指令集
var Instructions = []*Instruction{
	nil,
	// 1-10
	{"add", NIL},
	{"sub", NIL},
	{"mul", NIL},
	{"div", NIL},
	{"mod", NIL},
	{"lt", NIL},
	{"gt", NIL},
	{"geq", NIL},
	{"leq", NIL},
	{"neq", NIL},
	// 11
	{"eq", NIL},
	{"or", NIL},
	{"and", NIL},
	{"pow", NIL},
	{"neg", NIL},
	{"true", NIL},
	{"false", NIL},
	{"not", NIL},
	{"bit_and", NIL},
	{"bit_or", NIL},
	// 21-30
	{"xor", NIL},
	{"array", INT},
	{"index_load", NIL},
	{"slice_split", NIL},
	{"dict", INT},
	{"call", POLL},
	{"ret", NIL},
	{"br", INT},
	{"brt", INT},
	{"brf", INT},
	// 31-40
	{"br_nil", INT},
	{"cconst", INT},
	{"iconst", INT},
	{"fconst", POLL},
	{"sconst", POLL},
	{"slice_const", POLL},
	{"map_const", POLL},
	{"nil", NIL},
	{"load", INT},
	{"gload", INT},
	// 41-50
	{"fload", POLL},
	{"store", INT},
	{"gstore", INT},
	{"fstore", POLL},
	{"index_store", NIL}, // slice: index value; map: key value

	{"struct", POLL},
	{"pop", INT}, // pop n values
	{"build_tuple", INT},
	{"unpack", INT},
	// 51-60
	{"iter_next", INT},
	{"iter", NIL},
	{"iter_done", NIL},
	{"halt", NIL},

	{"load_env", NIL},
	{"rf_by_index", POLL},
	{"r_set_field", POLL},
	{"rmap_index", NIL},
	{"r_index", NIL},
	{"rindex_store", NIL},
	{"rset", NIL},
	{"rset_map_index", NIL},

	{"deref", NIL},
	{"new_ptr_value", NIL},

	{"print", INT}, // print n values
	{"printf", INT},
	{"println", INT},
	{"sprintf", INT},
	{"len", NIL},
	{"append", INT},
	{"append_expand", NIL},
	{"delete", NIL},
	{"copy", NIL},
	{"toString", NIL},
	{"convert", INT},
}
