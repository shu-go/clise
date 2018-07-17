package clise

import (
	"testing"

	"bitbucket.org/shu_go/gotwant"
)

func TestReverse(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		s := []string{}
		Reverse(s)
		gotwant.Test(t, s, []string{})
	})
	t.Run("One", func(t *testing.T) {
		s := []string{"a"}
		Reverse(s)
		gotwant.Test(t, s, []string{"a"})
	})
	t.Run("Two", func(t *testing.T) {
		s := []string{"a", "b"}
		Reverse(s)
		gotwant.Test(t, s, []string{"b", "a"})
	})
	t.Run("Three", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		Reverse(s)
		gotwant.Test(t, s, []string{"c", "b", "a"})
	})
}
