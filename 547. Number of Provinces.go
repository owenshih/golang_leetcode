package main

func findCircleNum(isConnected [][]int) int {
	result := 0
	nMax := len(isConnected)
	for n := 0; n < nMax; n++ {
		if isConnected[n][n] == 1 {
			result++
			mapTransform(isConnected, n, n)
		}
	}
	return result
}

func mapTransform(border [][]int, n, m int) {
	if m < 0 || n < 0 || m >= len(border[0]) || n >= len(border) || border[n][m] == 0 {
		return
	}
	border[n][m] = 0
	border[m][n] = 0
	for i := 0; i < len(border); i++ {
		mapTransform(border, i, n)
	}
}
