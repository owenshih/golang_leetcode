package golang

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	var currentMax int
	for _, value := range nums {
		if 1 == value {
			currentMax += 1
		} else {
			currentMax = 0
		}
		if currentMax > max {
			max = currentMax
		}
	}
	return max
}
