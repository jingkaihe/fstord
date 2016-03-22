package fstord

import "reflect"

func preFetchValues(enumerable, fun interface{}) (enums, funValue reflect.Value) {
	enums = reflect.ValueOf(enumerable)
	funValue = reflect.ValueOf(fun)
	return
}

func validBoolFun(fun reflect.Value, types ...reflect.Type) bool {
	return validFun(fun, types...) && fun.Type().Out(0).Kind() == reflect.Bool
}

func validReduceFun(fun reflect.Value, types ...reflect.Type) bool {
	tps := types[:(len(types) - 1)]
	it := types[len(types)-1]
	return validFun(fun, tps...) && fun.Type().Out(0) == it
}

func validFun(fun reflect.Value, types ...reflect.Type) bool {
	// invalid if fun is not a function
	if fun.Kind() != reflect.Func {
		return false
	}

	// invalid if the params count is not the same
	if fun.Type().NumIn() != len(types) {
		return false
	}

	// invalid if the out count is not 1 (for now)
	if fun.Type().NumOut() != 1 {
		return false
	}

	// invalid if any of the params type doesn't match
	for i := 0; i < len(types); i++ {
		if fun.Type().In(i) != types[i] {
			return false
		}
	}

	return true
}
