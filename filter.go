package fstord

import (
	"fmt"
	"reflect"
)

// Filter the enumerable - returns the elements for which fun returns true
func Filter(enumerable, fun interface{}) interface{} {
	enums, funValue := preFetchValues(enumerable, fun)

	switch enums.Kind() {
	case reflect.Slice:
		return filterOnSlice(enums, funValue)
	case reflect.Map:
		return filterOnMap(enums, funValue)
	default:
		panic(fmt.Sprintf("%s does not support Filter", enums.Type()))
	}
}

func filterOnSlice(enums, funValue reflect.Value) interface{} {
	et := enums.Type().Elem()
	if !validBoolFun(funValue, et) {
		panic(fmt.Sprintf("%s is not a valid type for fun %s", et, funValue))
	}

	rt := reflect.MakeSlice(reflect.SliceOf(et), 0, enums.Len())
	for i := 0; i < enums.Len(); i++ {
		if funValue.Call([]reflect.Value{enums.Index(i)})[0].Bool() == true {
			rt = reflect.Append(rt, enums.Index(i))
		}
	}
	return rt.Interface()
}

func filterOnMap(enums, funValue reflect.Value) interface{} {
	kt := enums.Type().Key()
	et := enums.Type().Elem()
	mapType := enums.Type()
	if !validBoolFun(funValue, kt, et) {
		panic(fmt.Sprintf("func (%s, %s) -> %s is invalid for filter", kt, et, funValue.Type()))
	}
	rt := reflect.MakeMap(mapType)

	for _, k := range enums.MapKeys() {
		v := enums.MapIndex(k)
		if funValue.Call([]reflect.Value{k, v})[0].Bool() == true {
			rt.SetMapIndex(k, v)
		}
	}
	return rt.Interface()
}
