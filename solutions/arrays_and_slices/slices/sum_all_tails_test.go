package slices

import (
	"testing"
)

func TestSumAllTails(t *testing.T) {
	t.Run("make the sum of some slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		assertSlicesEqual(t, actual, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		assertSlicesEqual(t, actual, expected)
	})

	t.Run("without args", func(t *testing.T) {
		actual := SumAllTails()
		expected := []int{}

		assertSlicesEqual(t, actual, expected)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{})
		expected := []int{0, 0}

		assertSlicesEqual(t, actual, expected)
	})

	t.Run("sum one slice", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2})
		expected := []int{2}

		assertSlicesEqual(t, actual, expected)
	})
}
