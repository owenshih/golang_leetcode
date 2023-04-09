package golang

import "strconv"

func updateBoard(board [][]byte, click []int) [][]byte {
	mLen := len(board)
	nLen := len(board[0])
	if string(board[click[0]][click[1]]) == "M" {
		board[click[0]][click[1]] = []byte("X")[0]
		return board
	}
	dfs(click[0], click[1], mLen, nLen, board)
	return board
}

func dfs(clickM, clickN, mLen, nLen int, grid [][]byte) {
	if clickM >= mLen || clickN >= nLen || clickM < 0 || clickN < 0 {
		return
	}
	if string(grid[clickM][clickN]) != "E" {
		return
	}
	bombCount := checkBomb(clickM, clickN, mLen, nLen, grid)
	switch bombCount {
	case 0:
		grid[clickM][clickN] = []byte("B")[0]
		direction := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		for _, k := range direction {
			if clickM+k[0] >= mLen || clickN+k[1] >= nLen || clickM+k[0] < 0 || clickN+k[1] < 0 {
				continue
			}
			dfs(clickM+k[0], clickN+k[1], mLen, nLen, grid)
		}
	default:
		grid[clickM][clickN] = strconv.Itoa(bombCount)[0]
	}
}

func checkBomb(clickM, clickN, mLen, nLen int, grid [][]byte) int {
	direction := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	bombCount := 0
	for _, v := range direction {
		if clickM+v[0] >= mLen || clickN+v[1] >= nLen || clickM+v[0] < 0 || clickN+v[1] < 0 {
			continue
		}
		if string(grid[clickM+v[0]][clickN+v[1]]) == "M" {
			bombCount++
		}
	}
	return bombCount
}
