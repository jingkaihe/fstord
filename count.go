package fstord

import (
	"fmt"
	"reflect"
)

// Count returns the number of elements of a slice/map that return true when invoke fun
func Count(enumerable, fun interface{}) int {
	enums := reflect.ValueOf(enumerable)
	fv := reflect.ValueOf(fun)

	switch enums.Kind() {
	case reflect.Slice:
		et := enums.Type().Elem()
		if !validBoolFun(fv, et) {
			panic(fmt.Sprintf("%s is not a valid type for fun %s", et, fv))
		}

		cnt := 0
		for i := 0; i < enums.Len(); i++ {
			res := fv.Call([]reflect.Value{enums.Index(i)})[0]
			if res.Bool() == true {
				cnt++
			}
		}
		return cnt
	case reflect.Map:
		kt := enums.Type().Key()
		et := enums.Type().Elem()
		if !validBoolFun(fv, kt, et) {
			panic(fmt.Sprintf("func %s is invalid", fun))
		}

		cnt := 0
		for _, k := range enums.MapKeys() {
			v := fv.Call([]reflect.Value{k, enums.MapIndex(k)})[0]
			if v.Bool() == true {
				cnt++
			}
		}

		return cnt
	default:
		panic(fmt.Sprintf("%s does not support Count", enums.Type()))
	}
}
