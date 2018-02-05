package clise

import (
	"reflect"
)

// Bisect provides bisection search.
// The slice must be sorted.
// Func cmp should return -1 toward index 0, 0 matched, 1 toward last index,
// like strings.Compare("search string" slice[i]) for asc-order.
func Bisect(slice interface{}, cmp func(i int) (dir int)) int {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice { // typeof srcSlice != []xxx
		panic("src is not a slice")
	}

	length := v.Len()
	if length == 0 {
		return -1
	}

	var l, p, r int = 0, (length - 1) / 2, length - 1
	for {
		//fmt.Fprintf(os.Stderr, "l, p, r = %v, %v, %v\n", l, p, r)
		dir := cmp(p)
		//fmt.Fprintf(os.Stderr, " -> dir = %v\n", dir)
		if dir == 0 {
			return p
		}

		if l == r {
			break
		}

		if dir < 0 {
			p -= 1
			l, p, r = l, (l+p)/2, p
		} else {
			p += 1
			l, p, r = p, (p+r)/2, r
		}
	}

	return -1
}

func Cmp(l, r interface{}) int {
	switch lv := l.(type) {
	case bool:
		rv, ok := r.(bool)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if rv {
			return -1
		}
		return 1

	case float32:
		rv, ok := r.(float32)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case float64:
		rv, ok := r.(float64)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case int:
		rv, ok := r.(int)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case uint:
		rv, ok := r.(uint)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case int8:
		rv, ok := r.(int8)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case int16:
		rv, ok := r.(int16)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case int32:
		rv, ok := r.(int32)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case int64:
		rv, ok := r.(int64)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case uint8:
		rv, ok := r.(uint8)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case uint16:
		rv, ok := r.(uint16)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case uint32:
		rv, ok := r.(uint32)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case uint64:
		rv, ok := r.(uint64)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case string:
		rv, ok := r.(string)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1

	case uintptr:
		rv, ok := r.(uintptr)
		if !ok {
			return -1
		}

		if lv == rv {
			return 0
		}
		if lv < rv {
			return -1
		}
		return 1
	}
	return -1
}
