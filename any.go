package fstord

import (
	"fmt"
	"reflect"
)

// Any Returns whether any elements of a slice/map return true when invoke fun
func Any(enumerable, fun interface{}) bool {
	mvs := reflect.ValueOf(enumerable)
	fv := reflect.ValueOf(fun)

	switch mvs.Kind() {
	case reflect.Slice:
		et := mvs.Type().Elem()
		if !validAnyFun(fv, et) {
			panic(fmt.Sprintf("%s is not a valid type for fun %s", et, fv))
		}

		for i := 0; i < mvs.Len(); i++ {
			res := fv.Call([]reflect.Value{mvs.Index(i)})[0]
			if res.Bool() == true {
				return true
			}
		}
		return false
	case reflect.Map:
		kt := mvs.Type().Key()
		et := mvs.Type().Elem()
		if !validAnyFun(fv, kt, et) {
			panic(fmt.Sprintf("func %s is invalid", fun))
		}

		for _, k := range mvs.MapKeys() {
			v := fv.Call([]reflect.Value{k, mvs.MapIndex(k)})[0]
			if v.Bool() == true {
				return true
			}
		}

		return false
	default:
		panic(fmt.Sprintf("%s does not support Map", mvs.Type()))
	}
}

func validAnyFun(fun reflect.Value, types ...reflect.Type) bool {
	return validFun(fun, types...) && fun.Type().Out(0).Kind() == reflect.Bool
}
