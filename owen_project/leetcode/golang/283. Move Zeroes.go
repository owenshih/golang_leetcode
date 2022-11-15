package golang

func moveZeroes(nums []int) {
	key := 0
	for _, value := range nums {
		if value != 0 {
			nums[key] = value
			key++
		}
	}
	for i := key; i < len(nums); i++ {
		nums[i] = 0
	}
}
