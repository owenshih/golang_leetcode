package golang

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	prevMax := nums[0]
	curMax := getMax(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		curMax, prevMax = getMax(prevMax+nums[i], curMax), curMax
	}
	return curMax
}

func getMax(h1, h2 int) int {
	if h1 < h2 {
		return h2
	}
	return h1
}
