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

// CompileOption is the option for compiling.
type CompileOption func(conf *conf.CompileConf)

// Env sets the environment for compiling.
func Env(env any) CompileOption {
	return func(conf *conf.CompileConf) {
		conf.Env = env
	}
}

// DumpCode dumps the compiled code.
func DumpCode() CompileOption {
	return func(conf *conf.CompileConf) {
		conf.DumpCode = true
	}
}

// Func is the function to define.
type Func struct {
	Name string
	Fn   any
}

// FastFunc is the fast function to define. NumIn and NumOut are the number of input and output parameters.
type FastFunc struct {
	Name   string
	NumIn  int
	NumOut int
	Fn     func([]any) []any
}

// DefineFast defines the fast functions.
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

// DefineFuncs defines the functions.
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

// BlockSystemFunc blocks the system function.
func BlockSystemFunc(name string) CompileOption {
	return func(conf *conf.CompileConf) {
		conf.DefineFuncs.BlockSystemFunc(name)
	}
}

// BlockAllSystemFunc blocks all system functions.
func BlockAllSystemFunc() CompileOption {
	return func(conf *conf.CompileConf) {
		conf.DefineFuncs.BlockAllSystemFunc()
	}
}

// NoOptimize disables optimize.
func NoOptimize() CompileOption {
	return func(conf *conf.CompileConf) {
		conf.Optimize = false
	}
}

// AddTypes adds the types that can be used in the func newFromType.
func AddTypes(types map[string]any) CompileOption {
	return func(conf *conf.CompileConf) {
		for k, v := range types {
			conf.TypesAvailable.Define(k, reflect.TypeOf(v))
		}
	}
}

// RunOption is the option for running.
type RunOption func(conf *conf.RunConf)

// Trace enables trace info. useful for debugging.
func Trace() RunOption {
	return func(conf *conf.RunConf) {
		conf.Trace = true
	}
}

// Output sets the output writer. default is os.Stdout.
func Output(w io.Writer) RunOption {
	return func(conf *conf.RunConf) {
		conf.Out = w
	}
}

// StackSize sets the stack size. default is 64.
func StackSize(size int) RunOption {
	return func(conf *conf.RunConf) {
		conf.StackSize = size
	}
}

// NoCache disables method and field index cache.
func NoCache() RunOption {
	return func(conf *conf.RunConf) {
		conf.MethodIndexCache = false
		conf.FieldIndexCache = false
	}
}

// Compile compiles the program.
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

// Run runs the program.
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

// Eval compiles and runs the program in one step.
func Eval(program string, env any, ops ...any) (ret *Result, err error) {
	var compileOps = []CompileOption{
		NoOptimize(),
	}
	var runOps = []RunOption{
		NoCache(),
	}
	compileOps = append(compileOps, Env(env))
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
