package golang

func findDiagonalOrder(mat [][]int) []int {
	yLen := len(mat)
	xLen := len(mat[0])
	var result []int
	reverseFlag := false
	startPoint := make([]int, 2)
	startPoint[0] = 0
	startPoint[1] = 0
	for {
		cr := getCurrentArray(mat, xLen, yLen, startPoint)
		if reverseFlag {
			for i, j := 0, len(cr)-1; i < j; i, j = i+1, j-1 {
				cr[i], cr[j] = cr[j], cr[i]
			}
		}
		result = append(result, cr...)
		if startPoint[0] == xLen-1 && startPoint[1] == yLen-1 {
			break
		}
		startPoint = getNextStratPoint(startPoint, xLen, yLen)
		reverseFlag = !reverseFlag
	}
	return result
}

func getCurrentArray(mat [][]int, xLen, yLen int, startPoint []int) []int {
	var cr []int
	xOffect, yOffect := 0, 0
	for {
		cr = append(cr, mat[startPoint[1]+yOffect][startPoint[0]+xOffect])
		xOffect++
		yOffect--
		if (startPoint[1]+yOffect) < 0 || (startPoint[0]+xOffect) > xLen-1 {
			break
		}
	}
	return cr

}

func getNextStratPoint(startPoint []int, xLen, yLen int) []int {
	if startPoint[0] == 0 {
		startPoint[1]++
		if startPoint[1] >= yLen {
			startPoint[0] = 1
			startPoint[1] = yLen - 1
		}
	} else {
		startPoint[0]++
	}
	return startPoint
}
