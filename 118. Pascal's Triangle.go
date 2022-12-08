package golang

func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		curRow := make([]int, i+1)
		for j := range curRow {
			if j == 0 || j == i {
				curRow[j] = 1
				continue
			}
			curRow[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, curRow)
	}
	return res
}
