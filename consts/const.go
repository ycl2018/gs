package consts

type SliceInitConst struct {
	Value []any
	Name  string
}

type MapInitConst struct {
	Map  map[*ConstNode]*ConstNode
	Name string
}
