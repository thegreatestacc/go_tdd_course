package slices

func SumAllTails(numbersToSum ...[]int) []int {
	result := make([]int, 0, len(numbersToSum))
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
