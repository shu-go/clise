package clise

import (
	"fmt"
	"strings"
	"testing"

	"bitbucket.org/shu/gotwant"
)

func ExampleCopyFiltered() {
	slice := []string{"a", "bb", "ccc", "dddd"}

	filtered := CopyFiltered(slice, // a slice
		func(i int) bool { return len(slice[i]) >= 2 },
		func(i int) bool { return slice[i][0] <= 'c' },
	).([]string)

	fmt.Printf("filtered = %#v\n", filtered)
	fmt.Printf("slice = %#v\n", slice)
	// Output: filtered = []string{"bb", "ccc"}
	// slice = []string{"a", "bb", "ccc", "dddd"}
}

func ExampleCopySliceInGoWay() {
	slice := []string{"a", "bb", "ccc", "dddd"}

	filtered := make([]string, 0, len(slice))
	for _, v := range slice {
		if len(v) >= 2 && v[0] <= 'c' {
			filtered = append(filtered, v)
		}
	}

	fmt.Printf("filtered = %#v\n", filtered)
	fmt.Printf("slice = %#v\n", slice)
	// Output: filtered = []string{"bb", "ccc"}
	// slice = []string{"a", "bb", "ccc", "dddd"}
}

func TestCopyFilteredInt(t *testing.T) {
	slice := []int{}
	if nil != CopyFiltered(slice, func(i int) bool { return slice[i]%2 == 0 }) {
		t.Errorf("not nil")
	}

	slice = []int{0}
	gotwant.Test(t,
		CopyFiltered(slice, func(i int) bool { return slice[i]%2 == 0 }).([]int),
		[]int{0})

	slice = []int{1}
	gotwant.Test(t,
		CopyFiltered(slice, func(i int) bool { return slice[i]%2 == 0 }).([]int),
		[]int{})

	slice = []int{0, 1, 2, 3, 4}
	gotwant.Test(t,
		CopyFiltered(slice, func(i int) bool { return slice[i]%2 == 0 }).([]int),
		[]int{0, 2, 4})

	slice = []int{0, 1, 2, 3, 4, 5}
	gotwant.Test(t,
		CopyFiltered(slice, func(i int) bool { return slice[i]%2 == 0 }).([]int),
		[]int{0, 2, 4})

	slice = []int{5, 4, 3, 2, 1, 0}
	gotwant.Test(t,
		CopyFiltered(slice, func(i int) bool { return slice[i]%2 == 0 }).([]int),
		[]int{4, 2, 0})
}

func TestCopyFilteredString(t *testing.T) {
	slice := []string{"hoge", "piyo", "foo", "bar", "baz"}
	gotwant.Test(t, CopyFiltered(slice, func(i int) bool { return strings.HasPrefix(slice[i], "b") }).([]string), []string{"bar", "baz"})

	slice = []string{"hoge", "piyo", "foo", "bar", "baz"}
	gotwant.Test(t, CopyFiltered(slice, func(i int) bool { return len(slice[i]) == 4 }).([]string), []string{"hoge", "piyo"})

	slice = []string{"hoge", "piyo", "foo", "bar", "baz"}
	gotwant.Test(t,
		CopyFiltered(slice,
			func(i int) bool { return len(slice[i]) == 4 },
			func(i int) bool { return strings.HasPrefix(slice[i], "b") },
		).([]string),
		[]string{})

	slice = []string{"hoge", "piyo", "foo", "bar", "baz"}
	gotwant.Test(t,
		CopyFiltered(slice,
			func(i int) bool { return len(slice[i]) == 4 },
			func(i int) bool { return strings.HasPrefix(slice[i], "p") },
		).([]string),
		[]string{"piyo"})
}

func BenchmarkCopyFiltered(b *testing.B) {
	orig := genIntSeq(sliceSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stopTimer(b)
		slice := make([]int, len(orig))
		copy(slice, orig)
		startTimer(b)

		a := CopyFiltered(slice, func(i int) bool {
			return slice[i]%10000/1000 == 1 || slice[i]%1000/100 == 1 || slice[i]%100/10 == 1 || slice[i]%10 == 1
		}).([]int)
		a[0] = a[0]
	}
}

func BenchmarkCopyGoWay(b *testing.B) {
	orig := genIntSeq(sliceSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stopTimer(b)
		slice := make([]int, len(orig))
		copy(slice, orig)
		startTimer(b)

		filterCopyInt(slice, func(i int) bool {
			return slice[i]%10000/1000 == 1 || slice[i]%1000/100 == 1 || slice[i]%100/10 == 1 || slice[i]%10 == 1
		})
	}
}

func BenchmarkCopyFilteredString(b *testing.B) {
	orig := genStringSeq(sliceSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stopTimer(b)
		slice := make([]string, len(orig))
		copy(slice, orig)
		startTimer(b)

		CopyFiltered(slice, func(i int) bool { return strings.Contains(slice[i], "1") })
	}
}

func BenchmarkCopyGoWayString(b *testing.B) {
	orig := genStringSeq(sliceSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stopTimer(b)
		slice := make([]string, len(orig))
		copy(slice, orig)
		startTimer(b)

		filterCopyString(slice, func(i int) bool { return strings.Contains(slice[i], "1") })
	}
}

func BenchmarkCopyFilteredMyStruct(b *testing.B) {
	orig := genMyStruct(sliceSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stopTimer(b)
		slice := make([]myStruct, len(orig))
		copy(slice, orig)
		startTimer(b)

		CopyFiltered(slice, func(i int) bool { return strings.Contains(slice[i].name, "1") })
	}
}

func BenchmarkCopyGoWayMyStruct(b *testing.B) {
	orig := genMyStruct(sliceSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stopTimer(b)
		slice := make([]myStruct, len(orig))
		copy(slice, orig)
		startTimer(b)

		filterCopyMyStruct(slice, func(i int) bool { return strings.Contains(slice[i].name, "1") })
	}
}

func filterCopyInt(slice []int, funcs ...func(i int) bool) []int {
	result := make([]int, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if !allok {
			continue
		}

		result = append(result, slice[i])
	}
	return result
}

func filterCopyString(slice []string, funcs ...func(i int) bool) []string {
	result := make([]string, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if !allok {
			continue
		}

		result = append(result, slice[i])
	}
	return result
}

func filterCopyMyStruct(slice []myStruct, funcs ...func(i int) bool) []myStruct {
	result := make([]myStruct, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			continue
		}

		result = append(result, slice[i])
	}
	return result
}
