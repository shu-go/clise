package clise

import "reflect"

// Reverse reverses the provided slice ([]string etc.) in place.
func Reverse(slice interface{}) {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice { // typeof slice != []xxx
		panic("not a slice")
	}

	length := rv.Len() // len(slice)
	if length == 0 {
		return
	}

	swap := reflect.Swapper(rv.Interface())
	for i := 0; i < length/2; i++ {
		swap(i, length-i-1)
	}
}
