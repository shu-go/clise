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
	if length == 0 {
		return nil
	}

	ptrfiltered := reflect.New(rv.Type())
	ptrfiltered.Elem().Set(
		//reflect.MakeSlice(rv.Type(), 0, length))
		reflect.MakeSlice(rv.Type(), length, length)) // copy is done by dest[j] = src[i], so it's allocated in advance
	filtered := ptrfiltered.Elem()

	//copy := MakeCopier(slice, ptrfiltered.Elem())
	copy := MakeCopier(slice, ptrfiltered.Elem().Interface())

	copiedLen := 0

	for i := 0; i < length; i++ {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if !allok {
			continue
		}

		// copy

		//filtered.Set(reflect.Append(filtered, rv.Index(i)))
		//filtered.Index(copiedLen).Set(rv.Index(i))
		copy(i, copiedLen)
		copiedLen++
	}

	// set actual length of filtered
	filtered.SetLen(copiedLen)

	return filtered.Interface()
}
