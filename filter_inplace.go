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

	remove := MakeRemover(ptrSlice)

	rmlength := length / 2
	if rmlength == 0 {
		rmlength = length
	}
	rmlist := make([]int, 0, rmlength)
	// mark in asc order
	for i := 0; i < length; i++ {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			continue
		}
		rmlist = append(rmlist, i)
	}

	// comact traversing in desc order
	lastS, lastE := -1, -1
	for i := len(rmlist) - 1; i >= 0; i-- {
		if lastE == -1 {
			lastS, lastE = rmlist[i], rmlist[i]
			continue
		}
		if rmlist[i] == lastS-1 {
			lastS = rmlist[i]
			continue
		}

		remove(lastS, lastE)

		lastS, lastE = rmlist[i], rmlist[i]
	}
	if lastE != -1 {
		remove(lastS, lastE)
	}
}

// FilterSimple is basically same to Slice, with less memory usage and less speed in most cases.
// This func removes elements one by one, not by range.
func FilterSimple(ptrSlice interface{}, funcs ...FilterFunc) {
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

	remove := MakeRemover(ptrSlice)

	for i := length - 1; i >= 0; i-- {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			continue
		}

		remove(i, i)
	}
}

// FilterSwapSimple is slow and much safe version of Slice.
func FilterSwapSimple(ptrSlice interface{}, funcs ...FilterFunc) {
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

	swap := reflect.Swapper(rv.Elem().Interface())

	ngboundary := length
	for i := length - 1; i >= 0; i-- {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			continue
		}

		if i < ngboundary-1 {
			//swap(i, ngboundary-1) //not stable
			for j := i; j < ngboundary-1; j++ {
				swap(j, j+1)
			}
		}
		ngboundary--
	}

	rv.Elem().SetLen(ngboundary)
}

func FilterSwap(ptrSlice interface{}, funcs ...FilterFunc) {
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
		if ! /* NOT */ allok {
			continue
		}

		movelist[i] = okindex
		okindex++
	}

	if okindex == length {
		return
	}

	swap := reflect.Swapper(rv.Elem().Interface())

	for i, v := range movelist {
		swap(i, v)
	}

	rv.Elem().SetLen(okindex)
}

// FilterSwapUnstable is faster version of FilterSwap.
// The order of slice may not be preserved.
func FilterSwapUnstable(ptrSlice interface{}, funcs ...FilterFunc) {
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

	swap := reflect.Swapper(rv.Elem().Interface())

	ngboundary := length
	for i := length - 1; i >= 0; i-- {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			continue
		}

		if i < ngboundary-1 {
			swap(i, ngboundary-1) //not stable
			//for j := i; j < ngboundary-1; j++ {
			//	swap(j, j+1)
			//}
		}
		ngboundary--
	}

	rv.Elem().SetLen(ngboundary)
}
