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
	InstrLShift
	InstrRShift
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
	InstrGoOuter
	InstrGoDefine
	InstrInitRef
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
	22: {"lshift", NIL},
	23: {"rshift", NIL},
	24: {"array", INT},
	25: {"index_load", NIL},
	26: {"slice_split", NIL},
	27: {"dict", INT},
	28: {"call", POLL},
	29: {"ret", NIL},
	30: {"br", INT},
	31: {"brt", INT},
	32: {"brf", INT},
	33: {"br_nil", INT},
	34: {"cconst", INT},
	35: {"iconst", INT},
	36: {"const", POLL},
	37: {"slice_const", POLL},
	38: {"map_const", POLL},
	39: {"nil", NIL},
	40: {"load", INT},
	41: {"gload", INT},
	42: {"fload", POLL},
	43: {"store", INT},
	44: {"gstore", INT},
	45: {"fstore", POLL},
	46: {"index_store", NIL},
	47: {"struct", POLL},
	48: {"pop", INT},
	49: {"build_tuple", INT},
	50: {"unpack", INT},
	51: {"iter_next", INT},
	52: {"iter", NIL},
	53: {"iter_done", NIL},
	54: {"halt", NIL},
	55: {"load_env", NIL},
	56: {"rf_by_index", POLL},
	57: {"r_set_field", POLL},
	58: {"rmap_index", NIL},
	59: {"r_index", NIL},
	60: {"rindex_store", NIL},
	61: {"rset", NIL},
	62: {"rset_map_index", NIL},
	63: {"deref", NIL},
	64: {"new_ptr_value", NIL},
	65: {"print", INT},
	66: {"printf", INT},
	67: {"println", INT},
	68: {"sprintf", INT},
	69: {"len", NIL},
	70: {"append", INT},
	71: {"append_expand", NIL},
	72: {"delete", NIL},
	73: {"toString", NIL},
	74: {"convert", INT},
	75: {"mload_byname", POLL},
	76: {"mload_byindex", INT},
	77: {"call_outer", INT},
	78: {"call_define", INT},
	79: {"go_outer", INT},
	80: {"go_define", POLL},
	81: {"init_ref", NIL},
}
