package golang

func search(nums []int, target int) int {
	min, mid, max := 0, 0, len(nums)-1
	finished := false
	for !finished {
		mid = (min + max) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[min] == target {
			return min
		}
		if nums[max] == target {
			return max
		}
		if mid == 0 || max-1 == min || min >= max {
			finished = true
		}
		if nums[mid] > target {
			max = mid - 1
		}
		if nums[mid] < target {
			min = mid + 1
		}
	}
	return -1
}
