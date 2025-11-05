package compile

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/vm"
)

type ForLoop struct {
	Token     antlr.Token
	Breaks    []*vm.StackInstr
	Continues []*vm.StackInstr
}
