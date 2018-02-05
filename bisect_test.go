package clise

import (
	"fmt"
	"strings"
	"testing"

	"bitbucket.org/shu/gotwant"
)

func ExampleBisect() {
	slice := []string{"a", "bb", "ccc", "dddd"}

	pos := Bisect(slice, func(i int) int {
		return strings.Compare("ccc", slice[i])
	})
	fmt.Printf("pos = %v\n", pos)

	slice = genStringSeq(10000)
	pos = Bisect(slice, func(i int) int {
		return strings.Compare("5963", slice[i])
	})
	fmt.Printf("pos = %v\n", pos)
	fmt.Printf("slice[pos] = %v\n", slice[pos])

	slice = genStringSeq(1000)
	pos = Bisect(slice, func(i int) int {
		return strings.Compare("5963", slice[i])
	})
	fmt.Printf("pos = %v\n", pos)

	// Output: pos = 2
	// pos = 5963
	// slice[pos] = 5963
	// pos = -1
}

func TestBisect(t *testing.T) {
	slice := []int{0, 1, 2}
	pos := Bisect(slice, func(i int) int {
		return Cmp(0, slice[i])
	})
	gotwant.Test(t, pos, 0)

	slice = []int{0, 1, 2}
	pos = Bisect(slice, func(i int) int {
		return Cmp(-1, slice[i])
	})
	gotwant.Test(t, pos, -1)

	slice = []int{0, 1, 2}
	pos = Bisect(slice, func(i int) int {
		return Cmp(-1, slice[i])
	})
	gotwant.Test(t, pos, -1)

	pos = Bisect(slice, func(i int) int {
		return Cmp(2, slice[i])
	})
	gotwant.Test(t, pos, 2)

	slice = []int{2, 1, 0}
	pos = Bisect(slice, func(i int) int {
		return Cmp(slice[i], 2)
	})
	gotwant.Test(t, pos, 0)
}

func BenchmarkBisect(b *testing.B) {
	b.Run("clise", func(b *testing.B) {
		slice := genStringSeq(1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Bisect(slice, func(i int) int {
				return strings.Compare("820", slice[i])
			})
		}
	})
	b.Run("goway", func(b *testing.B) {
		slice := genStringSeq(1000)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, v := range slice {
				if strings.Compare(v, "820") == 0 {
					break
				}
			}
		}
	})
}
