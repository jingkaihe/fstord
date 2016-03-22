package fstord

import (
	"fmt"
	"reflect"
)

// Every returns whether all the elements of a slice/map return true when invoke fun
// How the code works is kind like the Any func, but IMO is a good duplication
func Every(enumerable, fun interface{}) bool {
	enums, funValue := preFetchValues(enumerable, fun)

	switch enums.Kind() {
	case reflect.Slice:
		return everyOnSlice(enums, funValue)
	case reflect.Map:
		return everyOnMap(enums, funValue)
	default:
		panic(fmt.Sprintf("%s does not support Every", enums.Type()))
	}
}

func everyOnSlice(enums, funValue reflect.Value) bool {
	et := enums.Type().Elem()
	if !validBoolFun(funValue, et) {
		panic(fmt.Sprintf("%s is not a valid type for fun %s", et, funValue))
	}

	for i := 0; i < enums.Len(); i++ {
		res := funValue.Call([]reflect.Value{enums.Index(i)})[0]
		if res.Bool() == false {
			return false
		}
	}
	return true
}

func everyOnMap(enums, funValue reflect.Value) bool {
	kt := enums.Type().Key()
	et := enums.Type().Elem()
	if !validBoolFun(funValue, kt, et) {
		panic(fmt.Sprintf("func %s is invalid", funValue))
	}

	for _, k := range enums.MapKeys() {
		v := funValue.Call([]reflect.Value{k, enums.MapIndex(k)})[0]
		if v.Bool() == false {
			return false
		}
	}

	return true
}
