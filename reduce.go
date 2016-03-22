package fstord

import (
	"fmt"
	"reflect"
)

// Reduce for each element in the enumerable. The accumulated and element
// are passed into fun as arguments.
func Reduce(enumerable, fun, initial interface{}) interface{} {
	enums, funValue := preFetchValues(enumerable, fun)
	it := reflect.ValueOf(initial)

	switch enums.Kind() {
	case reflect.Slice:
		return reduceOnSlice(enums, funValue, it)
	case reflect.Map:
		return reduceOnMap(enums, funValue, it)
	default:
		panic(fmt.Sprintf("%s does not support Reduce", enums.Type()))
	}
}

func reduceOnSlice(enums, funValue, initial reflect.Value) interface{} {
	et := enums.Type().Elem()
	if !validReduceFun(funValue, initial.Type(), et, initial.Type()) {
		panic(fmt.Sprintf("fun %s is not a valid type", funValue))
	}

	rt := initial
	for i := 0; i < enums.Len(); i++ {
		rt = funValue.Call([]reflect.Value{rt, enums.Index(i)})[0]
	}
	return rt.Interface()
}

func reduceOnMap(enums, funValue, initial reflect.Value) interface{} {
	kt := enums.Type().Key()
	et := enums.Type().Elem()

	if !validReduceFun(funValue, initial.Type(), kt, et, initial.Type()) {
		panic(fmt.Sprintf("fun %s is not a valid type for map enumeration", funValue))
	}

	rt := initial
	for _, k := range enums.MapKeys() {
		v := enums.MapIndex(k)
		rt = funValue.Call([]reflect.Value{rt, k, v})[0]
	}

	return rt.Interface()
}
