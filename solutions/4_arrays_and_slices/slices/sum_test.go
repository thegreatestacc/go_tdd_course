package slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		actual := Sum(nums)
		expected := 15

		if actual != expected {
			t.Errorf("acutal is %d, expected is %d",
				actual, expected)
		}
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3}
		actual := Sum(nums)
		expected := 6

		if actual != expected {
			t.Errorf("acutal is %d, expected is %d",
				actual, expected)
		}
	})
}

func Sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return sum
}
