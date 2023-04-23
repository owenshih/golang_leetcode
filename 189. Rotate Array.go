package golang

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	res := append(nums[n-k:], nums[:n-k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = res[i]
	}
}
