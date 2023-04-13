package golang

func countBattleships(board [][]byte) int {
	boardCount := 0
	for m := 0; m < len(board); m++ {
		for n := 0; n < len(board[0]); n++ {
			fmt.Println(string(board[m][n]))
			if string(board[m][n]) == "X" {
				dfs(m, n, board)
				boardCount++
			}
		}
	}
	return boardCount
}

func dfs(m, n int, board [][]byte) {
	if m < 0 || n < 0 || m >= len(board) || n >= len(board[0]) {
		return
	}
	if string(board[m][n]) == "X" {
		board[m][n] = []byte(".")[0]
		dfs(m+1, n, board)
		dfs(m-1, n, board)
		dfs(m, n+1, board)
		dfs(m, n-1, board)
	}

}
