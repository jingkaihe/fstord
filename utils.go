package fstord

import "reflect"

func preFetchValues(enumerable, fun interface{}) (enums, funValue reflect.Value) {
	enums = reflect.ValueOf(enumerable)
	funValue = reflect.ValueOf(fun)
	return
}
