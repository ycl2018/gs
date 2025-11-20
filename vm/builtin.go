package vm

import (
	"fmt"
	"reflect"
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
	default:
		panic(fmt.Sprintf("can't length %T", x))
	}
}

func appendSlice(x any, values []any) any {
	switch x := x.(type) {
	case []any:
		return append(x, values...)
	case []string:
		return append(x, CopyToSliceByKind(values, reflect.String).([]string)...)
	case []int:
		return append(x, CopyToSliceByKind(values, reflect.Int).([]int)...)
	case []int64:
		return append(x, CopyToSliceByKind(values, reflect.Int64).([]int64)...)
	case []int8:
		return append(x, CopyToSliceByKind(values, reflect.Int8).([]int8)...)
	case []int16:
		return append(x, CopyToSliceByKind(values, reflect.Int16).([]int16)...)
	case []int32:
		return append(x, CopyToSliceByKind(values, reflect.Int32).([]int32)...)
	case []uint:
		return append(x, CopyToSliceByKind(values, reflect.Uint).([]uint)...)
	case []uint8:
		return append(x, CopyToSliceByKind(values, reflect.Uint8).([]uint8)...)
	case []uint16:
		return append(x, CopyToSliceByKind(values, reflect.Uint16).([]uint16)...)
	case []uint32:
		return append(x, CopyToSliceByKind(values, reflect.Uint32).([]uint32)...)
	case []uint64:
		return append(x, CopyToSliceByKind(values, reflect.Uint64).([]uint64)...)
	case []float32:
		return append(x, CopyToSliceByKind(values, reflect.Float32).([]float32)...)
	case []float64:
		return append(x, CopyToSliceByKind(values, reflect.Float64).([]float64)...)
	case []uintptr:
		return append(x, CopyToSliceByKind(values, reflect.Uintptr).([]uintptr)...)
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
			return append(x, CopyToSliceByKind(expandSlice, reflect.String).([]string)...)
		case []int:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Int).([]int)...)
		case []int8:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Int8).([]int8)...)
		case []int16:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Int16).([]int16)...)
		case []int32:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Int32).([]int32)...)
		case []int64:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Int64).([]int64)...)
		case []uint8:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Uint8).([]uint8)...)
		case []uint:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Uint).([]uint)...)
		case []uint16:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Uint16).([]uint16)...)
		case []uint32:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Uint32).([]uint32)...)
		case []uint64:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Uint64).([]uint64)...)
		case []float32:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Float32).([]float32)...)
		case []float64:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Float64).([]float64)...)
		case []uintptr:
			return append(x, CopyToSliceByKind(expandSlice, reflect.Uintptr).([]uintptr)...)
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
		delete(m, ToInt(k))
		return
	case map[int]any:
		delete(m, ToInt(k))
		return
	case map[int]bool:
		delete(m, ToInt(k))
		return
	case map[int]int:
		delete(m, ToInt(k))
		return
	case map[int64]string:
		delete(m, ToInt64(k))
		return
	case map[int64]any:
		delete(m, ToInt64(k))
		return
	case map[int64]int64:
		delete(m, ToInt64(k))
		return
	case map[int64]bool:
		delete(m, ToInt64(k))
		return
	default:
	}
	rm := reflect.ValueOf(m)
	if rm.Kind() != reflect.Map {
		panic(fmt.Sprintf("delete unsupport type:%T", m))
	}
	switch k.(type) {
	case int: // int may be vm value, map
		switch kt := rm.Type().Key().Kind(); kt {
		case reflect.Int:
			k = ToInt(k)
		case reflect.Int8:
			k = ToInt8(k)
		case reflect.Int16:
			k = ToInt16(k)
		case reflect.Int32:
			k = ToInt32(k)
		case reflect.Int64:
			k = ToInt64(k)
		case reflect.Uint:
			k = ToUint(k)
		case reflect.Uint8:
			k = ToUint8(k)
		case reflect.Uint16:
			k = ToUint16(k)
		case reflect.Uint32:
			k = ToUint32(k)
		case reflect.Uint64:
			k = ToUint64(k)
		case reflect.Uintptr:
			k = ToUintptr(k)
		case reflect.Float64:
			k = ToFloat64(k)
		case reflect.Float32:
			k = ToFloat32(k)
		}
	}

	rm.SetMapIndex(reflect.ValueOf(k), reflect.Value{})
}

func copySlice(src any) any {
	//TODO
	panic("to be copied")
}

func toString(v any) string {
	if v, ok := v.(fmt.Stringer); ok {
		return v.String()
	}
	return fmt.Sprintf("%v", v)
}

func convert(v any, kind reflect.Kind) any {
	switch kind {
	case reflect.Int8:
		return ToInt8(v)
	case reflect.Int16:
		return ToInt16(v)
	case reflect.Int32:
		return ToInt32(v)
	case reflect.Int64:
		return ToInt64(v)
	case reflect.Uint:
		return ToUint(v)
	case reflect.Uint8:
		return ToUint8(v)
	case reflect.Uint16:
		return ToUint16(v)
	case reflect.Uint32:
		return ToUint32(v)
	case reflect.Uint64:
		return ToUint64(v)
	case reflect.Float32:
		return ToFloat32(v)
	case reflect.Float64:
		return ToFloat64(v)
	case reflect.String:
		return toString(v)
	case reflect.Bool:
		return v.(bool)
	}
	panic(fmt.Sprintf("can't convert %T to kind:%s", v, kind))
}
