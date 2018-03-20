// +build !unsafe

package clise

import (
	"reflect"
)

// CopyFiltered filters the provided slice into a new slice.
// The original slice is kept unchanged.
func CopyFiltered(slice interface{}, funcs ...FilterFunc) interface{} {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("not a slice")
	}

	length := rv.Len()

	ptrfiltered := reflect.New(rv.Type())
	ptrfiltered.Elem().Set(
		//reflect.MakeSlice(rv.Type(), 0, length))
		reflect.MakeSlice(rv.Type(), length, length)) // copy is done by dest[j] = src[i], so it's allocated in advance
	filtered := ptrfiltered.Elem()

	reflect.Copy(filtered, rv)

	Filter(ptrfiltered.Interface(), funcs...)

	return filtered.Interface()
}
