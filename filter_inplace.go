package clise

import (
	"reflect"
)

// Filter filters the provided slice (pointer to slice : *[]string etc.) in place.
func Filter(ptrSlice interface{}, funcs ...FilterFunc) {
	rv := reflect.ValueOf(ptrSlice)
	if rv.Kind() != reflect.Ptr { // typeof ptrSlice != ptr
		panic("not a pointer")
	} else if rv.Elem().Kind() != reflect.Slice { // typeof *ptrSlice != []xxx
		panic("not a pointer to a slice")
	}

	length := rv.Elem().Len() // len(*ptrSlice)
	if length == 0 {
		return
	}

	movelist := make([]int, length)
	okindex := 0

	for i := 0; i < length; i++ {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			movelist[i] = okindex
			okindex++
		} else {
			movelist[i] = -1
		}
	}

	if okindex == length {
		return
	}

	swap := reflect.Swapper(rv.Elem().Interface())

	for i, v := range movelist {
		if v != -1 {
			swap(i, v)
		}
	}

	rv.Elem().SetLen(okindex)
}
