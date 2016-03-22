package fstord

import (
	"fmt"
	"reflect"
)

// Count returns the number of elements of a slice/map that return true when invoke fun
func Count(enumerable, fun interface{}) int {
	enums, fv := preFetchValues(enumerable, fun)

	switch enums.Kind() {
	case reflect.Slice:
		return countOnSlice(enums, fv)
	case reflect.Map:
		return countOnMap(enums, fv)
	default:
		panic(fmt.Sprintf("%s does not support Count", enums.Type()))
	}
}

func countOnSlice(enums, fv reflect.Value) int {
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
}

func countOnMap(enums, fv reflect.Value) int {
	kt := enums.Type().Key()
	et := enums.Type().Elem()
	if !validBoolFun(fv, kt, et) {
		panic(fmt.Sprintf("func %s is invalid", fv))
	}

	cnt := 0
	for _, k := range enums.MapKeys() {
		v := fv.Call([]reflect.Value{k, enums.MapIndex(k)})[0]
		if v.Bool() == true {
			cnt++
		}
	}
	return cnt
}
