package fstord

import (
	"fmt"
	"reflect"
)

// Map Returns a slice/map where each element is the result of invoking
// fun on each corresponding element of slice/map
func Map(enumerable, fun interface{}) interface{} {
	mvs := reflect.ValueOf(enumerable)
	fv := reflect.ValueOf(fun)

	switch mvs.Kind() {
	case reflect.Slice:
		et := mvs.Type().Elem()
		if !validFun(fv, et) {
			panic(fmt.Sprintf("%s is not a valid type for fun %s", et, fv))
		}

		rt := reflect.MakeSlice(reflect.SliceOf(fv.Type().Out(0)), mvs.Len(), mvs.Len())
		for i := 0; i < mvs.Len(); i++ {
			cs := fv.Call([]reflect.Value{mvs.Index(i)})[0]
			rt.Index(i).Set(cs)
		}
		return rt.Interface()
	case reflect.Map:
		kt := mvs.Type().Key()
		et := mvs.Type().Elem()
		mapType := reflect.MapOf(kt, fv.Type().Out(0))
		if !validFun(fv, kt, et) || mapType.Kind() != reflect.Map {
			panic(fmt.Sprintf("func (%s, %s) -> %s type is invalid", kt, et, mapType))
		}
		rt := reflect.MakeMap(mapType)

		for _, k := range mvs.MapKeys() {
			v := fv.Call([]reflect.Value{k, mvs.MapIndex(k)})[0]
			rt.SetMapIndex(k, v)
		}

		return rt.Interface()
	default:
		panic(fmt.Sprintf("%s does not support Map", mvs.Type()))
	}
}

func validFun(fun reflect.Value, types ...reflect.Type) bool {
	// invalid if fun is not a function
	if fun.Kind() != reflect.Func {
		return false
	}

	// invalid if the params count is not the same
	if fun.Type().NumIn() != len(types) {
		return false
	}

	// invalid if the out count is not 1 (for now)
	if fun.Type().NumOut() != 1 {
		return false
	}

	// invalid if any of the params type doesn't match
	for i := 0; i < len(types); i++ {
		if fun.Type().In(i) != types[i] {
			return false
		}
	}

	return true
}
