package slices

import (
	"slices"
	"testing"
)

func assertSlicesEqual(t testing.TB, actual, expected []int) {
	t.Helper()

	if !slices.Equal(actual, expected) {
		t.Errorf("actual is %v, expected is %v", actual, expected)
	}
}
