package compile

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycl2018/gs/consts"
	"github.com/ycl2018/gs/gen"
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
	isMap := rt.Kind() == reflect.Map
	isSlice := rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array
	isStruct := rt.Kind() == reflect.Struct
	isPtrStruct := rt.Kind() == reflect.Ptr && rt.Elem().Kind() == reflect.Struct
	valid := isMap || isSlice || isStruct || isPtrStruct
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
	s.Write(consts.InstrLoadEnv)
	if len(query) == 0 {
		return
	}
	var curType = s.Env.RType
	var brNils []*consts.StackInstr
	var indexId strings.Builder
	indexId.WriteString("$")
	for i := 0; i < len(query); i++ {
		curType = dePointer(curType)
		var fieldPath []*reflect.StructField
		// short path for struct field load
		for curType.Kind() == reflect.Struct && i < len(query) && query[i] == gen.GsLexerDOT {
			indexId.WriteString("." + ids[i+1])
			f, err := s.Env.IndexField(ids[i+1], curType)
			if err != nil {
				s.Log.ErrorToken(qid.GetStart(), err.Error())
				return
			}
			fieldPath = append(fieldPath, f)
			curType = dePointer(f.Type)
			i++
		}
		if len(fieldPath) > 0 {
			s.Write(consts.InstrRFByIndex, defineFieldIndexConst(indexId.String(), fieldPath, s.GlobalScope).GetAddress())
			fieldPath = fieldPath[:0]
			if i > len(query)-1 {
				return
			}
		}
		switch query[i] {
		case gen.GsLexerSAFE_DOT:
			// if true, then fieldLoad
			brNil := consts.NewStackInstr(consts.InstrBRNil, placeholder)
			brNils = append(brNils, brNil)
			s.WriteInstr(brNil)
			indexId.WriteString("." + ids[i+1])
			// must be safe dot
			fieldName := ids[i+1]
			curType = dePointer(curType)
			switch curType.Kind() {
			case reflect.Struct:
				f, err := s.Env.IndexField(fieldName, curType)
				if err != nil {
					s.Log.ErrorToken(qid.GetStart(), err.Error())
					return
				}
				curType = f.Type
				s.Write(consts.InstrRFByIndex, defineFieldIndexConst(indexId.String(), []*reflect.StructField{f}, s.GlobalScope).GetAddress())
			case reflect.Interface:
				// interface Load
				s.Write(consts.InstrFLoad, defineStringConst(fieldName, s.GlobalScope).GetAddress())
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		case gen.GsLexerSAFE_LBRACK:
			brNil := consts.NewStackInstr(consts.InstrBRNil, placeholder)
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
				s.Write(consts.InstrRMapIndex)
				curType = curType.Elem()
			case reflect.Array, reflect.Slice, reflect.String:
				s.Write(consts.InstrRIndex)
				curType = curType.Elem()
			case reflect.Interface:
				s.Write(consts.InstrIndexLoad)
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		}
	}
	s.FillTarget(brNils...)
}

// TODO: 优化loadQid/storeQid,如果一直是fieldLoad，一步到位，合并所有path，一次性加载

func (s *StackCompileVisitor) storeQidToEnv(qid gen.IQidContext) {
	env := qid.GetChild(0).(*gen.PrimaryContext).ENV()
	if env == nil {
		panic("not env")
	}
	if s.Env.Kind == reflect.Struct {
		s.Log.ErrorToken(qid.GetStart(), "Env type is struct, can't assign,try use pointer instead")
		return
	}
	var ids []string
	var query []int // tokenType
	var qExprs []*gen.ExprContext
	var j int // index of qExprs
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
	s.Write(consts.InstrLoadEnv)
	if len(query) == 0 {
		s.Write(consts.InstrRSet)
		return
	}
	var curType = s.Env.RType
	var indexId strings.Builder
	indexId.WriteString("$")
	for i := 0; i < len(query); i++ {
		curType = dePointer(curType)
		var fieldPath []*reflect.StructField
		// short path for fieldLoad
		for query[i] == gen.GsLexerDOT && curType.Kind() == reflect.Struct && i < len(query)-1 {
			indexId.WriteString("." + ids[i+1])
			f, err := s.Env.IndexField(ids[i+1], curType)
			if err != nil {
				s.Log.ErrorToken(qid.GetStart(), err.Error())
				return
			}
			fieldPath = append(fieldPath, f)
			curType = dePointer(f.Type)
			i++
		}
		if len(fieldPath) > 0 {
			s.Write(consts.InstrRFByIndex, defineFieldIndexConst(indexId.String(), fieldPath, s.GlobalScope).GetAddress())
			fieldPath = fieldPath[:0]
		}
		switch query[i] {
		case gen.GsLexerDOT:
			// fieldLoad
			indexId.WriteString("." + ids[i+1])
			fieldName := ids[i+1]
			switch curType.Kind() {
			case reflect.Struct:
				// must be the last field in path
				if i != len(query)-1 {
					panic("field path must be the last in assign left side")
				}
				f, err := s.Env.IndexField(fieldName, curType)
				if err != nil {
					s.Log.ErrorToken(qid.GetStart(), err.Error())
					return
				}
				s.Write(consts.InstrRSetField, defineFieldIndexConst(indexId.String(), []*reflect.StructField{f}, s.GlobalScope).GetAddress())
				return
			case reflect.Interface:
				// interface Load
				s.Write(consts.InstrFLoad, defineStringConst(fieldName, s.GlobalScope).GetAddress())
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		case gen.GsLexerLBRACK:
			// arrayLoad/mapLoad
			expr := qExprs[j]
			j++
			expr.Accept(s)
			indexId.WriteString("[" + expr.GetText() + "]")
			switch curType.Kind() {
			case reflect.Map:
				if i == len(query)-1 {
					s.Write(consts.InstrRSetMapIndex)
				} else {
					s.Write(consts.InstrRMapIndex)
					curType = curType.Elem()
				}
			case reflect.Array, reflect.Slice, reflect.String:
				s.Write(consts.InstrRIndex)
				curType = curType.Elem()
				if i == len(query)-1 {
					s.Write(consts.InstrRSet)
				}
			case reflect.Interface:
				if i == len(query)-1 {
					s.Write(consts.InstrIndexStore)
				} else {
					s.Write(consts.InstrIndexLoad)
				}
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		}
	}
}

func dePointer(curType reflect.Type) reflect.Type {
	if curType.Kind() == reflect.Pointer {
		curType = curType.Elem()
	}
	return curType
}
