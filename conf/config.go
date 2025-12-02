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
	Out              io.Writer
	StackSize        int
	Trace            bool
	FieldIndexCache  bool
	MethodIndexCache bool
}

func Default() CompileConf {
	return CompileConf{
		Optimize:    true,
		DefineFuncs: map[string]*consts.DefineFunc{},
	}
}

func DefaultRunConf() RunConf {
	return RunConf{
		Out:              io.Discard,
		FieldIndexCache:  true,
		MethodIndexCache: true,
	}
}
