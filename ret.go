package gs

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ycl2018/gs/consts"
)

type Result struct {
	value any
}

func (r *Result) RawValue() any {
	return r.value
}

func (r *Result) IsNil() bool {
	return r.value == nil
}

func (r *Result) MustString() string {
	s, err := r.String()
	if err != nil {
		panic(err)
	}
	return s
}

func (r *Result) Error() error {
	if r.value == nil {
		return nil
	}
	switch r.value.(type) {
	case error:
		return r.value.(error)
	default:
		panic(fmt.Errorf("value type:%T is not error", r.value))
	}
}

func (r *Result) String() (string, error) {
	if r.value == nil {
		return "", errors.New("value is nil")
	}
	switch r.value.(type) {
	case string:
		return r.value.(string), nil
	default:
		rv := reflect.ValueOf(r.value)
		switch rv.Kind() {
		case reflect.String:
			return rv.String(), nil
		default:
			return "", fmt.Errorf("value type:%T is not string", r.value)
		}
	}
}

func (r *Result) MustInt() int {
	i, err := r.Int()
	if err != nil {
		panic(err)
	}
	return i
}

func (r *Result) Int() (int, error) {
	if r.value == nil {
		return 0, errors.New("value is nil")
	}
	switch v := r.value.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case uintptr:
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	default:
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int(rv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return int(rv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return int(rv.Float()), nil
		default:
			return 0, fmt.Errorf("value type:%T is not or cannot convert to int", r.value)
		}
	}
}

func (r *Result) MustBool() bool {
	b, err := r.Bool()
	if err != nil {
		panic(err)
	}
	return b
}

func (r *Result) Bool() (bool, error) {
	if r.value == nil {
		return false, errors.New("value is nil")
	}
	switch v := r.value.(type) {
	case bool:
		return v, nil
	default:
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Bool:
			return rv.Bool(), nil
		default:
			return false, fmt.Errorf("value type:%T is not bool", r.value)
		}
	}
}

func (r *Result) MustFloat() float64 {
	f, err := r.Float()
	if err != nil {
		panic(err)
	}
	return f
}

func (r *Result) Float() (float64, error) {
	if r.value == nil {
		return 0, errors.New("value is nil")
	}
	switch v := r.value.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case uintptr:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return float64(rv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return float64(rv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return rv.Float(), nil
		default:
			return 0, fmt.Errorf("value type:%T is not or cannot convert to float64", r.value)
		}
	}
}

func (r *Result) Get(index int) *Result {
	t, ok := r.value.(consts.Tuple)
	if ok {
		if index < 0 || index >= t.Num {
			panic(fmt.Sprintf("index %d out of range", index))
		}
		return &Result{value: t.Values[index]}
	} else if index == 0 {
		return r
	}
	return &Result{value: nil}
}
