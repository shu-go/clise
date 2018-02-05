// +build unsafe

// Package clise provides EXPERIMENTAL generic slice functions.
package clise

import (
	"reflect"
	"unsafe"
)

// MakeRemover returns a function that removes the elements in the provided slice.
// MakeRemover requirs ptrSlice as a pointer to a slice (of any type).
//
// Returned function func(i, j int) removes elements [i, j] in ptrSlice.
func MakeRemover(ptrSlice interface{}) func(i, j int) {
	v := reflect.ValueOf(ptrSlice)
	if v.Kind() != reflect.Ptr { // typeof ptrSlice == ptr
		panic("not a pointer")
	} else if v.Elem().Kind() != reflect.Slice { // typeof *ptrSlice != []xxx
		panic("not a pointer to a slice")
	}

	// sizeof switch (sizes of builtin types and upto 64 bytes)
	switch reflect.ValueOf(ptrSlice).Type().Elem().Elem().Size() { // sizeof((*ptrSlice)[0])
	case 1:
		s := (*[]int8)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 2:
		s := (*[]int16)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 4:
		s := (*[]int32)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 8:
		s := (*[]int64)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 12:
		s := (*[][12]byte)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 16:
		s := (*[][16]byte)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 24:
		s := (*[][24]byte)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 32:
		s := (*[][32]byte)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 40:
		s := (*[][40]byte)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 48:
		s := (*[][48]byte)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 56:
		s := (*[][56]byte)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	case 64:
		s := (*[][64]byte)(unsafe.Pointer(v.Pointer()))
		return func(i, j int) {
			*s = append((*s)[:i], (*s)[j+1:]...)
		}
	}

	// fallback to reflect
	sv := reflect.ValueOf(ptrSlice)
	return func(i, j int) {
		elem := sv.Elem()
		elem.Set(reflect.AppendSlice(elem.Slice(0, i), elem.Slice(j+1, elem.Len())))
	}
}

// MakeCopier returns a function that copies the elements between the provided slices.
// MakeCopier requirs both slice has same length.
//
// Returned function func(srcI, destI int) copies an element from srcSlice[srcI] to destSlice[destI].
func MakeCopier(srcSlice, destSlice interface{}) func(srcI, destI int) {
	v := reflect.ValueOf(srcSlice)
	if v.Kind() != reflect.Slice { // typeof srcSlice != []xxx
		panic("src is not a slice")
	}
	sv := reflect.New(v.Type()) // var sv *[]xxx
	sv.Elem().Set(v)            // *sv = srcSlice

	v = reflect.ValueOf(destSlice)
	if v.Kind() != reflect.Slice { // typeof destSlice != []xxx
		panic("dest is not a slice")
	}
	dv := reflect.New(v.Type())
	dv.Elem().Set(v)

	// sizeof switch (sizes of builtin types and upto 64 bytes)
	switch reflect.ValueOf(srcSlice).Type().Elem().Size() { // sizeof((srcSlice)[0])
	case 1:
		s := *(*[]int8)(unsafe.Pointer(sv.Pointer()))
		d := *(*[]int8)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 2:
		s := *(*[]int16)(unsafe.Pointer(sv.Pointer()))
		d := *(*[]int16)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 4:
		s := *(*[]int32)(unsafe.Pointer(sv.Pointer()))
		d := *(*[]int32)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 8:
		s := *(*[]int64)(unsafe.Pointer(sv.Pointer()))
		d := *(*[]int64)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 12:
		s := *(*[][12]byte)(unsafe.Pointer(sv.Pointer()))
		d := *(*[][12]byte)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 16:
		s := *(*[][16]byte)(unsafe.Pointer(sv.Pointer()))
		d := *(*[][16]byte)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 24:
		s := *(*[][24]byte)(unsafe.Pointer(sv.Pointer()))
		d := *(*[][24]byte)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 32:
		s := *(*[][32]byte)(unsafe.Pointer(sv.Pointer()))
		d := *(*[][32]byte)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 40:
		s := *(*[][40]byte)(unsafe.Pointer(sv.Pointer()))
		d := *(*[][40]byte)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 48:
		s := *(*[][48]byte)(unsafe.Pointer(sv.Pointer()))
		d := *(*[][48]byte)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 56:
		s := *(*[][56]byte)(unsafe.Pointer(sv.Pointer()))
		d := *(*[][56]byte)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	case 64:
		s := *(*[][64]byte)(unsafe.Pointer(sv.Pointer()))
		d := *(*[][64]byte)(unsafe.Pointer(dv.Pointer()))
		return func(srcI, destI int) {
			d[destI] = s[srcI]
		}
	}

	// fallback to reflect
	sv = reflect.ValueOf(srcSlice)
	dv = reflect.ValueOf(destSlice)
	return func(srcI, destI int) {
		dv.Index(destI).Set(sv.Index(srcI))
	}
}

func MakeSetter(ptrSlice interface{}) func(i int, elem interface{}) {
	v := reflect.ValueOf(ptrSlice)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice { // typeof srcSlice != []xxx
		panic("src is not a pointer to a slice")
	}

	switch v.Type().Elem().Elem().Size() { // sizeof(*(ptrSlice)[0])
	case 1:
		s := (*[]int8)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*int8)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 2:
		s := (*[]int16)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*int16)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 4:
		s := (*[]int32)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*int32)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 8:
		s := (*[]int64)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*int64)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 12:
		s := (*[][12]byte)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[12]byte)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 16:
		s := (*[][16]byte)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[16]byte)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 24:
		s := (*[][24]byte)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[24]byte)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 32:
		s := (*[][32]byte)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[32]byte)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 40:
		s := (*[][40]byte)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[40]byte)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 48:
		s := (*[][48]byte)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[48]byte)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 56:
		s := (*[][56]byte)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[56]byte)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	case 64:
		s := (*[][64]byte)(unsafe.Pointer(v.Pointer()))
		return func(i int, elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[64]byte)(unsafe.Pointer(iface[1]))
			(*s)[i] = e
		}
	}

	// fallback to reflect
	return func(i int, elem interface{}) {
		v.Elem().Index(i).Set(reflect.ValueOf(elem))
	}
}

// MakeAppender returns a function that appends an elements to the provided slices.
// MakeAppender requirs ptrSlice as a pointer to a slice (of any type).
//
// Returned function func(elem interface{}) appends elem to ptrSlice.
func MakeAppender(ptrSlice interface{}) func(elem interface{}) {
	v := reflect.ValueOf(ptrSlice)
	if v.Kind() != reflect.Ptr { // typeof ptrSlice == ptr
		panic("not a pointer")
	} else if v.Elem().Kind() != reflect.Slice { // typeof *ptrSlice != []xxx
		panic("not a pointer to a slice")
	}

	// sizeof switch (sizes of builtin types and upto 64 bytes)
	switch v.Type().Elem().Elem().Size() { // sizeof((*ptrSlice)[0])
	case 1:
		s := (*[]int8)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*int8)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 2:
		s := (*[]int16)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*int16)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 4:
		s := (*[]int32)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*int32)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 8:
		s := (*[]int64)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*int64)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 12:
		s := (*[][12]byte)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[12]byte)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 16:
		s := (*[][16]byte)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[16]byte)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 24:
		s := (*[][24]byte)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[24]byte)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 32:
		s := (*[][32]byte)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[32]byte)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 40:
		s := (*[][40]byte)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[40]byte)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 48:
		s := (*[][48]byte)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[48]byte)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 56:
		s := (*[][56]byte)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[56]byte)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	case 64:
		s := (*[][64]byte)(unsafe.Pointer(v.Pointer()))
		return func(elem interface{}) {
			iface := *(*[2]uintptr)(unsafe.Pointer(&elem))
			e := *(*[64]byte)(unsafe.Pointer(iface[1]))
			*s = append(*s, e)
		}
	}

	// fallback to reflect
	return func(elem interface{}) {
		ev := reflect.ValueOf(elem)
		v.Elem().Set(reflect.Append(v.Elem(), ev)) // *s = append(*s, elem)
	}
}
