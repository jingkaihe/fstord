package fstord

import (
	"fmt"
	"reflect"
)

// Filter the enumerable - returns the elements for which fun returns true
func Filter(enumerable, fun interface{}) interface{} {
	enums := reflect.ValueOf(enumerable)
	fv := reflect.ValueOf(fun)

	switch enums.Kind() {
	case reflect.Slice:
		et := enums.Type().Elem()
		if !validBoolFun(fv, et) {
			panic(fmt.Sprintf("%s is not a valid type for fun %s", et, fv))
		}

		rt := reflect.MakeSlice(reflect.SliceOf(et), 0, enums.Len())
		for i := 0; i < enums.Len(); i++ {
			if fv.Call([]reflect.Value{enums.Index(i)})[0].Bool() == true {
				rt = reflect.Append(rt, enums.Index(i))
			}
		}
		return rt.Interface()
	case reflect.Map:
		kt := enums.Type().Key()
		et := enums.Type().Elem()
		mapType := enums.Type()
		if !validBoolFun(fv, kt, et) {
			panic(fmt.Sprintf("func (%s, %s) -> %s is invalid for filter", kt, et, fv.Type()))
		}
		rt := reflect.MakeMap(mapType)

		for _, k := range enums.MapKeys() {
			v := enums.MapIndex(k)
			if fv.Call([]reflect.Value{k, v})[0].Bool() == true {
				rt.SetMapIndex(k, v)
			}
		}
		return rt.Interface()
	default:
		panic(fmt.Sprintf("%s does not support Filter", enums.Type()))
	}
}

func validBoolFun(fun reflect.Value, types ...reflect.Type) bool {
	return validFun(fun, types...) && fun.Type().Out(0).Kind() == reflect.Bool
}
