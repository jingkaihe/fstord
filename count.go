package fstord

import (
	"fmt"
	"reflect"
)

// Count returns the number of elements of a slice/map that return true when invoke fun
func Count(enumerable, fun interface{}) int {
	mvs := reflect.ValueOf(enumerable)
	fv := reflect.ValueOf(fun)

	switch mvs.Kind() {
	case reflect.Slice:
		et := mvs.Type().Elem()
		if !validBoolFun(fv, et) {
			panic(fmt.Sprintf("%s is not a valid type for fun %s", et, fv))
		}

		cnt := 0
		for i := 0; i < mvs.Len(); i++ {
			res := fv.Call([]reflect.Value{mvs.Index(i)})[0]
			if res.Bool() == true {
				cnt++
			}
		}
		return cnt
	case reflect.Map:
		kt := mvs.Type().Key()
		et := mvs.Type().Elem()
		if !validBoolFun(fv, kt, et) {
			panic(fmt.Sprintf("func %s is invalid", fun))
		}

		cnt := 0
		for _, k := range mvs.MapKeys() {
			v := fv.Call([]reflect.Value{k, mvs.MapIndex(k)})[0]
			if v.Bool() == true {
				cnt++
			}
		}

		return cnt
	default:
		panic(fmt.Sprintf("%s does not support Count", mvs.Type()))
	}
}
