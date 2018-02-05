// +build unsafe

package clise

import (
	"fmt"
	"testing"
)

func ExampleFilter() {
	strSlice := []string{"a", "bb", "ccc", "dddd"}
	intSlice := []int{1, 2, 3, 4}

	FilterSimpleFast(&strSlice, // a pointer to a slice
		func(i int) bool { return len(strSlice[i]) >= 2 }, //          {"bb", "ccc", "dddd"}
		func(i int) bool { return strSlice[i][0] <= 'c' }, // and {"a", "bb", "ccc"}
	)
	FilterSimpleFast(&intSlice, // a pointer to a slice
		func(i int) bool { return intSlice[i]%2 == 0 }, // even
	)

	fmt.Printf("strSlice = %#v\n", strSlice)
	fmt.Printf("intSlice = %#v\n", intSlice)
	// Output: strSlice = []string{"bb", "ccc"}
	// intSlice = []int{2, 4}
}

func ExampleSliceInGoWay() {
	strSlice := []string{"a", "bb", "ccc", "dddd"}
	intSlice := []int{1, 2, 3, 4}

	for i := len(strSlice) - 1; i >= 0; i-- {
		if !(len(strSlice[i]) >= 2 && strSlice[i][0] <= 'c') {
			strSlice = append(strSlice[:i], strSlice[i+1:]...)
		}
	}
	for i := len(intSlice) - 1; i >= 0; i-- {
		if !(intSlice[i]%2 == 0) {
			intSlice = append(intSlice[:i], intSlice[i+1:]...)
		}
	}

	fmt.Printf("strSlice = %#v\n", strSlice)
	fmt.Printf("intSlice = %#v\n", intSlice)
	// Output: strSlice = []string{"bb", "ccc"}
	// intSlice = []int{2, 4}
}

func TestTypesafe(t *testing.T) {
	var panicked bool
	var slice interface{}
	slice = 1
	func() {
		panicked = false
		defer func() { recover(); panicked = true }()
		FilterSimpleFast(&slice, func(i int) bool { return true })
	}()
	if !panicked {
		t.Errorf("!?")
	}

	slice = "abc"
	func() {
		panicked = false
		defer func() { recover(); panicked = true }()
		FilterSimpleFast(&slice, func(i int) bool { return true })
	}()
	if !panicked {
		t.Errorf("!?")
	}

	slice = []int{1, 2, 3}
	func() {
		panicked = false
		defer func() { recover(); panicked = true }()
		FilterSimpleFast( /*not a ptr*/ slice, func(i int) bool { return true })
	}()
	if !panicked {
		t.Errorf("!?")
	}
}
