package consts

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/ycl2018/gs/gen"
)

var Predefine = map[string]*DefineFunc{
	// collections
	"in": {
		Name:   "in",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     In,
	},
	"index": {
		Name:   "index",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     Index,
	},
	// strings
	"hasPrefix": {
		Name:   "hasPrefix",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     HasPrefix,
	},
	"hasSuffix": {
		Name:   "hasSuffix",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     HasSuffix,
	},
	"trim": {
		Name:   "trim",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     Trim,
	},
	"trimPrefix": {
		Name:   "trimPrefix",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     TrimPrefix,
	},
	"trimSuffix": {
		Name:   "trimSuffix",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     TrimSuffix,
	},
	"trimSpace": {
		Name:   "trimSpace",
		NumIn:  1,
		NumOut: 1,
		Fast:   true,
		Fn:     TrimSpace,
	},
	"trimLeft": {
		Name:   "trimLeft",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     TrimLeft,
	},
	"trimRight": {
		Name:   "trimRight",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     TrimRight,
	},
	"toLower": {
		Name:   "toLower",
		NumIn:  1,
		NumOut: 1,
		Fast:   true,
		Fn:     toLower,
	},
	"toUpper": {
		Name:   "toUpper",
		NumIn:  1,
		NumOut: 1,
		Fast:   true,
		Fn:     toUpper,
	},
	"split": {
		Name:   "split",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     Split,
	},
	"join": {
		Name:   "join",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     Join,
	},
	// system
	"now": {
		Name:   "now",
		NumIn:  0,
		NumOut: 1,
		Fast:   true,
		Fn:     Now,
	},
	"parseTime": {
		Name:   "parseTime",
		NumIn:  2,
		NumOut: 2,
		Fast:   true,
		Fn:     ParseTime,
	},
	"parseDuration": {
		Name:   "parseDuration",
		NumIn:  1,
		NumOut: 2,
		Fast:   true,
		Fn:     ParseDuration,
	},
	"atoi": {
		Name:   "atoi",
		NumIn:  1,
		NumOut: 2,
		Fast:   true,
		Fn:     Atoi,
	},
	"itoa": {
		Name:   "itoa",
		NumIn:  1,
		NumOut: 1,
		Fast:   true,
		Fn:     Itoa,
	},
	"duration": {
		Name:   "duration",
		NumIn:  1,
		NumOut: 1,
		Fast:   true,
		Fn:     Duration,
	},
	// json
	"toJson": {
		Name:   "toJson",
		NumIn:  1,
		NumOut: 2,
		Fast:   true,
		Fn:     ToJson,
	},
	"fromJson": {
		Name:   "fromJson",
		NumIn:  1,
		NumOut: 2,
		Fast:   true,
		Fn:     FromJson,
	},
	"unmarshalJson": {
		Name:   "unmarshalJson",
		NumIn:  2,
		NumOut: 1,
		Fast:   true,
		Fn:     UnmarshalJson,
	},
}

func In(args []any) []any {
	obj := args[0]
	key := args[1]
	switch obj := obj.(type) {
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
			if ok, _ := gen.Eq(rv.Index(i).Interface(), key); ok {
				return []any{true}
			}
		}
		return []any{false}
	case reflect.Map:
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
	switch obj := obj.(type) {
	case []int:
		for i, v := range obj {
			if v == key.(int) {
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
		for i, v := range obj {
			if v == key.(int64) {
				return []any{i}
			}
		}
		return []any{-1}
	case []any:
		for i, v := range obj {
			if v == key {
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
			if rv.Index(i).Interface() == key {
				return []any{i}
			}
		}
		return []any{-1}
	default:
		panic(fmt.Sprintf("index: not support type %T", obj))
	}
}

func HasPrefix(args []any) []any {
	obj := args[0]
	key := args[1]
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
	switch obj := obj.(type) {
	case string:
		return []any{strings.ToLower(obj)}
	default:
		panic(fmt.Sprintf("toLower: not support type %T", obj))
	}
}
func toUpper(args []any) []any {
	obj := args[0]
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
	d := strconv.FormatInt(ToInt64(args[0]), 10)
	return []any{d}
}

func Duration(args []any) []any {
	obj := args[0]
	switch obj := obj.(type) {
	case time.Duration:
		return []any{obj}
	default:
		return []any{time.Duration(ToInt64(obj))}
	}
}

func ToInt64(value any) int64 {
	switch v := value.(type) {
	case int64:
		return v
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case uintptr:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	default:
		panic(fmt.Sprintf("ToInt64: unsupport %T", value))
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
