package conf

import (
	"io"
	"reflect"
)

type CompileConf struct {
	Env        any
	Optimize   bool
	DumpCode   bool
	DefineFunc map[string]reflect.Value
}

type RunConf struct {
	Trace     bool
	Env       any
	Out       io.Writer
	StackSize int
}
