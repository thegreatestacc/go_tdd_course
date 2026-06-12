package arrays

import "testing"

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	actual := Sum(nums)
	expected := 15

	if actual != expected {
		t.Errorf("actual is %d, expected is %d",
			actual, expected)
	}
}

func Sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return sum
}
