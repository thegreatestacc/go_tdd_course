package slices

import (
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("sum some slices", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{3, 4})
		expected := []int{3, 7}

		assertSlicesEqual(t, actual, expected)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		actual := SumAll([]int{}, []int{})
		expected := []int{0, 0}

		assertSlicesEqual(t, actual, expected)
	})

	t.Run("sum one slice", func(t *testing.T) {
		actual := SumAll([]int{3})
		expected := []int{3}

		assertSlicesEqual(t, actual, expected)
	})

	t.Run("sum without args", func(t *testing.T) {
		actual := SumAll()
		expected := []int{}

		assertSlicesEqual(t, actual, expected)
	})
}
