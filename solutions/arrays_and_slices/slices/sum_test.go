package slices

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		actual := Sum(nums)
		expected := 15

		assertCorrectMessageSum(t, actual, expected)
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3}
		actual := Sum(nums)
		expected := 6

		assertCorrectMessageSum(t, actual, expected)
	})

	t.Run("empty slice", func(t *testing.T) {
		nums := []int{}
		actual := Sum(nums)
		expected := 0

		assertCorrectMessageSum(t, actual, expected)
	})

	t.Run("slice with one element", func(t *testing.T) {
		nums := []int{1}
		actual := Sum(nums)
		expected := 1

		assertCorrectMessageSum(t, actual, expected)
	})

	t.Run("slice with negative nums and zero", func(t *testing.T) {
		nums := []int{0, -1, 3}
		actual := Sum(nums)
		expected := 2

		assertCorrectMessageSum(t, actual, expected)
	})
}

func assertCorrectMessageSum(t testing.TB, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual is '%v', expected is '%v'",
			actual, expected)
	}
}
