// +build unsafe

package clise

import (
	"fmt"
	"testing"
)

func ExampleMakeRemover() {
	slice := []string{"a", "bb", "ccc", "dddd"}

	remove := MakeRemover(&slice) // Note: give a pointer to the slice

	remove(1, 2)

	fmt.Printf("slice = %#v\n", slice)
	// Output: slice = []string{"a", "dddd"}
}

func ExampleMakeCopier() {
	slice := []string{"a", "bb", "ccc", "dddd"}
	dest := make([]string, 2) // Note: needs actual length, not only capacity

	copy := MakeCopier(slice, dest)

	copy(3, 0) // => dest = []string { "dddd", "" }
	copy(2, 0) // => dest = []string { "ccc", "" }
	copy(1, 1) // => dest = []string { "ccc", "bb" }

	fmt.Printf("dest = %#v\n", dest)
	// Output: dest = []string{"ccc", "bb"}
}

func ExampleMakeAppender() {
	slice := []string{"a", "bb", "ccc", "dddd"}

	apend := MakeAppender(&slice)

	apend("eeeee")

	fmt.Printf("slice = %#v\n", slice)
	// Output: slice = []string{"a", "bb", "ccc", "dddd", "eeeee"}
}

func ExampleMakeAppender2() {
	slice := []string{"a", "bb", "ccc", "dddd"}
	dest := []string{}

	apend := MakeAppender(&dest)

	for _, v := range slice {
		apend(v + "Z")
	}

	fmt.Printf("dest = %#v\n", dest)
	// Output: dest = []string{"aZ", "bbZ", "cccZ", "ddddZ"}
}

type testCaseSliceOpString struct {
	src  []string
	op   func(src []string) []string
	want []string
}

func testSliceOpString(t *testing.T, cases ...testCaseSliceOpString) {
	ok := true
	for _, c := range cases {
		dest := c.op(c.src)
		if len(dest) != len(c.want) {
			ok = false
		} else {
			for i := 0; i < len(dest); i++ {
				if dest[i] != c.want[i] {
					ok = false
					break
				}
			}
		}
		if !ok {
			t.Errorf("\nhave %#v\nwant %#v\n", dest, c.want)
		}
	}
}

func TestMakeCopier(t *testing.T) {
	testSliceOpString(t,
		testCaseSliceOpString{
			src: []string{"a"},
			op: func(src []string) []string {
				dest := make([]string, len(src))
				copy := MakeCopier(src, dest)
				copy(0, 0)
				return dest
			},
			want: []string{"a"},
		},
		testCaseSliceOpString{
			src: []string{"a", "b"},
			op: func(src []string) []string {
				dest := make([]string, len(src))
				copy := MakeCopier(src, dest)
				copy(1, 0)
				copy(0, 1)
				return dest
			},
			want: []string{"b", "a"},
		},
		testCaseSliceOpString{
			src: []string{},
			op: func(src []string) []string {
				return nil
			},
			want: nil,
		},
	)
}

func BenchmarkCopyOp(b *testing.B) {
	src := []string{"aaaaa", "bbbbb", "ccccc"}
	dest := make([]string, len(src))

	copy := MakeCopier(src, dest)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		srcI := b.N % len(src)
		destI := b.N % len(dest)
		copy(srcI, destI)
	}
}

func BenchmarkCopyOpGoWay(b *testing.B) {
	src := []string{"aaaaa", "bbbbb", "ccccc"}
	dest := make([]string, len(src))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		srcI := b.N % len(src)
		destI := b.N % len(dest)
		dest[destI] = src[srcI]
	}
}

func TestMakeRemover(t *testing.T) {
	testSliceOpString(t,
		testCaseSliceOpString{
			src: []string{},
			op: func(src []string) []string {
				_ = MakeRemover(&src)
				return src
			},
			want: []string{},
		},
		testCaseSliceOpString{
			src: []string{"a"},
			op: func(src []string) []string {
				remove := MakeRemover(&src)
				remove(0, 0)
				return src
			},
			want: []string{},
		},
		testCaseSliceOpString{
			src: []string{"a", "b", "c"},
			op: func(src []string) []string {
				remove := MakeRemover(&src)
				remove(0, 0)
				return src
			},
			want: []string{"b", "c"},
		},
		testCaseSliceOpString{
			src: []string{"a", "b", "c"},
			op: func(src []string) []string {
				remove := MakeRemover(&src)
				remove(0, 1)
				return src
			},
			want: []string{"c"},
		},
		testCaseSliceOpString{
			src: []string{"a", "b", "c"},
			op: func(src []string) []string {
				remove := MakeRemover(&src)
				remove(1, 2)
				return src
			},
			want: []string{"a"},
		},
		testCaseSliceOpString{
			src: []string{"a", "b", "c"},
			op: func(src []string) []string {
				remove := MakeRemover(&src)
				remove(0, 2)
				return src
			},
			want: []string{},
		},
		testCaseSliceOpString{
			src: []string{},
			op: func(src []string) []string {
				return nil
			},
			want: nil,
		},
	)
}

func BenchmarkRemoveOp(b *testing.B) {
	b.ResetTimer()
	//for i := 0; i < 10000; /*b.N*/ i++ {
	for i := 0; i < b.N; i++ {
		slice := genStringSeq(1000) //[]string{"aaaaa", "bbbbb", "ccccc"}
		remove := MakeRemover(&slice)

		i := b.N % len(slice)
		remove(i, i)
	}
}

func BenchmarkRemoveOpGoWay(b *testing.B) {
	b.ResetTimer()
	//for i := 0; i < 10000; /*b.N*/ i++ {
	for i := 0; i < b.N; i++ {
		slice := genStringSeq(1000) //[]string{"aaaaa", "bbbbb", "ccccc"}
		//remove := MakeRemover(&slice)

		i := b.N % len(slice)
		slice = append(slice[:i], slice[i+1:]...)
	}
}

func BenchmarkAppendOp(b *testing.B) {
	b.ResetTimer()
	//for i := 0; i < 10000; /*b.N*/ i++ {
	for i := 0; i < b.N; i++ {
		slice := []string{}
		apend := MakeAppender(&slice)

		apend("aaaa")
	}
}

func BenchmarkAppendOpGoWay(b *testing.B) {
	b.ResetTimer()
	//for i := 0; i < 10000; /*b.N*/ i++ {
	for i := 0; i < b.N; i++ {
		slice := []string{}
		//apend := MakeAppender(&slice)

		slice = append(slice, "aaaa")
	}
}
