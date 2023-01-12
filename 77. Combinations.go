package golang

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	loop(1, n, k, cur, &res)
	return res
}

func loop(startNum, maxNum, maxLen int, cur []int, res *[][]int) {
	if startNum == maxNum+1 {
		return
	}
	for i := startNum; i <= maxNum; i++ {
		tmpCur := make([]int, len(cur))
		copy(tmpCur, cur)
		tmpCur = append(tmpCur, i)
		if len(tmpCur) == maxLen {
			*res = append(*res, tmpCur)
			continue
		}
		loop(i+1, maxNum, maxLen, tmpCur, res)
	}
}
