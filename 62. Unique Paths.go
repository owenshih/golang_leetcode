package golang

func uniquePaths(m int, n int) int {
	maps := make([][]int, n)
	for i := 0; i < len(maps); i++ {
		maps[i] = make([]int, m)
		if i == 0 {
			for j := 0; j < m; j++ {
				maps[0][j] = 1
			}
		} else {
			maps[i][0] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			maps[j][i] = maps[j][i-1] + maps[j-1][i]
		}
	}
	return maps[n-1][m-1]
}
