package conf

import (
	"io"
	"os"
	"reflect"

	"github.com/ycl2018/gs/consts"
)

type CompileConf struct {
	Env            any
	Optimize       bool
	DumpCode       bool
	DefineFuncs    DefineFuncManager
	TypesAvailable DefineTypesManager
}

type RunConf struct {
	Out                     io.Writer
	Trace                   bool
	DisableFieldIndexCache  bool
	DisableMethodIndexCache bool
}

func Default() CompileConf {
	return CompileConf{
		Optimize:       true,
		DefineFuncs:    NewDefineFuncManager(),
		TypesAvailable: NewDefineTypesManager(),
	}
}

func DefaultRunConf() RunConf {
	return RunConf{
		Out: os.Stdout,
	}
}

type DefineFuncManager struct {
	userDefine    map[string]*consts.DefineFunc
	systemDefine  map[string]*consts.DefineFunc
	blockedSystem map[string]bool
}

func NewDefineFuncManager() DefineFuncManager {
	return DefineFuncManager{
		userDefine:    map[string]*consts.DefineFunc{},
		systemDefine:  consts.Predefine,
		blockedSystem: map[string]bool{},
	}
}

func (m DefineFuncManager) GetFunc(name string) *consts.DefineFunc {
	if f, ok := m.userDefine[name]; ok {
		return f
	}
	if m.blockedSystem[name] {
		return nil
	}
	if f, ok := m.systemDefine[name]; ok {
		return f
	}
	return nil
}

func (m DefineFuncManager) BlockSystemFunc(name string) {
	m.blockedSystem[name] = true
}

func (m DefineFuncManager) BlockAllSystemFunc() {
	m.systemDefine = map[string]*consts.DefineFunc{}
}

func (m DefineFuncManager) Define(name string, f *consts.DefineFunc) {
	if _, ok := m.systemDefine[name]; ok {
		panic("can not override system func " + name)
	}
	m.userDefine[name] = f
}

type DefineTypesManager struct {
	system     map[string]reflect.Type
	userDefine map[string]reflect.Type
}

func NewDefineTypesManager() DefineTypesManager {
	return DefineTypesManager{
		system:     consts.PredefineTypes,
		userDefine: map[string]reflect.Type{},
	}
}

func (m DefineTypesManager) GetType(name string) reflect.Type {
	if t, ok := m.system[name]; ok {
		return t
	}
	if t, ok := m.userDefine[name]; ok {
		return t
	}
	return nil
}

func (m DefineTypesManager) Define(name string, t reflect.Type) {
	if _, ok := m.system[name]; ok {
		panic("can not override system type " + name)
	}
	m.userDefine[name] = t
}
