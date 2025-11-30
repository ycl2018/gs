package conf

import (
	"io"
	"reflect"
)

type CompileConf struct {
	Env         any
	Optimize    bool
	DumpCode    bool
	DefineFuncs map[string]reflect.Value
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
		DefineFuncs: map[string]reflect.Value{},
	}
}
