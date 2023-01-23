package golang

func findPairs(nums []int, k int) int {
	pairCount := 0
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}

	for i := range numMap {
		if k == 0 {
			if numMap[i] > 1 {
				pairCount++
			}
			continue
		}
		if numMap[i+k] > 0 {
			pairCount++
		}
	}
	return pairCount
}
