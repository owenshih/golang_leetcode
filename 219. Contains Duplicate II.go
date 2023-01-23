package golang

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]interface{})
	for i, v := range nums {
		val, _ := numMap[v]
		if val == nil {
			numMap[v] = i
			continue
		}
		checkVal := val.(int) - i
		if checkVal < 0 {
			checkVal = 0 - checkVal
		}
		if checkVal <= k {
			return true
		}
		numMap[v] = i
	}
	return false
}
