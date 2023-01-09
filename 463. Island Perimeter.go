package golang

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	mMax, nMax := len(grid), len(grid[0])
	for m := 0; m < mMax; m++ {
		for n := 0; n < nMax; n++ {
			if grid[m][n] == 0 {
				continue
			}
			islandPerimeter := islandPerimeterCheck(grid, m, n, mMax, nMax)
			perimeter += islandPerimeter
		}
	}
	return perimeter
}

func islandPerimeterCheck(grid [][]int, m, n, mMax, nMax int) int {
	perimeter := 0
	dirs := [4][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	for _, dir := range dirs {
		i := m + dir[0]
		j := n + dir[1]
		if i < 0 || j < 0 || i >= mMax || j >= nMax {
			perimeter++
			continue
		}
		if grid[i][j] == 0 {
			perimeter++
		}
	}
	return perimeter
}
