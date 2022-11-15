package golang

func removeElement(nums []int, val int) int {
	key := 0
	for index, value := range nums {
		if value != val {
			nums[key] = nums[index]
			key++
		}
	}
	return key
}
