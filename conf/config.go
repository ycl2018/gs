package conf

import (
	"io"

	"github.com/ycl2018/gs/consts"
)

type CompileConf struct {
	Env         any
	Optimize    bool
	DumpCode    bool
	DefineFuncs map[string]*consts.DefineFunc
}

type RunConf struct {
	Trace     bool
	Env       any
	Out       io.Writer
	StackSize int
}

func Default() CompileConf {
	return CompileConf{
		Optimize:    true,
		DefineFuncs: map[string]*consts.DefineFunc{},
	}
}
