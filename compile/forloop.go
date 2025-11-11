package compile

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/consts"
)

type ForLoop struct {
	Token     antlr.Token
	Breaks    []*consts.StackInstr
	Continues []*consts.StackInstr
}
