package consts

type Debugger struct {
	Table []Info // key: instruct index, value: [line, column]
}

type Info struct {
	Line int
}
