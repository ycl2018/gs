package consts

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/ycl2018/gs/gen"
)

// Predefined functions.
const (
	FnIn            = "in"
	FnIndex         = "index"
	FnSort          = "sort"
	FnHasPrefix     = "hasPrefix"
	FnHasSuffix     = "hasSuffix"
	FnTrim          = "trim"
	FnTrimPrefix    = "trimPrefix"
	FnTrimSuffix    = "trimSuffix"
	FnTrimSpace     = "trimSpace"
	FnTrimLeft      = "trimLeft"
	FnTrimRight     = "trimRight"
	FnToLower       = "toLower"
	FnToUpper       = "toUpper"
	FnSplit         = "split"
	FnJoin          = "join"
	FnNow           = "now"
	FnParseTime     = "parseTime"
	FnParseDuration = "parseDuration"
	FnAtoi          = "atoi"
	FnItoA          = "itoa"
	FnDuration      = "duration"
	FnToJSON        = "toJson"
	FnFromJSON      = "fromJson"
	FnUnmarshalJSON = "unmarshalJson"
)

// Predefine is the predefine functions. **DO NOT MODIFY**.
var Predefine = func() map[string]*DefineFunc {
	var ret = map[string]*DefineFunc{}
	define := func(name string, numIn, numOut int, fn any) {
		ret[name] = &DefineFunc{
			Name:   name,
			NumIn:  numIn,
			NumOut: numOut,
			Fast:   true,
			Fn:     fn,
		}
	}
	define(FnIn, 2, 1, In)
	define(FnIndex, 2, 1, Index)
	define(FnSort, 1, 0, Sort)
	define(FnHasPrefix, 2, 1, HasPrefix)
	define(FnHasSuffix, 2, 1, HasSuffix)
	define(FnTrim, 2, 1, Trim)
	define(FnTrimPrefix, 2, 1, TrimPrefix)
	define(FnTrimSuffix, 2, 1, TrimSuffix)
	define(FnTrimSpace, 1, 1, TrimSpace)
	define(FnTrimLeft, 2, 1, TrimLeft)
	define(FnTrimRight, 2, 1, TrimRight)
	define(FnToLower, 1, 1, toLower)
	define(FnToUpper, 1, 1, toUpper)
	define(FnSplit, 2, 1, Split)
	define(FnJoin, 2, 1, Join)
	define(FnNow, 0, 1, Now)
	define(FnParseTime, 2, 2, ParseTime)
	define(FnParseDuration, 1, 2, ParseDuration)
	define(FnAtoi, 1, 2, Atoi)
	define(FnItoA, 1, 1, Itoa)
	define(FnDuration, 1, 1, Duration)
	define(FnToJSON, 1, 2, ToJson)
	define(FnFromJSON, 1, 2, FromJson)
	define(FnUnmarshalJSON, 2, 1, UnmarshalJson)
	return ret
}()

var PredefineTypes = map[string]reflect.Type{
	"int":           reflect.TypeOf(0),
	"int8":          reflect.TypeOf(int8(0)),
	"int16":         reflect.TypeOf(int16(0)),
	"int32":         reflect.TypeOf(int32(0)),
	"int64":         reflect.TypeOf(int64(0)),
	"uint":          reflect.TypeOf(uint(0)),
	"uint8":         reflect.TypeOf(uint8(0)),
	"uint16":        reflect.TypeOf(uint16(0)),
	"uint32":        reflect.TypeOf(uint32(0)),
	"uint64":        reflect.TypeOf(uint64(0)),
	"float32":       reflect.TypeOf(float32(0)),
	"float64":       reflect.TypeOf(float64(0)),
	"string":        reflect.TypeOf(""),
	"uintptr":       reflect.TypeOf(uintptr(0)),
	"time.Time":     reflect.TypeOf(time.Time{}),
	"time.Duration": reflect.TypeOf(time.Duration(0)),
	"struct{}":      reflect.TypeOf(struct{}{}),
}

func In(args []any) []any {
	obj := args[0]
	key := args[1]
	if obj == nil {
		return []any{false}
	}
	switch obj := obj.(type) {
	case []any:
		for i := 0; i < len(obj); i++ {
			if gen.Eq(obj[i], key, false) {
				return []any{true}
			}
		}
		return []any{false}
	case map[string]string:
		if _, ok := obj[key.(string)]; ok {
			return []any{true}
		}
		return []any{false}
	case map[string]struct{}:
		if v, ok := key.(string); !ok {
			return []any{false}
		} else {
			if _, ok := obj[v]; ok {
				return []any{true}
			}
		}
		return []any{false}
	case map[int]struct{}:
		if v, ok := gen.TryToInt(key); !ok {
			return []any{false}
		} else {
			if _, ok := obj[v]; ok {
				return []any{true}
			}
		}
		return []any{false}
	case map[float64]struct{}:
		if v, ok := gen.TryToFloat64(key); !ok {
			return []any{false}
		} else {
			if _, ok := obj[v]; ok {
				return []any{true}
			}
		}
		return []any{false}
	case map[int64]struct{}:
		if v, ok := gen.TryToInt64(key); !ok {
			return []any{false}
		} else {
			if _, ok := obj[v]; ok {
				return []any{true}
			}
		}
		return []any{false}
	case []int:
		if key, ok := gen.TryToInt(key); !ok {
			return []any{false}
		} else {
			for _, v := range obj {
				if v == key {
					return []any{true}
				}
			}
		}
		return []any{false}
	case []int64:
		if key, ok := gen.TryToInt64(key); !ok {
			return []any{false}
		} else {
			for _, v := range obj {
				if v == key {
					return []any{true}
				}
			}
		}
		return []any{false}
	case []string:
		for _, v := range obj {
			if v == key.(string) {
				return []any{true}
			}
		}
		return []any{false}
	}
	rv := reflect.ValueOf(obj)
	kind := rv.Kind()
	switch kind {
	case reflect.Slice, reflect.Array, reflect.String:
		for i := 0; i < rv.Len(); i++ {
			if ok := gen.Eq(rv.Index(i).Interface(), key, false); ok {
				return []any{true}
			}
		}
		return []any{false}
	case reflect.Map:
		if key == nil {
			key = reflect.Zero(rv.Type().Key())
		}
		if rv.MapIndex(reflect.ValueOf(key)).IsValid() {
			return []any{true}
		}
		return []any{false}
	default:
		panic(fmt.Sprintf("in: not support type %T", obj))
	}
}

func Index(args []any) []any {
	obj := args[0]
	key := args[1]
	if obj == nil {
		return []any{false}
	}
	switch obj := obj.(type) {
	case []int:
		key := gen.ToInt(key)
		for i, v := range obj {
			if v == key {
				return []any{i}
			}
		}
		return []any{-1}
	case []string:
		for i, v := range obj {
			if v == key.(string) {
				return []any{i}
			}
		}
		return []any{-1}
	case string:
		return []any{strings.Index(obj, key.(string))}
	case []int64:
		key := gen.ToInt64(key)
		for i, v := range obj {
			if v == key {
				return []any{i}
			}
		}
		return []any{-1}
	case []any:
		for i, v := range obj {
			if ok := gen.Eq(v, key, false); ok {
				return []any{i}
			}
		}
		return []any{-1}
	}
	rv := reflect.ValueOf(obj)
	kind := rv.Kind()
	switch kind {
	case reflect.Slice, reflect.Array, reflect.String:
		for i := 0; i < rv.Len(); i++ {
			if ok := gen.Eq(rv.Index(i).Interface(), key, false); ok {
				return []any{i}
			}
		}
		return []any{-1}
	default:
		panic(fmt.Sprintf("index: not support type %T", obj))
	}
}

// Sort support basic kind slice or []any of all basic kind values that can compare
func Sort(args []any) []any {
	s := args[0]
	if s == nil {
		return nil
	}
	switch s := s.(type) {
	case []string:
		slices.Sort(s)
	case []int:
		slices.Sort(s)
	case []int8:
		slices.Sort(s)
	case []int16:
		slices.Sort(s)
	case []int32:
		slices.Sort(s)
	case []int64:
		slices.Sort(s)
	case []float32:
		slices.Sort(s)
	case []float64:
		slices.Sort(s)
	case []uint:
		slices.Sort(s)
	case []uint8:
		slices.Sort(s)
	case []uint16:
		slices.Sort(s)
	case []uint32:
		slices.Sort(s)
	case []uint64:
		slices.Sort(s)
	case []uintptr:
		slices.Sort(s)
	case []any:
		slices.SortFunc(s, func(a, b any) int {
			ok := gen.Lt(a, b)
			if ok {
				return -1
			}
			return 1
		})
	default:
		panic(fmt.Sprintf("sort: not support type %T,please use user define function to sort", s))
	}
	return nil
}

func HasPrefix(args []any) []any {
	obj := args[0]
	key := args[1]
	if obj == nil || key == nil {
		return []any{false}
	}
	switch obj := obj.(type) {
	case string:
		return []any{strings.HasPrefix(obj, key.(string))}
	default:
		panic(fmt.Sprintf("startWith: not support type %T", obj))
	}
}

func HasSuffix(args []any) []any {
	obj := args[0]
	key := args[1]
	if obj == nil || key == nil {
		return []any{false}
	}
	switch obj := obj.(type) {
	case string:
		return []any{strings.HasSuffix(obj, key.(string))}
	default:
		panic(fmt.Sprintf("endWith: not support type %T", obj))
	}
}

func Trim(args []any) []any {
	obj := args[0]
	cutset := args[1]
	if obj == nil || cutset == nil {
		return []any{false}
	}
	switch obj := obj.(type) {
	case string:
		return []any{strings.Trim(obj, cutset.(string))}
	default:
		panic(fmt.Sprintf("trim: not support type %T", obj))
	}
}

func TrimPrefix(args []any) []any {
	obj := args[0]
	cutset := args[1]
	if obj == nil || cutset == nil {
		return []any{false}
	}
	switch obj := obj.(type) {
	case string:
		return []any{strings.TrimPrefix(obj, cutset.(string))}
	default:
		panic(fmt.Sprintf("trimPrefix: not support type %T", obj))
	}
}

func TrimSuffix(args []any) []any {
	obj := args[0]
	cutset := args[1]
	if obj == nil || cutset == nil {
		return []any{false}
	}
	switch obj := obj.(type) {
	case string:
		return []any{strings.TrimSuffix(obj, cutset.(string))}
	default:
		panic(fmt.Sprintf("trimSuffix: not support type %T", obj))
	}
}

func TrimSpace(args []any) []any {
	obj := args[0]
	switch obj := obj.(type) {
	case string:
		return []any{strings.TrimSpace(obj)}
	default:
		panic(fmt.Sprintf("trimSpace: not support type %T", obj))
	}
}

func TrimLeft(args []any) []any {
	obj := args[0]
	cutset := args[1]
	switch obj := obj.(type) {
	case string:
		return []any{strings.TrimLeft(obj, cutset.(string))}
	default:
		panic(fmt.Sprintf("trimLeft: not support type %T", obj))
	}
}

func TrimRight(args []any) []any {
	obj := args[0]
	cutset := args[1]
	switch obj := obj.(type) {
	case string:
		return []any{strings.TrimRight(obj, cutset.(string))}
	default:
		panic(fmt.Sprintf("trimRight: not support type %T", obj))
	}
}

func toLower(args []any) []any {
	obj := args[0]
	if obj == nil {
		return []any{""}
	}
	switch obj := obj.(type) {
	case string:
		return []any{strings.ToLower(obj)}
	default:
		panic(fmt.Sprintf("toLower: not support type %T", obj))
	}
}
func toUpper(args []any) []any {
	obj := args[0]
	if obj == nil {
		return []any{""}
	}
	switch obj := obj.(type) {
	case string:
		return []any{strings.ToUpper(obj)}
	default:
		panic(fmt.Sprintf("toUpper: not support type %T", obj))
	}
}

func Split(args []any) []any {
	obj := args[0]
	sep := args[1]
	if obj == nil || sep == nil {
		return []any{nil}
	}
	switch obj := obj.(type) {
	case string:
		return []any{strings.Split(obj, sep.(string))}
	default:
		panic(fmt.Sprintf("split: not support type %T", obj))
	}
}

func Join(args []any) []any {
	obj := args[0]
	sep := args[1]
	if obj == nil || sep == nil {
		return []any{nil}
	}
	switch obj := obj.(type) {
	case []string:
		return []any{strings.Join(obj, sep.(string))}
	case []any:
		return []any{strings.Join(ToStringSlice(obj), sep.(string))}
	default:
		panic(fmt.Sprintf("join: not support type %T", obj))
	}
}

// ToStringSlice converts []any to []string
func ToStringSlice(s []any) []string {
	if s == nil {
		return nil
	}
	ret := make([]string, len(s))
	for i, v := range s {
		ret[i] = v.(string)
	}
	return ret
}

// time

func Now(_ []any) []any {
	return []any{time.Now()}
}

func ParseTime(args []any) []any {
	layout := args[0]
	value := args[1]
	if layout == nil || value == nil {
		return []any{nil, fmt.Errorf("parsetime err: invalid layout(%v) time(%v) ", layout, value)}
	}
	switch layout := layout.(type) {
	case string:
		t, err := time.Parse(layout, value.(string))
		return []any{t, err}
	default:
		panic(fmt.Sprintf("parseTime: not support type %T", args[0]))
	}
}

func ParseDuration(args []any) []any {
	obj := args[0]
	switch obj := obj.(type) {
	case string:
		d, err := time.ParseDuration(obj)
		return []any{d, err}
	default:
		panic(fmt.Sprintf("parseDuration: not support type %T", obj))
	}
}

func Atoi(args []any) []any {
	obj := args[0]
	switch obj := obj.(type) {
	case string:
		d, err := strconv.Atoi(obj)
		return []any{int64(d), err}
	default:
		panic(fmt.Sprintf("atoi: not support type %T", obj))
	}
}

func Itoa(args []any) []any {
	d := strconv.FormatInt(gen.ToInt64(args[0]), 10)
	return []any{d}
}

func Duration(args []any) []any {
	obj := args[0]
	switch obj := obj.(type) {
	case time.Duration:
		return []any{obj}
	default:
		return []any{time.Duration(gen.ToInt64(obj))}
	}
}

func ToJson(args []any) []any {
	obj := args[0]
	b, err := json.MarshalString(obj)
	return []any{b, err}
}

func FromJson(args []any) []any {
	obj := args[0]
	var v map[string]any
	err := json.UnmarshalString(obj.(string), &v)
	return []any{v, err}
}

func UnmarshalJson(args []any) []any {
	str, obj := args[0], args[1]
	err := json.UnmarshalString(str.(string), &obj)
	return []any{err}
}
