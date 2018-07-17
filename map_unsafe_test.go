// +build unsafe

package clise

import (
	"strconv"
	"testing"

	"bitbucket.org/shu_go/gotwant"
)

func TestMap(t *testing.T) {
	t.Run("", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		gotwant.Test(
			t,
			Map(slice, func(i int) interface{} {
				return "value" + strconv.Itoa(slice[i])
			}).([]string),
			[]string{"value1", "value2", "value3", "value4", "value5"},
		)
	})
	t.Run("Empty", func(t *testing.T) {
		slice := []int{}
		gotwant.Test(
			t,
			Map(slice, func(i int) interface{} {
				return "value" + strconv.Itoa(slice[i])
			}),
			nil,
		)
	})
}

func BenchmarkMap(b *testing.B) {
	slice := genIntSeq(sliceSize)
	for i := 0; i < b.N; i++ {
		Map(slice, func(i int) interface{} {
			return "value" + strconv.Itoa(slice[i])
		})
	}
}

func BenchmarkMapManual(b *testing.B) {
	slice := genIntSeq(sliceSize)
	for i := 0; i < b.N; i++ {
		sslice := make([]string, len(slice))
		for i := range slice {
			sslice[i] = "value" + strconv.Itoa(slice[i])
		}
	}
}
