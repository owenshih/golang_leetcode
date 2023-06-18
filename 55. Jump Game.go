package golang

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	res := jumpdfs(nums, 0)
	if res {
		return true
	}
	return false
}

func jumpdfs(nums []int, index int) bool {
	if index == len(nums)-1 {
		return true
	}
	if nums[index] == 0 || nums[index] == -1 {
		nums[index] = -1
		return false
	}
	for j := 1; j <= nums[index]; j++ {
		res := jumpdfs(nums, index+j)
		if res {
			return true
		}
	}
	nums[index] = -1
	return false
}
