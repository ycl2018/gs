package consts

import (
	"reflect"
	"sync"
)

type RuntimeCache struct {
	fieldIndex  map[reflect.Type]map[string][]int       // type -> fieldName -> index
	methodIndex map[reflect.Type]map[string]MethodIndex // type -> methodName -> index
	fieldLock   sync.RWMutex
	methodLock  sync.RWMutex
}

func NewRuntimeCache() *RuntimeCache {
	return &RuntimeCache{
		fieldIndex:  make(map[reflect.Type]map[string][]int),
		methodIndex: make(map[reflect.Type]map[string]MethodIndex),
	}
}

func (r *RuntimeCache) FetchFieldIndex(t reflect.Type, fieldName string) ([]int, bool) {
	r.fieldLock.RLock()
	defer r.fieldLock.RUnlock()
	if fields, ok := r.fieldIndex[t]; ok {
		if indexes, ok := fields[fieldName]; ok {
			return indexes, ok
		}
	}
	return nil, false
}

func (r *RuntimeCache) SetFieldIndex(t reflect.Type, fieldName string, indexes []int) {
	r.fieldLock.Lock()
	defer r.fieldLock.Unlock()
	if fields, ok := r.fieldIndex[t]; ok {
		fields[fieldName] = indexes
	} else {
		r.fieldIndex[t] = map[string][]int{
			fieldName: indexes,
		}
	}
}

type Convert int

const (
	No Convert = iota
	Elem
	PtrTo
)

type MethodIndex struct {
	Index    []int
	Convert  Convert
	Ptr      bool
	IsMethod bool
}

func (r *RuntimeCache) FetchMethodIndex(t reflect.Type, methodName string) (MethodIndex, bool) {
	r.methodLock.RLock()
	defer r.methodLock.RUnlock()
	if methods, ok := r.methodIndex[t]; ok {
		if index, ok := methods[methodName]; ok {
			return index, ok
		}
	}
	return MethodIndex{}, false
}

func (r *RuntimeCache) SetMethodIndex(t reflect.Type, methodName string, index MethodIndex) {
	r.methodLock.Lock()
	defer r.methodLock.Unlock()
	if methods, ok := r.methodIndex[t]; ok {
		methods[methodName] = index
	} else {
		r.methodIndex[t] = map[string]MethodIndex{
			methodName: index,
		}
	}
}
