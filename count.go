package fstord

import (
	"fmt"
	"reflect"
)

// Count returns the number of elements of a slice/map that return true when invoke fun
func Count(enumerable, fun interface{}) int {
	enums, funValue := preFetchValues(enumerable, fun)

	switch enums.Kind() {
	case reflect.Slice:
		return countOnSlice(enums, funValue)
	case reflect.Map:
		return countOnMap(enums, funValue)
	default:
		panic(fmt.Sprintf("%s does not support Count", enums.Type()))
	}
}

func countOnSlice(enums, funValue reflect.Value) int {
	et := enums.Type().Elem()
	if !validBoolFun(funValue, et) {
		panic(fmt.Sprintf("%s is not a valid type for fun %s", et, funValue))
	}

	cnt := 0
	for i := 0; i < enums.Len(); i++ {
		res := funValue.Call([]reflect.Value{enums.Index(i)})[0]
		if res.Bool() == true {
			cnt++
		}
	}
	return cnt
}

func countOnMap(enums, funValue reflect.Value) int {
	kt := enums.Type().Key()
	et := enums.Type().Elem()
	if !validBoolFun(funValue, kt, et) {
		panic(fmt.Sprintf("func %s is invalid", funValue))
	}

	cnt := 0
	for _, k := range enums.MapKeys() {
		v := funValue.Call([]reflect.Value{k, enums.MapIndex(k)})[0]
		if v.Bool() == true {
			cnt++
		}
	}
	return cnt
}
