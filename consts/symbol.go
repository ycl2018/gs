package consts

type FunctionSymbol struct {
	Name   string
	Args   uint8
	Locals uint16
	Addr   int32
}

type Instruction struct {
	Name       string
	OpRandType int // 最长3个操作数
}

type Const struct {
	Value any
	Kind  ConstKind
}

type ConstStructDef struct {
	Name        string
	MemberNames []string
}
