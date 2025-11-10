package compile

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/gen"
	"github.com/ycl2018/gs/vm"
)

type Env struct {
	RType reflect.Type
	Kind  reflect.Kind
	Cache map[reflect.Type]map[string]*reflect.StructField
}

func NewEnv(v any) *Env {
	if v == nil {
		return nil
	}
	rt := reflect.TypeOf(v)
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	isMap := rt.Kind() == reflect.Map
	isSlice := rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array
	isStruct := rt.Kind() == reflect.Struct
	valid := isMap || isSlice || isStruct
	if !valid {
		panic("invalid type")
	}
	return &Env{
		RType: rt,
		Kind:  rt.Kind(),
		Cache: make(map[reflect.Type]map[string]*reflect.StructField),
	}
}

type FiledIndex struct {
	Index []int
	Name  string
	From  reflect.Kind
}

func (e *Env) IndexField(field string, from reflect.Type) (*reflect.StructField, error) {
	if cache, ok := e.Cache[from]; ok {
		if f, ok := cache[field]; ok {
			return f, nil
		}
	}

	kind := from.Kind()
	if kind != reflect.Struct {
		return nil, fmt.Errorf("invalid type %s", from.String())
	}
	f, ok := from.FieldByName(field)
	if ok && f.IsExported() {
		if cache, ok := e.Cache[from]; ok {
			cache[field] = &f
		} else {
			e.Cache[from] = map[string]*reflect.StructField{field: &f}
		}
		return &f, nil
	}
	return nil, fmt.Errorf("field %s not found", field)
}

func (s *StackCompileVisitor) loadQidFromEnv(qid gen.IQidContext) {
	env := qid.GetChild(0).(*gen.PrimaryContext).ENV()
	if env == nil {
		panic("not env")
	}
	var ids []string
	var query []int // tokenType
	var qExprs []*gen.ExprContext
	for i, child := range qid.GetChildren() {
		if i == 0 {
			ids = append(ids, "$")
			continue
		}
		if node, ok := child.(antlr.TerminalNode); ok {
			switch t := node.GetSymbol().GetTokenType(); t {
			case gen.GsLexerDOT, gen.GsLexerSAFE_DOT, gen.GsLexerLBRACK, gen.GsLexerSAFE_LBRACK:
				query = append(query, t)
			case gen.GsLexerID:
				ids = append(ids, node.GetText())
			}
		} else if e, ok := child.(*gen.ExprContext); ok {
			qExprs = append(qExprs, e)
		}
	}
	s.Write(vm.InstrLoadEnv)
	if len(query) == 0 {
		return
	}
	s.Write(vm.InstrRV)
	var curType = s.Env.RType
	var brNils []*vm.StackInstr
	var indexId strings.Builder
	indexId.WriteString("$")
	for i := 0; i < len(query); i++ {
		switch query[i] {
		case gen.GsLexerSAFE_DOT:
			// if true, then fieldLoad
			brNil := vm.NewStackInstr(vm.InstrBRNil, placeholder)
			brNils = append(brNils, brNil)
			s.WriteInstr(brNil)
			fallthrough
		case gen.GsLexerDOT:
			indexId.WriteString("." + ids[i+1])
			// fieldLoad
			fieldName := ids[i+1]
			if curType.Kind() == reflect.Pointer {
				curType = curType.Elem()
				s.Write(vm.InstrRElem)
			}
			switch curType.Kind() {
			case reflect.Struct:
				f, err := s.Env.IndexField(fieldName, curType)
				if err != nil {
					s.Log.ErrorToken(qid.GetStart(), err.Error())
					return
				}
				curType = f.Type
				s.Write(vm.InstrRFByIndex, defineFieldIndexConst(indexId.String(), f, s.GlobalScope).GetAddress())
				if i == len(query)-1 {
					s.Write(vm.InstrInterface)
				}
			case reflect.Interface:
				// interface Load
				s.Write(vm.InstrFLoad, defineStringConst(fieldName, s.GlobalScope).GetAddress())
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		case gen.GsLexerSAFE_LBRACK:
			brNil := vm.NewStackInstr(vm.InstrBRNil, placeholder)
			brNils = append(brNils, brNil)
			s.WriteInstr(brNil)
			fallthrough
		case gen.GsLexerLBRACK:
			// arrayLoad/mapLoad
			expr := qExprs[i]
			indexId.WriteString("[" + expr.GetText() + "]")
			expr.Accept(s)
			switch curType.Kind() {
			case reflect.Map:
				s.Write(vm.InstrRMapIndex)
				curType = curType.Elem()
			case reflect.Array, reflect.Slice, reflect.String:
				s.Write(vm.InstrRIndex)
				curType = curType.Elem()
			case reflect.Interface:
				s.Write(vm.InstrIndexLoad)
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
			if i == len(query)-1 {
				s.Write(vm.InstrInterface) // convert reflect.Value to interface
			}
		}
	}
	s.FillTarget(brNils...)
}

func (s *StackCompileVisitor) storeQidToEnv(qid gen.IQidContext) {
	env := qid.GetChild(0).(*gen.PrimaryContext).ENV()
	if env == nil {
		panic("not env")
	}
	var ids []string
	var query []int // tokenType
	var qExprs []*gen.ExprContext
	for i, child := range qid.GetChildren() {
		if i == 0 {
			ids = append(ids, "$")
			continue
		}
		if node, ok := child.(antlr.TerminalNode); ok {
			switch t := node.GetSymbol().GetTokenType(); t {
			case gen.GsLexerDOT, gen.GsLexerLBRACK:
				query = append(query, t)
			case gen.GsLexerSAFE_DOT, gen.GsLexerSAFE_LBRACK:
				s.Log.ErrorToken(node.GetSymbol(), "syntax error:can't use %s in assign left side", node.GetSymbol().GetText())
				return
			case gen.GsLexerID:
				ids = append(ids, node.GetText())
			}
		} else if e, ok := child.(*gen.ExprContext); ok {
			qExprs = append(qExprs, e)
		}
	}
	s.Write(vm.InstrLoadEnv)
	s.Write(vm.InstrRV)
	if len(ids) == 1 {
		s.Write(vm.InstrRSet)
		return
	}
	var curType = s.Env.RType
	var indexId strings.Builder
	indexId.WriteString("$")
	for i := 0; i < len(query)-1; i++ {
		switch query[i] {
		case gen.GsLexerDOT:
			// fieldLoad
			indexId.WriteString("." + ids[i+1])
			fieldName := ids[i+1]
			if curType.Kind() == reflect.Pointer {
				curType = curType.Elem()
				s.Write(vm.InstrRElem)
			}
			switch curType.Kind() {
			case reflect.Struct:
				f, err := s.Env.IndexField(fieldName, curType)
				if err != nil {
					s.Log.ErrorToken(qid.GetStart(), err.Error())
					return
				}
				curType = f.Type
				s.Write(vm.InstrRFByIndex, defineFieldIndexConst(indexId.String(), f, s.GlobalScope).GetAddress())
			case reflect.Interface:
				// interface Load
				s.Write(vm.InstrFLoad, defineStringConst(fieldName, s.GlobalScope).GetAddress())
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		case gen.GsLexerLBRACK:
			// arrayLoad/mapLoad
			expr := qExprs[i]
			expr.Accept(s)
			indexId.WriteString("[" + expr.GetText() + "]")
			switch curType.Kind() {
			case reflect.Map:
				s.Write(vm.InstrRMapIndex)
				curType = curType.Elem()
			case reflect.Array, reflect.Slice, reflect.String:
				s.Write(vm.InstrRIndex)
				curType = curType.Elem()
			case reflect.Interface:
				s.Write(vm.InstrIndexLoad)
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		}
	}
	// lastQuery
	switch query[len(query)-1] {
	case gen.GsLexerDOT:
		fieldName := ids[len(ids)-1]
		if curType.Kind() == reflect.Pointer {
			curType = curType.Elem()
			s.Write(vm.InstrRElem)
		}
		switch curType.Kind() {
		case reflect.Struct:
			f, err := s.Env.IndexField(fieldName, curType)
			if err != nil {
				s.Log.ErrorToken(qid.GetStart(), err.Error())
				return
			}
			curType = f.Type
			s.Write(vm.InstrRFByIndex, defineFieldIndexConst(indexId.String(), f, s.GlobalScope).GetAddress())
			s.Write(vm.InstrRSet) //
		case reflect.Interface:
			// interface Load
			s.Write(vm.InstrFStore, defineStringConst(fieldName, s.GlobalScope).GetAddress())
		default:
			s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
			return
		}
	case gen.GsLexerLBRACK:
		// arrayStore/mapStore
		expr := qExprs[len(qExprs)-1]
		expr.Accept(s)
		switch curType.Kind() {
		case reflect.Map:
			s.Write(vm.InstrRSetMapIndex)
		case reflect.Array, reflect.Slice, reflect.String:
			s.Write(vm.InstrRIndex)
			s.Write(vm.InstrRSet)
		case reflect.Interface:
			s.Write(vm.InstrIndexStore)
		default:
			s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
			return
		}
	}
}
