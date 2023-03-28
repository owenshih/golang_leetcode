package golang

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	nMax := len(grid[0])
	mMax := len(grid)
	for n := 0; n < nMax; n++ {
		for m := 0; m < mMax; m++ {
			if grid[m][n] == 0 {
				continue
			}
			curArea := dfs(n, m, grid)
			if curArea > maxArea {
				maxArea = curArea
			}
		}
	}
	return maxArea
}

func dfs(n, m int, grid [][]int) int {
	islandArea := 0
	if n < 0 || m < 0 || n >= len(grid[0]) || m >= len(grid) {
		return islandArea
	}
	if grid[m][n] == 1 {
		grid[m][n] = 0
		islandArea += 1
		islandArea += dfs(n-1, m, grid)
		islandArea += dfs(n+1, m, grid)
		islandArea += dfs(n, m-1, grid)
		islandArea += dfs(n, m+1, grid)
	}
	return islandArea

}
