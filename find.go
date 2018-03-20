package clise

import (
	"reflect"
)

func Find(slice interface{}, dest interface{}, funcs ...FilterFunc) bool {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("not a slice")
	}

	dv := reflect.ValueOf(dest)
	if dv.Kind() != reflect.Ptr {
		panic("not a pointer")
	}

	if dv.Type().Elem() != rv.Type().Elem() { // typeof *dv != typeof rv[0]
		panic("type not match")
	}

	if !dv.Elem().CanSet() {
		panic("dest is not CanSet, pass dest like &someValue")
	}

	length := rv.Len()

	for i := 0; i < length; i++ {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			dv.Elem().Set(rv.Index(i)) // *dv = rv[i]
			return true
		}
	}

	return false
}
