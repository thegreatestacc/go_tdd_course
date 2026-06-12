package slices

import (
	"slices"
	"testing"
)

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}

	if !slices.Equal(actual, expected) {
		t.Errorf("actual is %v, expected is %v",
			actual, expected)
	}
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}
	return result
}
