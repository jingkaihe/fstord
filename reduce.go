package fstord

import (
	"fmt"
	"reflect"
)

// Reduce for each element in the enumerable. The accumulated and element
// are passed into fun as arguments.
func Reduce(enumerable, fun, initial interface{}) interface{} {
	mvs := reflect.ValueOf(enumerable)
	fv := reflect.ValueOf(fun)
	it := reflect.ValueOf(initial)

	switch mvs.Kind() {
	case reflect.Slice:
		et := mvs.Type().Elem()
		if !validReduceFun(fv, it.Type(), et, it.Type()) {
			panic(fmt.Sprintf("fun %s is not a valid type", fv))
		}

		rt := it
		for i := 0; i < mvs.Len(); i++ {
			rt = fv.Call([]reflect.Value{rt, mvs.Index(i)})[0]
		}
		return rt.Interface()
	default:
		panic(fmt.Sprintf("%s does not support Map", mvs.Type()))
	}
}

func validReduceFun(fun reflect.Value, types ...reflect.Type) bool {
	tps := types[:(len(types) - 1)]
	it := types[len(types)-1]
	return validFun(fun, tps...) && fun.Type().Out(0) == it
}
