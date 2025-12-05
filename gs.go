package gs

import (
	"fmt"
	"io"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/compile"
	"github.com/ycl2018/gs/conf"
	"github.com/ycl2018/gs/consts"
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

type Func struct {
	Name string
	Fn   any
}

type FastFunc struct {
	Name   string
	NumIn  int
	NumOut int
	Fn     func([]any) []any
}

func DefineFast(fns ...FastFunc) CompileOption {
	return func(conf *conf.CompileConf) {
		for _, fn := range fns {
			conf.DefineFuncs.Define(fn.Name, &consts.DefineFunc{
				Name:   fn.Name,
				NumIn:  fn.NumIn,
				NumOut: fn.NumOut,
				Fast:   true,
				Fn:     fn.Fn,
			})
		}
	}
}

func DefineFuncs(fns ...Func) CompileOption {
	return func(conf *conf.CompileConf) {
		for _, fn := range fns {
			rv := reflect.ValueOf(fn.Fn)
			if rv.Kind() != reflect.Func {
				panic(fmt.Sprintf("define %s type %T is not function", fn.Name, fn))
			}
			conf.DefineFuncs.Define(fn.Name, &consts.DefineFunc{
				Name:   fn.Name,
				NumIn:  rv.Type().NumIn(),
				NumOut: rv.Type().NumOut(),
				Fn:     rv,
			})
		}
	}
}

func BlockSystemFunc(name string) CompileOption {
	return func(conf *conf.CompileConf) {
		conf.DefineFuncs.BlockSystemFunc(name)
	}
}

func BlockAllSystemFunc() CompileOption {
	return func(conf *conf.CompileConf) {
		conf.DefineFuncs.BlockAllSystemFunc()
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

func NoCache() RunOption {
	return func(conf *conf.RunConf) {
		conf.MethodIndexCache = false
		conf.FieldIndexCache = false
	}
}

func Compile(code string, ops ...CompileOption) (*vm.Code, error) {
	p := compile.NewGsCompiler()
	var config = conf.Default()
	for _, op := range ops {
		op(&config)
	}
	ret, err := p.Compile(antlr.NewInputStream(code), &config)
	if err == nil && config.DumpCode {
		fmt.Println(p.Dump())
	}
	return ret, err
}

func Run(code *vm.Code, env any, ops ...RunOption) (ret *Result, err error) {
	config := conf.DefaultRunConf()
	for _, op := range ops {
		op(&config)
	}
	interpreter := vm.NewInterpreter(code, env, &config)
	err = interpreter.Run()
	if err != nil {
		return nil, err
	}
	return &Result{value: interpreter.Result}, nil
}

func Eval(program string, env any, ops ...any) (ret *Result, err error) {
	var compileOps []CompileOption
	var runOps []RunOption
	for _, op := range ops {
		switch op := op.(type) {
		case CompileOption:
			compileOps = append(compileOps, op)
		case RunOption:
			runOps = append(runOps, op)
		default:
			panic(fmt.Sprintf("EvalOption type %T is not CompileOption or RunOption", op))
		}
	}
	code, err := Compile(program, compileOps...)
	if err != nil {
		return nil, err
	}
	return Run(code, env, runOps...)
}
