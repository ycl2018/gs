package compile

import (
	"fmt"
	"reflect"
	"strings"

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
	s.Write(consts.InstrLoadEnv)
	accessors := qid.AllAccessor()
	if len(accessors) == 0 {
		return
	}
	var curType = s.Env.RType
	var brNils []*consts.StackInstr
	var indexId strings.Builder
	indexId.WriteString("$")
	curType = dePointer(curType)
	for i := 0; i < len(accessors); i++ {
		curType = dePointer(curType)
		var fieldPath []*reflect.StructField
		pa, ok := accessors[i].(*gen.PropertyAccessContext)
		// for consistent struct field load
		for ok && curType.Kind() == reflect.Struct && i < len(accessors) && pa.DOT() != nil {
			fieldName := pa.ID().GetText()
			indexId.WriteString("." + fieldName)
			f, err := s.Env.IndexField(fieldName, curType)
			if err != nil {
				s.Log.ErrorToken(qid.GetStart(), err.Error())
				return
			}
			fieldPath = append(fieldPath, f)
			curType = dePointer(f.Type)
			i++
			if i < len(accessors) {
				pa, ok = accessors[i].(*gen.PropertyAccessContext)
			}
		}
		if len(fieldPath) > 0 {
			s.Write(consts.InstrRFByIndex, defineFieldIndexConst(indexId.String(), fieldPath, s.GlobalScope).GetAddress())
			fieldPath = fieldPath[:0]
			if i > len(accessors)-1 {
				return
			}
		}
		switch a := accessors[i].(type) {
		case *gen.PropertyAccessContext:
			fieldName := a.ID().GetText()
			indexId.WriteString("." + fieldName)
			curType = dePointer(curType)
			if a.SAFE_DOT() != nil {
				brNil := consts.NewStackInstr(consts.InstrBRNil, placeholder)
				brNils = append(brNils, brNil)
				s.WriteInstr(brNil)
			}
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
		case *gen.IndexAccessContext:
			if a.SAFE_LBRACK() != nil {
				brNil := consts.NewStackInstr(consts.InstrBRNil, placeholder)
				brNils = append(brNils, brNil)
				s.WriteInstr(brNil)
			}
			switch t := a.GetChild(1).(type) {
			case *gen.ExprContext:
				indexId.WriteString("[" + t.GetText() + "]")
				t.Accept(s)
			case *gen.SliceExprContext:
				indexId.WriteString("[" + t.GetText() + "]")
				t.Accept(s)
			}
			switch curType.Kind() {
			case reflect.Map:
				s.Write(consts.InstrRMapIndex)
				curType = curType.Elem()
			case reflect.Array, reflect.Slice, reflect.String:
				s.Write(consts.InstrIndexLoad)
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
	return
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
	s.Write(consts.InstrLoadEnv)
	accessors := qid.AllAccessor()
	if len(accessors) == 0 {
		s.Write(consts.InstrRSet)
		return
	}
	var curType = s.Env.RType
	var indexId strings.Builder
	indexId.WriteString("$")
	curType = dePointer(curType)
	for i := 0; i < len(accessors); i++ {
		curType = dePointer(curType)
		var fieldPath []*reflect.StructField
		pa, ok := accessors[i].(*gen.PropertyAccessContext)
		// for consistent struct field load
		for ok && curType.Kind() == reflect.Struct && i < len(accessors)-1 && pa.DOT() != nil {
			fieldName := pa.ID().GetText()
			indexId.WriteString("." + fieldName)
			f, err := s.Env.IndexField(fieldName, curType)
			if err != nil {
				s.Log.ErrorToken(qid.GetStart(), err.Error())
				return
			}
			fieldPath = append(fieldPath, f)
			curType = dePointer(f.Type)
			i++
			if i < len(accessors) {
				pa, ok = accessors[i].(*gen.PropertyAccessContext)
			}
		}
		if len(fieldPath) > 0 {
			s.Write(consts.InstrRFByIndex, defineFieldIndexConst(indexId.String(), fieldPath, s.GlobalScope).GetAddress())
			fieldPath = fieldPath[:0]
			if i > len(accessors)-1 {
				return
			}
		}
		switch a := accessors[i].(type) {
		case *gen.PropertyAccessContext:
			fieldName := a.ID().GetText()
			indexId.WriteString("." + fieldName)
			curType = dePointer(curType)
			if a.SAFE_DOT() != nil {
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
			switch curType.Kind() {
			case reflect.Struct:
				if i != len(accessors)-1 {
					panic("unexpected code")
				}
				f, err := s.Env.IndexField(fieldName, curType)
				if err != nil {
					s.Log.ErrorToken(qid.GetStart(), err.Error())
					return
				}
				s.Write(consts.InstrRSetField, defineFieldIndexConst(indexId.String(), []*reflect.StructField{f}, s.GlobalScope).GetAddress())
			case reflect.Interface:
				// interface Load
				s.Write(consts.InstrFLoad, defineStringConst(fieldName, s.GlobalScope).GetAddress())
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		case *gen.IndexAccessContext:
			if a.SAFE_LBRACK() != nil {
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
			switch t := a.GetChild(1).(type) {
			case *gen.ExprContext:
				indexId.WriteString("[" + t.GetText() + "]")
				t.Accept(s)
			case *gen.SliceExprContext:
				indexId.WriteString("[" + t.GetText() + "]")
				if i == len(accessors)-1 {
					s.Log.ErrorToken(a.GetStart(), "syntax error:can't assign to slice split")
					return
				}
				t.Accept(s)
			}
			switch curType.Kind() {
			case reflect.Map:
				if i == len(accessors)-1 {
					s.Write(consts.InstrRSetMapIndex)
				} else {
					s.Write(consts.InstrRMapIndex)
					curType = curType.Elem()
				}
			case reflect.Array, reflect.Slice, reflect.String:
				curType = curType.Elem()
				if i == len(accessors)-1 {
					s.Write(consts.InstrRIndexStore)
				} else {
					s.Write(consts.InstrRIndex)
				}
			case reflect.Interface:
				if i == len(accessors)-1 {
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
	return
}

func dePointer(curType reflect.Type) reflect.Type {
	if curType.Kind() == reflect.Pointer {
		curType = curType.Elem()
	}
	return curType
}
