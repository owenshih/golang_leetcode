package golang

import "fmt"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	res = append(res, make([]int, 0))
	loop(0, nums, cur, &res)
	return res
}

func loop(startIndex int, nums, cur []int, res *[][]int) {
	for i := startIndex; i < len(nums); i++ {
		tmpCur := make([]int, len(cur))
		copy(tmpCur, cur)
		tmpCur = append(tmpCur, nums[i])
		fmt.Println(tmpCur)
		*res = append(*res, tmpCur)
		loop(i+1, nums, tmpCur, res)
	}
}
