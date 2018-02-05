package clise

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"bitbucket.org/shu/gotwant"
)

const (
	sliceSize = 10000
	pause     = false
)

var (
	stopTimer  func(*testing.B) = func(*testing.B) {}
	startTimer func(*testing.B) = func(*testing.B) {}
)

func init() {
	fmt.Printf("sliceSize: %d\n", sliceSize)
	fmt.Printf("sizeof myStruct: %d\n", reflect.TypeOf(myStruct{}).Size())
	fmt.Printf("sizeof bigStruct: %d\n", reflect.TypeOf(bigStruct{}).Size())

	if pause {
		stopTimer = func(b *testing.B) { b.StopTimer() }
		startTimer = func(b *testing.B) { b.StartTimer() }
	}
}

func TestIntInGoWay(t *testing.T) {
	slice := []int{0}
	for i := len(slice) - 1; i >= 0; i-- {
		v := slice[i]
		if v%2 != 0 {
			slice = append(slice[:i], slice[i+1:]...)
			//slice = append(slice, v)
		}
	}
	gotwant.Test(t,
		slice,
		[]int{0})

	slice = []int{1}
	for i := len(slice) - 1; i >= 0; i-- {
		v := slice[i]
		if v%2 != 0 {
			slice = append(slice[:i], slice[i+1:]...)
			//slice = append(slice, v)
		}
	}
	gotwant.Test(t,
		slice,
		[]int{})

	slice = []int{0, 1, 2, 3, 4}
	for i := len(slice) - 1; i >= 0; i-- {
		v := slice[i]
		if v%2 != 0 {
			slice = append(slice[:i], slice[i+1:]...)
			//slice = append(slice, v)
		}
	}
	gotwant.Test(t,
		slice,
		[]int{0, 2, 4})

	slice = []int{0, 1, 2, 3, 4, 5}
	for i := len(slice) - 1; i >= 0; i-- {
		v := slice[i]
		if v%2 != 0 {
			slice = append(slice[:i], slice[i+1:]...)
			//slice = append(slice, v)
		}
	}
	gotwant.Test(t,
		slice,
		[]int{0, 2, 4})

	slice = []int{5, 4, 3, 2, 1, 0}
	for i := len(slice) - 1; i >= 0; i-- {
		v := slice[i]
		if v%2 != 0 {
			slice = append(slice[:i], slice[i+1:]...)
			//slice = append(slice, v)
		}
	}
	gotwant.Test(t,
		slice,
		[]int{4, 2, 0})

}

func BenchmarkReinitialization(b *testing.B) {
	orig := genIntSeq(sliceSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//stopTimer(b)
		slice := make([]int, len(orig))
		copy(slice, orig)
		//startTimer(b)
	}
}

func BenchmarkMakeInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//stopTimer(b)
		_ = make([]int, sliceSize)
		//startTimer(b)
	}
}

func BenchmarkMakeMyStruct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//stopTimer(b)
		_ = make([]myStruct, sliceSize)
		//startTimer(b)
	}
}

func genInt8Seq(n int) []int8 {
	s := make([]int8, n, n)
	for i := 0; i < n; i++ {
		s[i] = int8(i % 256)
	}
	return s
}

func genInt64Seq(n int) []int64 {
	s := make([]int64, n, n)
	for i := 0; i < n; i++ {
		s[i] = int64(i)
	}
	return s
}

func genIntSeq(n int) []int {
	s := make([]int, n, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

func genStringSeq(n int) []string {
	s := make([]string, n, n)
	for i := 0; i < n; i++ {
		s[i] = fmt.Sprintf("%d", i)
	}
	return s
}

type myStruct struct {
	name           string
	value1, value2 int
	value3         float64
}

func genMyStruct(n int) []myStruct {
	s := make([]myStruct, n, n)
	for i := 0; i < n; i++ {
		s[i] = myStruct{
			name:   strconv.Itoa(i),
			value1: i,
			value2: i * 2,
			value3: float64(i),
			//value4: float64(i) * 2,
		}
	}
	return s
}

type bigStruct struct {
	name     string
	hogehoge [100]byte
}

func genBigStruct(n int) []bigStruct {
	s := make([]bigStruct, n, n)
	for i := 0; i < n; i++ {
		s[i] = bigStruct{
			name: strconv.Itoa(i),
		}
	}
	return s
}
