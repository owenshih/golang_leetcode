package golang

func maxSubArray(nums []int) int {
	maxVal := nums[0]
	for i := range nums {
		if i == 0 {
			continue
		}
		if (nums[i] + nums[i-1]) > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		} else {
			nums[i] = nums[i]
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return maxVal
}
