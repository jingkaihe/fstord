package fstord

import (
	"fmt"
	"reflect"
)

// Map Returns a slice/map where each element is the result of invoking
// fun on each corresponding element of slice/map
func Map(enumerable, fun interface{}) interface{} {
	enums, funValue := preFetchValues(enumerable, fun)

	switch enums.Kind() {
	case reflect.Slice:
		return mapOnSlice(enums, funValue)
	case reflect.Map:
		return mapOnMap(enums, funValue)
	default:
		panic(fmt.Sprintf("%s does not support Map", enums.Type()))
	}
}

func mapOnSlice(enums, funValue reflect.Value) interface{} {
	et := enums.Type().Elem()
	if !validFun(funValue, et) {
		panic(fmt.Sprintf("%s is not a valid type for fun %s", et, funValue))
	}

	rt := reflect.MakeSlice(reflect.SliceOf(funValue.Type().Out(0)), enums.Len(), enums.Len())
	for i := 0; i < enums.Len(); i++ {
		cs := funValue.Call([]reflect.Value{enums.Index(i)})[0]
		rt.Index(i).Set(cs)
	}
	return rt.Interface()
}

func mapOnMap(enums, funValue reflect.Value) interface{} {
	kt := enums.Type().Key()
	et := enums.Type().Elem()
	mapType := reflect.MapOf(kt, funValue.Type().Out(0))
	if !validFun(funValue, kt, et) || mapType.Kind() != reflect.Map {
		panic(fmt.Sprintf("func (%s, %s) -> %s type is invalid", kt, et, mapType))
	}
	rt := reflect.MakeMap(mapType)

	for _, k := range enums.MapKeys() {
		v := funValue.Call([]reflect.Value{k, enums.MapIndex(k)})[0]
		rt.SetMapIndex(k, v)
	}

	return rt.Interface()
}
