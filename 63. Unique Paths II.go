package golang

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	h := len(obstacleGrid)
	w := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[h-1][w-1] == 1 {
		return 0
	}
	route := make([][]int, h)
	for i := 0; i < len(obstacleGrid); i++ {
		route[i] = make([]int, w)
	}
	route[0][0] = 1

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if (i == 0 && j == 0) || obstacleGrid[i][j] == 1 {
				continue
			}

			if i == 0 && j != 0 {
				route[i][j] = route[0][j-1]
				continue
			}
			if j == 0 && i != 0 {
				route[i][j] = route[i-1][0]
				continue
			}
			route[i][j] = route[i-1][j] + route[i][j-1]
		}
	}
	return route[h-1][w-1]
}
