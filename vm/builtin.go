package vm

import (
	"fmt"
	"reflect"

	"github.com/ycl2018/gs/gen"
)

func length(x any) int {
	switch x := x.(type) {
	case []any:
		return len(x)
	case map[any]any:
		return len(x)
	case string:
		return len(x)
	case []string:
		return len(x)
	case []int:
		return len(x)
	case []int8:
		return len(x)
	case []int64:
		return len(x)
	case map[string]any:
		return len(x)
	case map[string]int:
		return len(x)
	case map[string]string:
		return len(x)
	case map[string]int64:
		return len(x)
	case []int32:
		return len(x)
	case []float32:
		return len(x)
	case []float64:
		return len(x)
	}
	xv := reflect.ValueOf(x)
	switch xv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String, reflect.Chan:
		return xv.Len()
	case reflect.Pointer:
		if xv.Elem().Kind() == reflect.Array {
			return xv.Elem().Len()
		}
	default:
	}
	panic(fmt.Sprintf("can't length %T", x))
}

func appendSlice(x any, values []any) any {
	switch x := x.(type) {
	case []any:
		return append(x, values...)
	case []string:
		return append(x, gen.CopyToSliceByKind(values, reflect.String).([]string)...)
	case []int:
		return append(x, gen.CopyToSliceByKind(values, reflect.Int).([]int)...)
	case []int64:
		return append(x, gen.CopyToSliceByKind(values, reflect.Int64).([]int64)...)
	case []int8:
		return append(x, gen.CopyToSliceByKind(values, reflect.Int8).([]int8)...)
	case []int16:
		return append(x, gen.CopyToSliceByKind(values, reflect.Int16).([]int16)...)
	case []int32:
		return append(x, gen.CopyToSliceByKind(values, reflect.Int32).([]int32)...)
	case []uint:
		return append(x, gen.CopyToSliceByKind(values, reflect.Uint).([]uint)...)
	case []uint8:
		return append(x, gen.CopyToSliceByKind(values, reflect.Uint8).([]uint8)...)
	case []uint16:
		return append(x, gen.CopyToSliceByKind(values, reflect.Uint16).([]uint16)...)
	case []uint32:
		return append(x, gen.CopyToSliceByKind(values, reflect.Uint32).([]uint32)...)
	case []uint64:
		return append(x, gen.CopyToSliceByKind(values, reflect.Uint64).([]uint64)...)
	case []float32:
		return append(x, gen.CopyToSliceByKind(values, reflect.Float32).([]float32)...)
	case []float64:
		return append(x, gen.CopyToSliceByKind(values, reflect.Float64).([]float64)...)
	case []uintptr:
		return append(x, gen.CopyToSliceByKind(values, reflect.Uintptr).([]uintptr)...)
	default:
		xv := reflect.ValueOf(x)
		return reflect.AppendSlice(xv, reflect.ValueOf(values)).Interface()
	}
}

func appendSliceExpand(x any, expandSlice any) any {
	switch expandSlice := expandSlice.(type) {
	case []any: // short path for: from vm slice to env
		switch x := x.(type) {
		case []any:
			return append(x, expandSlice...)
		case []string:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.String).([]string)...)
		case []int:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Int).([]int)...)
		case []int8:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Int8).([]int8)...)
		case []int16:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Int16).([]int16)...)
		case []int32:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Int32).([]int32)...)
		case []int64:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Int64).([]int64)...)
		case []uint8:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Uint8).([]uint8)...)
		case []uint:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Uint).([]uint)...)
		case []uint16:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Uint16).([]uint16)...)
		case []uint32:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Uint32).([]uint32)...)
		case []uint64:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Uint64).([]uint64)...)
		case []float32:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Float32).([]float32)...)
		case []float64:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Float64).([]float64)...)
		case []uintptr:
			return append(x, gen.CopyToSliceByKind(expandSlice, reflect.Uintptr).([]uintptr)...)
		}
	}
	// expandSlice is not from env
	switch x := x.(type) {
	case []any:
		// short path for: env to vm slice
		switch expandSlice := expandSlice.(type) {
		case []any:
			return append(x, expandSlice...)
		case []string:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []int:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []int8:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []int16:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []int32:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []int64:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []uint16:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []uint32:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []uint64:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []float32:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []float64:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		case []uintptr:
			for i := 0; i < len(expandSlice); i++ {
				x = append(x, expandSlice[i])
			}
			return x
		}
	case []string:
		return append(x, expandSlice.([]string)...)
	case []int:
		return append(x, expandSlice.([]int)...)
	case []int64:
		return append(x, expandSlice.([]int64)...)
	case []int32:
		return append(x, expandSlice.([]int32)...)
	case []uint8:
		return append(x, expandSlice.([]uint8)...)
	case []uint:
		return append(x, expandSlice.([]uint)...)
	case []uint16:
		return append(x, expandSlice.([]uint16)...)
	case []uint32:
		return append(x, expandSlice.([]uint32)...)
	case []uint64:
		return append(x, expandSlice.([]uint64)...)
	case []float32:
		return append(x, expandSlice.([]float32)...)
	case []float64:
		return append(x, expandSlice.([]float64)...)
	}
	xv := reflect.ValueOf(x)
	return reflect.AppendSlice(xv, reflect.ValueOf(expandSlice)).Interface()
}

func deleteMap(m, k any) {
	switch m := m.(type) {
	case map[any]any:
		delete(m, k)
		return
	case map[string]any:
		delete(m, k.(string))
		return
	case map[string]int:
		delete(m, k.(string))
		return
	case map[string]int64:
		delete(m, k.(string))
		return
	case map[string]bool:
		delete(m, k.(string))
		return
	case map[int]string:
		delete(m, k.(int))
		return
	case map[int]any:
		delete(m, k.(int))
		return
	case map[int]bool:
		delete(m, k.(int))
		return
	case map[int]int:
		delete(m, k.(int))
		return
	case map[int64]string:
		delete(m, k.(int64))
		return
	case map[int64]any:
		delete(m, k.(int64))
		return
	case map[int64]int64:
		delete(m, k.(int64))
		return
	case map[int64]bool:
		delete(m, k.(int64))
		return
	default:
	}
	rm := reflect.ValueOf(m)
	if rm.Kind() != reflect.Map {
		panic(fmt.Sprintf("delete unsupport type:%T", m))
	}
	rm.SetMapIndex(reflect.ValueOf(k), reflect.Value{})
}

var mapKindToType = map[reflect.Kind]reflect.Type{
	reflect.Int:     reflect.TypeOf(int(0)),
	reflect.Int8:    reflect.TypeOf(int8(0)),
	reflect.Int16:   reflect.TypeOf(int16(0)),
	reflect.Int32:   reflect.TypeOf(int32(0)),
	reflect.Int64:   reflect.TypeOf(int64(0)),
	reflect.Uint:    reflect.TypeOf(uint(0)),
	reflect.Uint8:   reflect.TypeOf(uint8(0)),
	reflect.Uint16:  reflect.TypeOf(uint16(0)),
	reflect.Uint32:  reflect.TypeOf(uint32(0)),
	reflect.Uint64:  reflect.TypeOf(uint64(0)),
	reflect.Float32: reflect.TypeOf(float32(0)),
	reflect.Float64: reflect.TypeOf(float64(0)),
	reflect.String:  reflect.TypeOf(""),
	reflect.Bool:    reflect.TypeOf(false),
}

func convert(v any, kind reflect.Kind) any {
	switch kind {
	case reflect.Int:
		if val, ok := gen.TryToInt(v); ok {
			return val
		}
	case reflect.Int8:
		if val, ok := gen.TryToInt8(v); ok {
			return val
		}
	case reflect.Int16:
		if val, ok := gen.TryToInt16(v); ok {
			return val
		}
	case reflect.Int32:
		if val, ok := gen.TryToInt32(v); ok {
			return val
		}
	case reflect.Int64:
		if val, ok := gen.TryToInt64(v); ok {
			return val
		}
	case reflect.Uint:
		if val, ok := gen.TryToUint(v); ok {
			return val
		}
	case reflect.Uint8:
		if val, ok := gen.TryToUint8(v); ok {
			return val
		}
	case reflect.Uint16:
		if val, ok := gen.TryToUint16(v); ok {
			return val
		}
	case reflect.Uint32:
		if val, ok := gen.TryToUint32(v); ok {
			return val
		}
	case reflect.Uint64:
		if val, ok := gen.TryToUint64(v); ok {
			return val
		}
	case reflect.Float32:
		if val, ok := gen.TryToFloat32(v); ok {
			return val
		}
	case reflect.Float64:
		if val, ok := gen.TryToFloat64(v); ok {
			return val
		}
	case reflect.String:
		switch v.(type) {
		case string:
			return v.(string)
		default:
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.String {
				return rv.String()
			}
			panic(fmt.Sprintf("can't convert %T to string", v))
		}
	case reflect.Bool:
		switch v.(type) {
		case bool:
			return v.(bool)
		}
	default:
		panic(fmt.Sprintf("convert unsupport type:%T", v))
	}
	rv := reflect.ValueOf(v)
	if rv.CanConvert(mapKindToType[kind]) {
		return rv.Convert(mapKindToType[kind]).Interface()
	}
	panic(fmt.Sprintf("can't convert %T to %s", v, kind))
}
