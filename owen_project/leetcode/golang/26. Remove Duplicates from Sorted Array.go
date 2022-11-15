package golang

func removeDuplicates(nums []int) int {
	key := 0
	currentInt := 0
	for index, value := range nums {
		if value == currentInt && index != 0 {
			continue
		}
		nums[key] = nums[index]
		currentInt = nums[index]
		key++
	}
	return key
}
