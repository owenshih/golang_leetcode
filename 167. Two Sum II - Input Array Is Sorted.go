package golang

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for {
		if (numbers[left] + numbers[right]) == target {
			break
		}
		if (numbers[left] + numbers[right]) > target {
			right--
		} else {
			left++
		}
	}
	result := []int{left + 1, right + 1}
	return result
}
