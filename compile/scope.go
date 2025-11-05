package compile

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ycl2018/gs/vm"
)

const (
	GlobalScopeName = "global"
	LocalScopeName  = "local"
)

type Scope interface {
	Resolve(string) Symbol
	GetName() string
	Define(symbol Symbol)
	ParentScope() Scope
}

type GlobalScope struct {
	Symbols           map[string]Symbol
	Consts            map[string]*ConstSymbol
	LocalVarAllocator uint16
}

func NewGlobalScope() *GlobalScope {
	for i, j := 0, 0; i < 1 && j < 10; i, j = i+1, j+1 {

	}
	return &GlobalScope{
		Symbols: map[string]Symbol{},
		Consts:  map[string]*ConstSymbol{},
	}
}

func (g *GlobalScope) Define(symbol Symbol) {
	switch s := symbol.(type) {
	case *ConstSymbol:
		s.SetScope(g)
		name := s.Name
		if g.Consts[name] != nil {
			return
		}
		s.SetAddress(len(g.Consts))
		g.Consts[s.GetName()] = s
	case *StructSymbol:
		if g.Symbols[symbol.GetName()] != nil {
			return
		}
		g.Symbols[symbol.GetName()] = symbol
		symbol.SetScope(g)
		// fields: structNameIndex+fieldName1Index+fieldName2Index+...
		fields := make([]int, len(s.Fields)+1)
		// trimPrefix: index::
		if index := strings.Index(s.Name, "::"); index != -1 {
			index := strings.Index(s.Name, "::") + 2
			fields[0] = defineStringConst(s.Name[index:], g).GetAddress()
		} else {
			fields[0] = defineStringConst(s.Name, g).GetAddress()
		}
		for i, field := range s.Fields {
			sym := defineStringConst(field.GetName(), g)
			fields[i+1] = sym.GetAddress()
		}
		constSym := &ConstSymbol{
			Name:   s.Name,
			Kind:   vm.ConstStruct,
			Fields: fields,
		}
		g.Define(constSym) // 存储结构体常量到全局常量池
		s.SetAddress(constSym.GetAddress())
	case *FunctionSymbol:
		if g.Symbols[symbol.GetName()] != nil {
			return
		}
		g.Symbols[symbol.GetName()] = symbol
		symbol.SetScope(g)
		funcNameIndex := defineStringConst(s.Name, g).GetAddress()
		// 函数名，参数数，局部变量数，代码地址
		fields := []int{funcNameIndex, int(len(s.FormalArgs)), int(s.LocalNums()), s.CodeAddr}
		constSym := &ConstSymbol{
			Name:   s.Name,
			Kind:   vm.ConstFunc,
			Fields: fields,
		}
		g.Define(constSym) // 存储函数常量到全局常量池
		s.SetAddress(constSym.GetAddress())
	default:
		if g.Symbols[symbol.GetName()] != nil {
			return
		}
		symbol.SetAddress(int(g.LocalVarAllocator))
		g.LocalVarAllocator++
		g.Symbols[symbol.GetName()] = symbol
		symbol.SetScope(g)
	}
}

func (g *GlobalScope) DefineOrGetConst(symbol *ConstSymbol) (Symbol, bool) {
	name := symbol.Name
	if g.Consts[name] != nil {
		return g.Consts[name], false
	}
	g.Define(symbol)
	return symbol, true
}

func (g *GlobalScope) Resolve(name string) Symbol {
	if v, ok := g.Symbols[name]; ok {
		return v
	}
	return nil
}

func (g *GlobalScope) GetName() string {
	return GlobalScopeName
}
func (g *GlobalScope) ParentScope() Scope {
	return nil
}

func (g *GlobalScope) Dump() string {
	var sb strings.Builder
	// 全局符号
	sb.WriteString(fmt.Sprintf(".globals: %d\n", g.LocalVarAllocator))
	// 全局变量
	// 按地址排序
	var syms []Symbol
	for _, s := range g.Symbols {
		syms = append(syms, s)
	}
	sort.Slice(syms, func(i, j int) bool {
		return syms[i].GetAddress() < syms[j].GetAddress()
	})
	for _, s := range syms {
		sb.WriteString(DumpSymbol(s, nil))
	}
	// 全局常量池
	sb.WriteString(fmt.Sprintf("\nConstant Pool:%d\n", len(g.Consts)))
	// 按地址排序
	var consts []Symbol
	for _, s := range g.Consts {
		consts = append(consts, s)
	}
	sort.Slice(consts, func(i, j int) bool {
		return consts[i].GetAddress() < consts[j].GetAddress()
	})
	for _, s := range consts {
		sb.WriteString(DumpSymbol(s, consts))
	}
	return sb.String()
}

type LocalScope struct {
	ID                int32
	Symbols           map[string]Symbol
	EnclosingScope    Scope
	BaseAllocAddr     int32
	LocalVarAllocator int32
}

func (l *LocalScope) Define(symbol Symbol) {
	switch s := symbol.(type) {
	case *ConstSymbol:
		panic("const symbol define in local scope")
	case *StructSymbol:
		if l.Symbols[symbol.GetName()] != nil {
			panic(fmt.Sprintf("duplicate struct %s in local scope", symbol.GetName()))
			return
		}
		// 同时存储在全局 scope 中
		globalScope := l.EnclosingScope
		for globalScope != nil && globalScope.GetName() != GlobalScopeName {
			globalScope = globalScope.ParentScope()
		}
		gName := fmt.Sprintf("%d::%s", l.ID, symbol.GetName())
		globalScope.Define(&StructSymbol{
			Fields: s.Fields,
			// 通过 scope ID 来避免命名冲突
			BaseSymbol: BaseSymbol{
				Name:        gName,
				DefineToken: s.DefineToken,
			},
		})
		sym := globalScope.Resolve(gName).(*StructSymbol)
		symbol.SetAddress(sym.GetAddress()) // 指向全局常量池中的结构体定义
		l.Symbols[symbol.GetName()] = symbol
		symbol.SetScope(l)
	default:
		if l.Symbols[symbol.GetName()] != nil {
			return
		}
		symbol.SetAddress(int(l.BaseAllocAddr) + int(l.LocalVarAllocator))
		l.LocalVarAllocator++
		l.Symbols[symbol.GetName()] = symbol
		symbol.SetScope(l)
	}
}

func (l *LocalScope) Resolve(name string) Symbol {
	if v, ok := l.Symbols[name]; ok {
		return v
	}
	if l.EnclosingScope != nil {
		return l.EnclosingScope.Resolve(name)
	}
	return nil
}

func (l *LocalScope) GetName() string {
	return LocalScopeName
}
func (l *LocalScope) ParentScope() Scope {
	return l.EnclosingScope
}
