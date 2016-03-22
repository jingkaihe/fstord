package fstord

import (
	"fmt"
	"reflect"
)

// Reduce for each element in the enumerable. The accumulated and element
// are passed into fun as arguments.
func Reduce(enumerable, fun, initial interface{}) interface{} {
	enums, fv := preFetchValues(enumerable, fun)
	it := reflect.ValueOf(initial)

	switch enums.Kind() {
	case reflect.Slice:
		et := enums.Type().Elem()
		if !validReduceFun(fv, it.Type(), et, it.Type()) {
			panic(fmt.Sprintf("fun %s is not a valid type", fun))
		}

		rt := it
		for i := 0; i < enums.Len(); i++ {
			rt = fv.Call([]reflect.Value{rt, enums.Index(i)})[0]
		}
		return rt.Interface()
	case reflect.Map:
		kt := enums.Type().Key()
		et := enums.Type().Elem()

		if !validReduceFun(fv, it.Type(), kt, et, it.Type()) {
			panic(fmt.Sprintf("fun %s is not a valid type for map enumeration", fun))
		}

		rt := it
		for _, k := range enums.MapKeys() {
			v := enums.MapIndex(k)
			rt = fv.Call([]reflect.Value{rt, k, v})[0]
		}

		return rt.Interface()
	default:
		panic(fmt.Sprintf("%s does not support Reduce", enums.Type()))
	}
}

func validReduceFun(fun reflect.Value, types ...reflect.Type) bool {
	tps := types[:(len(types) - 1)]
	it := types[len(types)-1]
	return validFun(fun, tps...) && fun.Type().Out(0) == it
}
