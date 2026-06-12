package slices

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, expected, actual []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual %v expected %v", actual, expected)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, actual, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, actual, expected)
	})
}

func SumAllTails(numbersToSum ...[]int) (result []int) {
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			result = append(result, 0)
		} else {
			tail := nums[1:]
			result = append(result, Sum(tail))
		}
	}
	return result
}
