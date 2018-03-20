package clise

import (
	"strings"
	"testing"

	"bitbucket.org/shu/gotwant"
)

func TestFilterInt(t *testing.T) {
	t.Run("clise", func(t *testing.T) {
		slice := []int{0}
		Filter(&slice, func(i int) bool { return slice[i]%2 == 0 })
		gotwant.Test(t,
			slice,
			[]int{0})

		slice = []int{1}
		Filter(&slice, func(i int) bool { return slice[i]%2 == 0 })
		gotwant.Test(t,
			slice,
			[]int{})

		slice = []int{0, 1, 2, 3, 4}
		Filter(&slice, func(i int) bool { return slice[i]%2 == 0 })
		gotwant.Test(t,
			slice,
			[]int{0, 2, 4})

		slice = []int{0, 1, 2, 3, 4, 5}
		Filter(&slice, func(i int) bool { return slice[i]%2 == 0 })
		gotwant.Test(t,
			slice,
			[]int{0, 2, 4})

		slice = []int{5, 4, 3, 2, 1, 0}
		Filter(&slice, func(i int) bool { return slice[i]%2 == 0 })
		gotwant.Test(t,
			slice,
			[]int{4, 2, 0})
	})
}

func TestFilterString(t *testing.T) {
	t.Run("clise", func(t *testing.T) {
		slice := []string{"hoge", "piyo", "foo", "bar", "baz"}
		Filter(&slice, func(i int) bool { return strings.HasPrefix(slice[i], "b") })
		gotwant.Test(t, slice, []string{"bar", "baz"})

		slice = []string{"hoge", "piyo", "foo", "bar", "baz"}
		Filter(&slice, func(i int) bool { return len(slice[i]) == 4 })
		gotwant.Test(t, slice, []string{"hoge", "piyo"})

		slice = []string{"hoge", "piyo", "foo", "bar", "baz"}
		Filter(&slice,
			func(i int) bool { return len(slice[i]) == 4 },
			func(i int) bool { return strings.HasPrefix(slice[i], "b") },
		)
		gotwant.Test(t, slice, []string{})

		slice = []string{"hoge", "piyo", "foo", "bar", "baz"}
		Filter(&slice,
			func(i int) bool { return len(slice[i]) == 4 },
			func(i int) bool { return strings.HasPrefix(slice[i], "p") },
		)
		gotwant.Test(t, slice, []string{"piyo"})
	})
	t.Run("goway", func(t *testing.T) {
		slice := []string{"hoge", "piyo", "foo", "bar", "baz"}
		filterInPlaceString(&slice, func(i int) bool { return strings.HasPrefix(slice[i], "b") })
		gotwant.Test(t, slice, []string{"bar", "baz"})

		slice = []string{"hoge", "piyo", "foo", "bar", "baz"}
		filterInPlaceString(&slice, func(i int) bool { return len(slice[i]) == 4 })
		gotwant.Test(t, slice, []string{"hoge", "piyo"})

		slice = []string{"hoge", "piyo", "foo", "bar", "baz"}
		filterInPlaceString(&slice, func(i int) bool { return len(slice[i]) == 4 })
		filterInPlaceString(&slice, func(i int) bool { return strings.HasPrefix(slice[i], "b") })
		gotwant.Test(t, slice, []string{})

		slice = []string{"hoge", "piyo", "foo", "bar", "baz"}
		filterInPlaceString(&slice, func(i int) bool { return len(slice[i]) == 4 })
		filterInPlaceString(&slice, func(i int) bool { return strings.HasPrefix(slice[i], "p") })
		gotwant.Test(t, slice, []string{"piyo"})
	})
}

func BenchmarkFilterInt(b *testing.B) {
	b.Run("clise", func(b *testing.B) {
		orig := genIntSeq(sliceSize)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			stopTimer(b)
			slice := make([]int, len(orig))
			copy(slice, orig)
			startTimer(b)

			Filter(&slice, func(i int) bool {
				return slice[i]%10000/1000 == 1 || slice[i]%1000/100 == 1 || slice[i]%100/10 == 1 || slice[i]%10 == 1
			})
		}
	})
	b.Run("goway", func(b *testing.B) {
		orig := genIntSeq(sliceSize)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			stopTimer(b)
			slice := make([]int, len(orig))
			copy(slice, orig)
			startTimer(b)

			filterInPlaceInt(&slice, func(i int) bool {
				return slice[i]%10000/1000 == 1 || slice[i]%1000/100 == 1 || slice[i]%100/10 == 1 || slice[i]%10 == 1
			})
		}
	})
	b.Run("goway2", func(b *testing.B) {
		orig := genIntSeq(sliceSize)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			stopTimer(b)
			slice := make([]int, len(orig))
			copy(slice, orig)
			startTimer(b)

			filterInPlaceInt2(&slice, func(i int) bool {
				return slice[i]%10000/1000 == 1 || slice[i]%1000/100 == 1 || slice[i]%100/10 == 1 || slice[i]%10 == 1
			})
		}
	})
}

func BenchmarkFilterMyStruct(b *testing.B) {
	b.Run("clise", func(b *testing.B) {
		orig := genMyStruct(sliceSize)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			stopTimer(b)
			slice := make([]myStruct, len(orig))
			copy(slice, orig)
			startTimer(b)

			Filter(&slice, func(i int) bool { return strings.Contains(slice[i].name, "1") })
		}
	})
	b.Run("goway", func(b *testing.B) {
		orig := genMyStruct(sliceSize)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			stopTimer(b)
			slice := make([]myStruct, len(orig))
			copy(slice, orig)
			startTimer(b)

			filterInPlaceMyStruct2(&slice, func(i int) bool { return strings.Contains(slice[i].name, "1") })
		}
	})
}

func BenchmarkFilterGoWayBigStruct(b *testing.B) {
	b.Run("clise", func(b *testing.B) {
		orig := genBigStruct(sliceSize)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			stopTimer(b)
			slice := make([]bigStruct, len(orig))
			copy(slice, orig)
			startTimer(b)

			filterInPlaceBigStruct(&slice, func(i int) bool { return strings.Contains(slice[i].name, "1") })
		}
	})
	b.Run("goway", func(b *testing.B) {
		orig := genBigStruct(sliceSize)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			stopTimer(b)
			slice := make([]bigStruct, len(orig))
			copy(slice, orig)
			startTimer(b)

			filterInPlaceBigStruct2(&slice, func(i int) bool { return strings.Contains(slice[i].name, "1") })
		}
	})
}

func filterInPlaceInt(ptrslice *[]int, funcs ...func(i int) bool) {
	for i := len(*ptrslice) - 1; i >= 0; i-- {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			continue
		}

		*ptrslice = append((*ptrslice)[:i], (*ptrslice)[i+1:]...)
	}
}

func filterInPlaceInt2(ptrslice *[]int, funcs ...func(i int) bool) {
	length := len(*ptrslice)

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

		//remove(lastS, lastE)
		*ptrslice = append((*ptrslice)[:lastS], (*ptrslice)[lastE+1:]...)

		lastS, lastE = rmlist[i], rmlist[i]
	}
	if lastE != -1 {
		//remove(lastS, lastE)
		*ptrslice = append((*ptrslice)[:lastS], (*ptrslice)[lastE+1:]...)
	}
}

func filterInPlaceMyStruct2(ptrslice *[]myStruct, funcs ...func(i int) bool) {
	length := len(*ptrslice)

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

		//remove(lastS, lastE)
		*ptrslice = append((*ptrslice)[:lastS], (*ptrslice)[lastE+1:]...)

		lastS, lastE = rmlist[i], rmlist[i]
	}
	if lastE != -1 {
		//remove(lastS, lastE)
		*ptrslice = append((*ptrslice)[:lastS], (*ptrslice)[lastE+1:]...)
	}
}

func filterInPlaceBigStruct2(ptrslice *[]bigStruct, funcs ...func(i int) bool) {
	length := len(*ptrslice)

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

		//remove(lastS, lastE)
		*ptrslice = append((*ptrslice)[:lastS], (*ptrslice)[lastE+1:]...)

		lastS, lastE = rmlist[i], rmlist[i]
	}
	if lastE != -1 {
		//remove(lastS, lastE)
		*ptrslice = append((*ptrslice)[:lastS], (*ptrslice)[lastE+1:]...)
	}
}

func filterInPlaceString(ptrslice *[]string, funcs ...func(i int) bool) {
	for i := len(*ptrslice) - 1; i >= 0; i-- {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			continue
		}

		*ptrslice = append((*ptrslice)[:i], (*ptrslice)[i+1:]...)
	}
}

func filterInPlaceBigStruct(ptrslice *[]bigStruct, funcs ...func(i int) bool) {
	for i := len(*ptrslice) - 1; i >= 0; i-- {
		allok := true
		for _, f := range funcs {
			if !f(i) {
				allok = false
			}
		}
		if allok {
			continue
		}

		*ptrslice = append((*ptrslice)[:i], (*ptrslice)[i+1:]...)
	}
}
