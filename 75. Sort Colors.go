package golang

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		for i := left + 1; i <= right; i++ {
			if nums[left] > nums[i] {
				nums[left], nums[i] = nums[i], nums[left]
			}
		}
		left++
	}
}
