package clise

import (
	"testing"

	"bitbucket.org/shu_go/gotwant"
)

func TestFind(t *testing.T) {
	ss := []int{1, 2, 3}

	var s int
	found := Find(ss, &s)
	gotwant.Test(t, found, true)
	gotwant.Test(t, s, ss[0])

	found = Find(ss, &s, func(i int) bool { return ss[i] == 2 })
	gotwant.Test(t, found, true)
	gotwant.Test(t, s, ss[1])

	olds := s
	found = Find(ss, &s, func(i int) bool { return ss[i] == 9 })
	gotwant.Test(t, found, false)
	gotwant.Test(t, s, olds)
}

func BenchmarkFind(b *testing.B) {
	b.Run("clise", func(b *testing.B) {
		ss := genIntSeq(sliceSize)
		var s int
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Find(ss, &s, func(j int) bool { return ss[j] == i })
		}
	})
	b.Run("goway", func(b *testing.B) {
		ss := genIntSeq(sliceSize)
		var s int
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			findInt(ss, &s, func(j int) bool { return ss[j] == i })
		}
	})

}

func findInt(slice []int, dest *int, funcs ...func(i int) bool) bool {
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

		*dest = slice[i]
		return true
	}
	return false
}
