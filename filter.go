package fstord

import (
	"fmt"
	"reflect"
)

// Filter the enumerable - returns the elements for which fun returns true
func Filter(enumerable, fun interface{}) interface{} {
	mvs := reflect.ValueOf(enumerable)
	fv := reflect.ValueOf(fun)

	switch mvs.Kind() {
	case reflect.Slice:
		et := mvs.Type().Elem()
		if !validFilterFun(fv, et) {
			panic(fmt.Sprintf("%s is not a valid type for fun %s", et, fv))
		}

		rt := reflect.MakeSlice(reflect.SliceOf(et), 0, mvs.Len())
		for i := 0; i < mvs.Len(); i++ {
			if fv.Call([]reflect.Value{mvs.Index(i)})[0].Bool() == true {
				rt = reflect.Append(rt, mvs.Index(i))
			}
		}
		return rt.Interface()
	case reflect.Map:
		kt := mvs.Type().Key()
		et := mvs.Type().Elem()
		mapType := mvs.Type()
		if !validFilterFun(fv, kt, et) {
			panic(fmt.Sprintf("func (%s, %s) -> %s is invalid for filter", kt, et, fv.Type()))
		}
		rt := reflect.MakeMap(mapType)

		for _, k := range mvs.MapKeys() {
			v := mvs.MapIndex(k)
			if fv.Call([]reflect.Value{k, v})[0].Bool() == true {
				rt.SetMapIndex(k, v)
			}
		}
		return rt.Interface()
	default:
		panic(fmt.Sprintf("%s does not support Map", mvs.Type()))
	}
}

func validFilterFun(fun reflect.Value, types ...reflect.Type) bool {
	return validFun(fun, types...) && fun.Type().Out(0).Kind() == reflect.Bool
}
