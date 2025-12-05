package compile

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ycl2018/gs/consts"
	"github.com/ycl2018/gs/gen"
)

type Env struct {
	RType       reflect.Type
	Kind        reflect.Kind
	Cache       map[reflect.Type]map[string]*reflect.StructField
	MethodCache map[reflect.Type]map[string]*reflect.Method
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
		RType:       rt,
		Kind:        rt.Kind(),
		Cache:       make(map[reflect.Type]map[string]*reflect.StructField),
		MethodCache: make(map[reflect.Type]map[string]*reflect.Method),
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
	if ok {
		if !f.IsExported() {
			return nil, fmt.Errorf("field %s is not exported for type %s", field, from)
		}
		if cache, ok := e.Cache[from]; ok {
			cache[field] = &f
		} else {
			e.Cache[from] = map[string]*reflect.StructField{field: &f}
		}
		return &f, nil
	}
	return nil, fmt.Errorf("field %s not found for type %s", field, from)
}

func (e *Env) IndexMethod(method string, from reflect.Type) (*reflect.Method, error) {
	if cache, ok := e.MethodCache[from]; ok {
		if f, ok := cache[method]; ok {
			return f, nil
		}
	}
	kind := from.Kind()
	if kind != reflect.Struct {
		return nil, fmt.Errorf("invalid type %s", from.String())
	}
	f, ok := from.MethodByName(method)
	if ok && f.IsExported() {
		if cache, ok := e.MethodCache[from]; ok {
			cache[method] = &f
		} else {
			e.MethodCache[from] = map[string]*reflect.Method{method: &f}
		}
		return &f, nil
	}
	f, ok = reflect.PointerTo(from).MethodByName(method)
	if ok {
		if !f.IsExported() {
			return nil, fmt.Errorf("method %s is not exported for type %s", method, from)
		}
		if cache, ok := e.MethodCache[from]; ok {
			cache[method] = &f
		} else {
			e.MethodCache[from] = map[string]*reflect.Method{method: &f}
		}
		return &f, nil
	}
	return nil, fmt.Errorf("method %s not found for type %s", method, from)
}

func (s *StackCompileVisitor) loadQidFromEnv(qid gen.IQidContext) {
	s.Write(consts.InstrLoadEnv, qid.GetStart())
	accessors := qid.AllAccessor()
	if len(accessors) == 0 {
		return
	}
	var curType = s.Env.RType
	var indexId strings.Builder
	indexId.WriteString(EnvText)
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
			s.Write(consts.InstrRFByIndex, qid.GetStart(), defineFieldIndexConst(indexId.String(), fieldPath, s.GlobalScope).GetAddress())
			fieldPath = fieldPath[:0]
			if i > len(accessors)-1 {
				return
			}
			indexId.WriteString(".")
		}
		switch a := accessors[i].(type) {
		case *gen.PropertyAccessContext:
			fieldName := a.ID().GetText()
			indexId.WriteString("." + fieldName)
			curType = dePointer(curType)
			switch curType.Kind() {
			case reflect.Struct:
				f, err := s.Env.IndexField(fieldName, curType)
				if err != nil {
					s.Log.ErrorToken(qid.GetStart(), err.Error())
					return
				}
				curType = f.Type
				s.Write(consts.InstrRFByIndex, a.GetStart(), defineFieldIndexConst(indexId.String(), []*reflect.StructField{f}, s.GlobalScope).GetAddress())
			case reflect.Interface:
				// interface Load
				if curType.NumMethod() > 0 {
					s.Log.ErrorToken(pa.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
					return
				}
				s.Write(consts.InstrFLoad, a.GetStart(), defineStringConst(fieldName, s.GlobalScope).GetAddress())
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		case *gen.IndexAccessContext:
			switch t := a.GetChild(1).(type) {
			case *gen.SliceExprContext:
				if !canSliceSplit(curType) {
					s.Log.ErrorToken(a.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
					return
				}
				indexId.WriteString("[" + t.GetText() + "]")
				t.Accept(s)
			case *gen.ExprContext:
				indexId.WriteString("[" + t.GetText() + "]")
				t.Accept(s)
				switch curType.Kind() {
				case reflect.Map:
					s.Write(consts.InstrRMapIndex, a.GetStart())
					curType = curType.Elem()
				case reflect.Array, reflect.Slice, reflect.String:
					s.Write(consts.InstrIndexLoad, a.GetStart())
					curType = curType.Elem()
				case reflect.Interface:
					if curType.NumMethod() > 0 {
						s.Log.ErrorToken(pa.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
						return
					}
					s.Write(consts.InstrIndexLoad, a.GetStart())
				default:
					s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
					return
				}
			}
		}
	}
	return
}

func (s *StackCompileVisitor) storeQidToEnv(qid gen.IQidContext) {
	if s.Env.Kind == reflect.Struct {
		s.Log.ErrorToken(qid.GetStart(), "Env type is struct, can't assign,try use pointer instead")
		return
	}
	s.Write(consts.InstrLoadEnv, qid.GetStart())
	accessors := qid.AllAccessor()
	if len(accessors) == 0 {
		s.Write(consts.InstrRSet, qid.GetStart())
		return
	}
	var curType = s.Env.RType
	var indexId strings.Builder
	indexId.WriteString(EnvText)
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
			s.Write(consts.InstrRFByIndex, qid.GetStart(), defineFieldIndexConst(indexId.String(), fieldPath, s.GlobalScope).GetAddress())
			fieldPath = fieldPath[:0]
			if i > len(accessors)-1 {
				return
			}
			indexId.WriteString(".")
		}
		switch a := accessors[i].(type) {
		case *gen.PropertyAccessContext:
			fieldName := a.ID().GetText()
			indexId.WriteString("." + fieldName)
			curType = dePointer(curType)
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
				s.Write(consts.InstrRSetField, a.GetStart(), defineFieldIndexConst(indexId.String(), []*reflect.StructField{f}, s.GlobalScope).GetAddress())
			case reflect.Interface:
				// interface Load
				if i == len(accessors)-1 {
					s.Write(consts.InstrFStore, a.GetStart(), defineStringConst(fieldName, s.GlobalScope).GetAddress())
				} else {
					if curType.NumMethod() > 0 {
						s.Log.ErrorToken(pa.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
						return
					}
					s.Write(consts.InstrFLoad, a.GetStart(), defineStringConst(fieldName, s.GlobalScope).GetAddress())
				}
			default:
				s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		case *gen.IndexAccessContext:
			switch t := a.GetChild(1).(type) {
			case *gen.ExprContext:
				indexId.WriteString("[" + t.GetText() + "]")
				t.Accept(s)
				switch curType.Kind() {
				case reflect.Map:
					if i == len(accessors)-1 {
						s.Write(consts.InstrRSetMapIndex, a.GetStart())
					} else {
						s.Write(consts.InstrRMapIndex, a.GetStart())
						curType = curType.Elem()
					}
				case reflect.Array, reflect.Slice, reflect.String:
					curType = curType.Elem()
					if i == len(accessors)-1 {
						s.Write(consts.InstrRIndexStore, a.GetStart())
					} else {
						s.Write(consts.InstrRIndex, a.GetStart())
					}
				case reflect.Interface:
					if i == len(accessors)-1 {
						s.Write(consts.InstrIndexStore, a.GetStart())
					} else {
						if curType.NumMethod() > 0 {
							s.Log.ErrorToken(pa.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
							return
						}
						s.Write(consts.InstrIndexLoad, a.GetStart())
					}
				default:
					s.Log.ErrorToken(qid.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
					return
				}
			case *gen.SliceExprContext:
				if !canSliceSplit(curType) {
					s.Log.ErrorToken(a.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
					return
				}
				indexId.WriteString("[" + t.GetText() + "]")
				if i == len(accessors)-1 {
					s.Log.ErrorToken(a.GetStart(), "syntax error:can't assign to slice split")
					return
				}
				t.Accept(s)
			}
		}
	}
	return
}

func (s *StackCompileVisitor) loadOuterFuncFromEnv(ctx *gen.OuterCallContext, accessors []gen.IAccessorContext) {
	s.Write(consts.InstrLoadEnv, ctx.GetStart())
	if len(accessors) == 0 {
		return
	}
	var curType = s.Env.RType
	var indexId strings.Builder
	indexId.WriteString(EnvText)
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
				s.Log.ErrorToken(pa.GetStart(), err.Error())
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
			s.Write(consts.InstrRFByIndex, ctx.GetStart(), defineFieldIndexConst(indexId.String(), fieldPath, s.GlobalScope).GetAddress())
			fieldPath = fieldPath[:0]
			if i > len(accessors)-1 {
				return
			}
			indexId.WriteString(".")
		}
		switch a := accessors[i].(type) {
		case *gen.PropertyAccessContext:
			fieldName := a.ID().GetText()
			indexId.WriteString("." + fieldName)
			curType = dePointer(curType)
			switch curType.Kind() {
			case reflect.Struct:
				if i != len(accessors)-1 {
					panic("unexpected code")
				}
				m, err := s.Env.IndexMethod(fieldName, curType)
				if m != nil {
					s.Write(consts.InstrMLoadByIndex, ctx.GetStart(), m.Index)
					return
				}
				// field
				f, err := s.Env.IndexField(fieldName, curType)
				if err != nil {
					s.Log.ErrorToken(a.GetStart(), err.Error())
					return
				}
				if f.Type.Kind() == reflect.Func || f.Type.Kind() == reflect.Interface {
					s.Write(consts.InstrRFByIndex, a.GetStart(), defineFieldIndexConst(indexId.String(), []*reflect.StructField{f}, s.GlobalScope).GetAddress())
					return
				}
				s.Log.ErrorToken(a.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			case reflect.Interface:
				// interface Load
				if i == len(accessors)-1 {
					if curType.NumMethod() > 0 {
						if methodByName, ok := curType.MethodByName(fieldName); ok {
							s.Write(consts.InstrMLoadByIndex, a.GetStart(), methodByName.Index)
						} else {
							s.Log.ErrorToken(pa.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
							return
						}
					} else {
						s.Write(consts.InstrMLoadByName, a.GetStart(), defineStringConst(fieldName, s.GlobalScope).GetAddress())
					}
				} else {
					if curType.NumMethod() > 0 {
						s.Log.ErrorToken(pa.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
						return
					}
					s.Write(consts.InstrFLoad, a.GetStart(), defineStringConst(fieldName, s.GlobalScope).GetAddress())
				}
			default:
				s.Log.ErrorToken(a.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
				return
			}
		case *gen.IndexAccessContext:
			switch t := a.GetChild(1).(type) {
			case *gen.ExprContext:
				indexId.WriteString("[" + t.GetText() + "]")
				t.Accept(s)
				switch curType.Kind() {
				case reflect.Map:
					s.Write(consts.InstrRMapIndex, a.GetStart())
					curType = curType.Elem()
				case reflect.Array, reflect.Slice, reflect.String:
					s.Write(consts.InstrRIndex, a.GetStart())
					curType = curType.Elem()
				case reflect.Interface:
					if curType.NumMethod() > 0 {
						s.Log.ErrorToken(pa.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
						return
					}
					s.Write(consts.InstrIndexLoad, a.GetStart())
				default:
					s.Log.ErrorToken(ctx.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
					return
				}
			case *gen.SliceExprContext:
				if !canSliceSplit(curType) {
					s.Log.ErrorToken(a.GetStart(), "syntax error:invalid %s on type:%s", indexId.String(), curType.String())
					return
				}
				indexId.WriteString("[" + t.GetText() + "]")
				if i == len(accessors)-1 {
					s.Log.ErrorToken(a.GetStart(), "syntax error:can't assign to slice split")
					return
				}
				t.Accept(s)
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

func canSliceSplit(curType reflect.Type) bool {
	return curType.Kind() == reflect.Slice || curType.Kind() == reflect.Array || curType.Kind() == reflect.String || (curType.Kind() == reflect.Interface && curType.NumMethod() == 0)
}
