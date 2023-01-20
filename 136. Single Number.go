package golang

func singleNumber(nums []int) int {
	numsMap := make(map[int]int)
	for i := range nums {
		numsMap[nums[i]]++
	}
	for i, v := range numsMap {
		if v == 1 {
			return i
		}
	}
	return 0
}
