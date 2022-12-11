package golang

func minCostToMoveChips(position []int) int {
	even, odd := 0, 0
	for _, v := range position {
		if v%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}
	if even > odd {
		return odd
	} else {
		return even
	}
}
