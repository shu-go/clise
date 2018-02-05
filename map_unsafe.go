// +build unsafe

package clise

import (
	"reflect"
)

func Map(slice interface{}, mapper func(i int) interface{}) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic("not a slice")
	}

	length := v.Len()
	if length == 0 {
		return nil
	}

	var dv reflect.Value
	//var apend func(elem interface{})
	var set func(i int, elem interface{})

	for i := 0; i < length; i++ {
		elem := mapper(i)
		if i == 0 {
			dtyp := reflect.SliceOf(reflect.TypeOf(elem))

			dv = reflect.New(dtyp) // slice *[]xxx
			//dv.Elem().Set(reflect.MakeSlice(dtyp, 0, length)) // *slice = make([]xxx, length)
			dv.Elem().Set(reflect.MakeSlice(dtyp, length, length)) // *slice = make([]xxx, length)

			//apend = MakeAppender(dv.Interface())
			set = MakeSetter(dv.Interface())
		}
		//apend(elem)
		set(i, elem)
	}

	return dv.Elem().Interface()
}
