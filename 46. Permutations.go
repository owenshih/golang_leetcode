package golang

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	answerLen := len(nums)
	loop(answerLen, nums, cur, &res)
	return res
}

func loop(answerLen int, nums, cur []int, res *[][]int) {
	for i := 0; i < len(nums); i++ {
		tmpCur := make([]int, len(cur))
		copy(tmpCur, cur)
		tmpCur = append(tmpCur, nums[i])
		if answerLen == len(tmpCur) {
			*res = append(*res, tmpCur)
			return
		}
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		newNums = append(newNums[:i], newNums[i+1:]...)
		loop(answerLen, newNums, tmpCur, res)
	}
}
