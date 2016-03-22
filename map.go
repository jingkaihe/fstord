package fstord

import (
	"fmt"
	"reflect"
)

// Map Returns a slice/map where each element is the result of invoking
// fun on each corresponding element of slice/map
func Map(enumerable, fun interface{}) interface{} {
	enums, fv := preFetchValues(enumerable, fun)

	switch enums.Kind() {
	case reflect.Slice:
		return mapOnSlice(enums, fv)
	case reflect.Map:
		return mapOnMap(enums, fv)
	default:
		panic(fmt.Sprintf("%s does not support Map", enums.Type()))
	}
}

func mapOnSlice(enums, fv reflect.Value) interface{} {
	et := enums.Type().Elem()
	if !validFun(fv, et) {
		panic(fmt.Sprintf("%s is not a valid type for fun %s", et, fv))
	}

	rt := reflect.MakeSlice(reflect.SliceOf(fv.Type().Out(0)), enums.Len(), enums.Len())
	for i := 0; i < enums.Len(); i++ {
		cs := fv.Call([]reflect.Value{enums.Index(i)})[0]
		rt.Index(i).Set(cs)
	}
	return rt.Interface()
}

func mapOnMap(enums, fv reflect.Value) interface{} {
	kt := enums.Type().Key()
	et := enums.Type().Elem()
	mapType := reflect.MapOf(kt, fv.Type().Out(0))
	if !validFun(fv, kt, et) || mapType.Kind() != reflect.Map {
		panic(fmt.Sprintf("func (%s, %s) -> %s type is invalid", kt, et, mapType))
	}
	rt := reflect.MakeMap(mapType)

	for _, k := range enums.MapKeys() {
		v := fv.Call([]reflect.Value{k, enums.MapIndex(k)})[0]
		rt.SetMapIndex(k, v)
	}

	return rt.Interface()
}
