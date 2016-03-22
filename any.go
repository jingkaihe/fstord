package fstord

import (
	"fmt"
	"reflect"
)

// Any returns whether any elements of a slice/map return true when invoke fun
func Any(enumerable, fun interface{}) bool {
	enums, fv := preFetchValues(enumerable, fun)

	switch enums.Kind() {
	case reflect.Slice:
		et := enums.Type().Elem()
		if !validBoolFun(fv, et) {
			panic(fmt.Sprintf("%s is not a valid type for fun %s", et, fv))
		}

		for i := 0; i < enums.Len(); i++ {
			res := fv.Call([]reflect.Value{enums.Index(i)})[0]
			if res.Bool() == true {
				return true
			}
		}
		return false
	case reflect.Map:
		kt := enums.Type().Key()
		et := enums.Type().Elem()
		if !validBoolFun(fv, kt, et) {
			panic(fmt.Sprintf("func %s is invalid", fun))
		}

		for _, k := range enums.MapKeys() {
			v := fv.Call([]reflect.Value{k, enums.MapIndex(k)})[0]
			if v.Bool() == true {
				return true
			}
		}

		return false
	default:
		panic(fmt.Sprintf("%s does not support Any", enums.Type()))
	}
}
