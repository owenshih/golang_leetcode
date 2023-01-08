package golang

func numIslands(grid [][]byte) int {
	numberOfIsland := 0
	mMax, nMax := len(grid),  len(grid[0])
	for m := 0; m < mMax; m++{
		for n := 0; n < nMax; n++{
			if(string(grid[m][n]) == "1"){
				numberOfIsland ++
				changeLandToSea(grid, m, n, mMax, nMax)
			}
		}
	}
	return numberOfIsland
}

func changeLandToSea(grid [][]byte, m, n, mMax, nMax int){
	if(m < 0 || n < 0 || m >= mMax || n >= nMax || string(grid[m][n]) == "0"){
		return
	}
	grid[m][n] = 48\
	changeLandToSea(grid, m + 1, n, mMax, nMax);
	changeLandToSea(grid, m - 1, n, mMax, nMax);
	changeLandToSea(grid, m, n + 1, mMax, nMax);
	changeLandToSea(grid, m, n - 1, mMax, nMax);
}