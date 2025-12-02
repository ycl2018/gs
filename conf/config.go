package conf

import (
	"io"

	"github.com/ycl2018/gs/consts"
)

type CompileConf struct {
	Env         any
	Optimize    bool
	DumpCode    bool
	DefineFuncs DefineFuncManager
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
		DefineFuncs: NewDefineFuncManager(),
	}
}

func DefaultRunConf() RunConf {
	return RunConf{
		Out:              io.Discard,
		FieldIndexCache:  true,
		MethodIndexCache: true,
	}
}
