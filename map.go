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
	var set func(i int, elem interface{})
	//var apend func(elem interface{})

	for i := 0; i < length; i++ {
		elem := mapper(i)
		if i == 0 {
			dtyp := reflect.SliceOf(reflect.TypeOf(elem))

			dv = reflect.New(dtyp)
			dv.Elem().Set(reflect.MakeSlice(dtyp, length, length))

			//apend = MakeAppender(dv.Interface())
			set = MakeSetter(dv.Elem().Interface())
		}
		//apend(elem)
		set(i, elem)
	}

	return dv.Elem().Interface()
}
