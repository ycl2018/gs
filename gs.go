package gs

import (
	"fmt"
	"io"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/compile"
	"github.com/ycl2018/gs/conf"
	"github.com/ycl2018/gs/vm"
)

type CompileOption func(conf *conf.CompileConf)

func Env(env any) CompileOption {
	return func(conf *conf.CompileConf) {
		conf.Env = env
	}
}

func DumpCode() CompileOption {
	return func(conf *conf.CompileConf) {
		conf.DumpCode = true
	}
}

func DefineFunc(name string, fn any) CompileOption {
	return func(conf *conf.CompileConf) {
		fn := reflect.ValueOf(fn)
		if fn.Kind() != reflect.Func {
			panic(fmt.Sprintf("fn %s argument must be a function, not %T", name, fn))
		}
		conf.DefineFunc[name] = fn
	}
}

type RunOption func(conf *conf.RunConf)

func Trace() RunOption {
	return func(conf *conf.RunConf) {
		conf.Trace = true
	}
}

func Output(w io.Writer) RunOption {
	return func(conf *conf.RunConf) {
		conf.Out = w
	}
}

func StackSize(size int) RunOption {
	return func(conf *conf.RunConf) {
		conf.StackSize = size
	}
}

func Compile(code string, ops ...CompileOption) (*vm.Code, error) {
	p := compile.NewGsCompiler()
	var config conf.CompileConf
	for _, op := range ops {
		op(&config)
	}
	ret, err := p.Compile(antlr.NewInputStream(code), &config)
	if err == nil && config.DumpCode {
		fmt.Println(p.Dump())
	}
	return ret, err
}

func Run(code *vm.Code, env any, ops ...RunOption) error {
	config := conf.RunConf{}
	for _, op := range ops {
		op(&config)
	}
	interpreter := vm.NewInterpreter(code, env, &config)
	return interpreter.Run()
}
