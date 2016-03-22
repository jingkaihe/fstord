package fstord

import (
	"fmt"
	"reflect"
)

// Every returns whether all the elements of a slice/map return true when invoke fun
// How the code works is kind like the Any func, but IMO is a good duplication
func Every(enumerable, fun interface{}) bool {
	enums, fv := preFetchValues(enumerable, fun)

	switch enums.Kind() {
	case reflect.Slice:
		return everyOnSlice(enums, fv)
	case reflect.Map:
		return everyOnMap(enums, fv)
	default:
		panic(fmt.Sprintf("%s does not support Every", enums.Type()))
	}
}

func everyOnSlice(enums, fv reflect.Value) bool {
	et := enums.Type().Elem()
	if !validBoolFun(fv, et) {
		panic(fmt.Sprintf("%s is not a valid type for fun %s", et, fv))
	}

	for i := 0; i < enums.Len(); i++ {
		res := fv.Call([]reflect.Value{enums.Index(i)})[0]
		if res.Bool() == false {
			return false
		}
	}
	return true
}

func everyOnMap(enums, fv reflect.Value) bool {
	kt := enums.Type().Key()
	et := enums.Type().Elem()
	if !validBoolFun(fv, kt, et) {
		panic(fmt.Sprintf("func %s is invalid", fv))
	}

	for _, k := range enums.MapKeys() {
		v := fv.Call([]reflect.Value{k, enums.MapIndex(k)})[0]
		if v.Bool() == false {
			return false
		}
	}

	return true
}
