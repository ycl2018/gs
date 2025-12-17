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
	InstrConvert
	InstrMLoadByName // method load by name: push peek value's method value or func type field value
	InstrMLoadByIndex
	InstrCallOuter
	InstrCallDefine
	InstrGoOuter
	InstrGoDefine
	InstrInitRef
	InstrNewFromType
)

var Instructions = []*Instruction{
	nil,
	InstrAdd:          {"add", NIL},
	InstrSub:          {"sub", NIL},
	InstrMul:          {"mul", NIL},
	InstrDiv:          {"div", NIL},
	InstrMod:          {"mod", NIL},
	InstrLT:           {"lt", NIL},
	InstrGT:           {"gt", NIL},
	InstrGEQ:          {"geq", NIL},
	InstrLEQ:          {"leq", NIL},
	InstrNEQ:          {"neq", NIL},
	InstrEQ:           {"eq", NIL},
	InstrOR:           {"or", NIL},
	InstrAND:          {"and", NIL},
	InstrNeg:          {"neg", NIL},
	InstrTrue:         {"true", NIL},
	InstrFalse:        {"false", NIL},
	InstrNot:          {"not", NIL},
	InstrBitAND:       {"bitAnd", NIL},
	InstrBitOR:        {"bitOr", NIL},
	InstrXOR:          {"xor", NIL},
	InstrLShift:       {"lshift", NIL},
	InstrRShift:       {"rshift", NIL},
	InstrArray:        {"array", INT},
	InstrIndexLoad:    {"indexLoad", NIL},
	InstrSliceSplit:   {"sliceSplit", NIL},
	InstrDict:         {"dict", INT},
	InstrCall:         {"call", POLL},
	InstrReturn:       {"ret", NIL},
	InstrBR:           {"br", INT},
	InstrBRT:          {"brt", INT},
	InstrBRF:          {"brf", INT},
	InstrBRNil:        {"br_nil", INT},
	InstrCConst:       {"cconst", INT},
	InstrIConst:       {"iconst", INT},
	InstrConst:        {"const", POLL},
	InstrSliceConst:   {"sliceConst", POLL},
	InstrMapConst:     {"mapConst", POLL},
	InstrNil:          {"nil", NIL},
	InstrLoad:         {"load", INT},
	InstrGLoad:        {"gload", INT},
	InstrFLoad:        {"fload", POLL},
	InstrStore:        {"store", INT},
	InstrGStore:       {"gstore", INT},
	InstrFStore:       {"fstore", POLL},
	InstrIndexStore:   {"indexStore", NIL},
	InstrStruct:       {"struct", POLL},
	InstrPop:          {"pop", INT},
	InstrBuildTuple:   {"build_tuple", INT},
	InstrUnpack:       {"unpack", INT},
	InstrIterNext:     {"iter_next", INT},
	InstrIter:         {"iter", NIL},
	InstrIterDone:     {"iterDone", NIL},
	InstrHalt:         {"halt", NIL},
	InstrLoadEnv:      {"loadEnv", NIL},
	InstrRFByIndex:    {"rfByIndex", POLL},
	InstrRSetField:    {"rSetField", POLL},
	InstrRMapIndex:    {"rMapIndex", NIL},
	InstrRIndex:       {"rIndex", NIL},
	InstrRIndexStore:  {"rIndexStore", NIL},
	InstrRSet:         {"rSet", NIL},
	InstrRSetMapIndex: {"rSetMapIndex", NIL},
	InstrDeref:        {"deref", NIL},
	InstrNewPtrValue:  {"newPtrValue", NIL},
	InstrPrint:        {"print", INT},
	InstrPrintf:       {"printf", INT},
	InstrPrintln:      {"println", INT},
	InstrSprintf:      {"sprintf", INT},
	InstrLen:          {"len", NIL},
	InstrAppend:       {"append", INT},
	InstrAppendExpand: {"append_expand", NIL},
	InstrDelete:       {"delete", NIL},
	InstrConvert:      {"convert", INT},
	InstrMLoadByName:  {"mloadByName", POLL},
	InstrMLoadByIndex: {"mloadByIndex", INT},
	InstrCallOuter:    {"callOuter", INT},
	InstrCallDefine:   {"callDefine", INT},
	InstrGoOuter:      {"goOuter", INT},
	InstrGoDefine:     {"goDefine", POLL},
	InstrInitRef:      {"initRef", NIL},
	InstrNewFromType:  {"newFromType", NIL},
}
