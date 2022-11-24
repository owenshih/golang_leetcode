package golang

func searchInsert(nums []int, target int) int {
	min, mid, max := 0, 0, len(nums)-1
	finished := false
	for !finished {
		if target == nums[min] {
			return min
		}
		if target == nums[max] {
			return max
		}
		if target < nums[min] {
			return min
		}
		if target > nums[max] {
			return max + 1
		}
		if max-min == 1 {
			return min + 1
		}
		mid = (min + max) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] < target && target < nums[max] {
			min = mid + 1
		}
		if nums[mid] > target && target > nums[min] {
			max = mid - 1
		}
	}
	return 0
}
