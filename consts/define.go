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
	InstrConst // load const
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
	InstrStruct     // push struct on stack
	InstrPop        // pop stack
	InstrBuildTuple // build tuple
	InstrUnpack     // unpack tuple
	InstrIterNext   // push iterVal
	InstrIter       // push iter state
	InstrIterDone   // check iter done
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
	InstrPrint
	InstrPrintf
	InstrPrintln
	InstrSprintf
	InstrLen
	InstrAppend
	InstrAppendExpand
	InstrDelete
	InstrToString
	InstrConvert
	InstrMLoadByName // method load by name: push peek value's method value or func type field value
	InstrMLoadByIndex
	InstrCallOuter
	InstrCallDefine
)

var Instructions = []*Instruction{
	nil,
	1:  {"add", NIL},
	2:  {"sub", NIL},
	3:  {"mul", NIL},
	4:  {"div", NIL},
	5:  {"mod", NIL},
	6:  {"lt", NIL},
	7:  {"gt", NIL},
	8:  {"geq", NIL},
	9:  {"leq", NIL},
	10: {"neq", NIL},
	11: {"eq", NIL},
	12: {"or", NIL},
	13: {"and", NIL},
	14: {"pow", NIL},
	15: {"neg", NIL},
	16: {"true", NIL},
	17: {"false", NIL},
	18: {"not", NIL},
	19: {"bit_and", NIL},
	20: {"bit_or", NIL},
	21: {"xor", NIL},
	22: {"array", INT},
	23: {"index_load", NIL},
	24: {"slice_split", NIL},
	25: {"dict", INT},
	26: {"call", POLL},
	27: {"ret", NIL},
	28: {"br", INT},
	29: {"brt", INT},
	30: {"brf", INT},
	31: {"br_nil", INT},
	32: {"cconst", INT},
	33: {"iconst", INT},
	34: {"const", POLL},
	35: {"slice_const", POLL},
	36: {"map_const", POLL},
	37: {"nil", NIL},
	38: {"load", INT},
	39: {"gload", INT},
	40: {"fload", POLL},
	41: {"store", INT},
	42: {"gstore", INT},
	43: {"fstore", POLL},
	44: {"index_store", NIL},
	45: {"struct", POLL},
	46: {"pop", INT},
	47: {"build_tuple", INT},
	48: {"unpack", INT},
	49: {"iter_next", INT},
	50: {"iter", NIL},
	51: {"iter_done", NIL},
	52: {"halt", NIL},
	53: {"load_env", NIL},
	54: {"rf_by_index", POLL},
	55: {"r_set_field", POLL},
	56: {"rmap_index", NIL},
	57: {"r_index", NIL},
	58: {"rindex_store", NIL},
	59: {"rset", NIL},
	60: {"rset_map_index", NIL},
	61: {"deref", NIL},
	62: {"new_ptr_value", NIL},
	63: {"print", INT},
	64: {"printf", INT},
	65: {"println", INT},
	66: {"sprintf", INT},
	67: {"len", NIL},
	68: {"append", INT},
	69: {"append_expand", NIL},
	70: {"delete", NIL},
	71: {"toString", NIL},
	72: {"convert", INT},
	73: {"mload_byname", NIL},
	74: {"mload_byindex", NIL},
	75: {"call_outer", INT},
	76: {"call_define", INT},
}
