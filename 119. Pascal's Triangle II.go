package golang

func getRow(rowIndex int) []int {
	var result = make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				result[0] = 1
				continue
			}
			result[j] = result[j] + result[j-1]
		}
	}
	return result
}
