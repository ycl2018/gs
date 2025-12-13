package vm

import (
	"reflect"
	"sync"

	"github.com/modern-go/reflect2"
)

func fieldByIndex(obj any, fields []reflect2.StructField) any {
	if obj == nil {
		panic("obj is nil")
	}
	obj, _ = indirect(obj)
	for _, field := range fields {
		obj = field.Get(obj)
	}
	return obj
}

func indirect(obj any) (any, reflect2.Type) {
	tye := reflect2.TypeOf(obj)
	for tye.Kind() == reflect.Ptr {
		tye = tye.(*reflect2.UnsafePtrType).Elem()
		obj = tye.Indirect(obj)
	}
	return obj, tye
}

var fieldByNameCache sync.Map

func fieldByName(obj any, name string) any {
	obj, itype := indirect(obj)
	switch itype.Kind() {
	case reflect.Struct:
		// cache
		rType := itype.(*reflect2.UnsafeStructType).RType()
		value, ok := fieldByNameCache.Load(rType)
		if ok {
			load, ok2 := value.(*sync.Map).Load(name)
			if ok2 {
				return load.(*reflect2.UnsafeStructField).Get(obj)
			}
		}
		structField := itype.(*reflect2.UnsafeStructType).FieldByName(name)
		get := structField.Get(obj)
		if get != nil {
			// set cache
		}
		return get
	case reflect.Interface:
		byName, ok := itype.(*reflect2.UnsafeEFaceType).MethodByName(name)
		if ok {
			byName
		}


	}
}
