package golang

func fib(n int) int {
	firstTow, pre := 0, 1
	if n == 0 || n == 1 {
		return n
	}
	for i := 2; i <= n; i++ {
		if i == n {
			return firstTow + pre
		}
		firstTow, pre = pre, firstTow+pre
	}
	return 0
}
