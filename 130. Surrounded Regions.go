package golang

func solve(board [][]byte) {
	nLen, mLen := len(board), len(board[0])
	for m := 0; m < mLen; m++ {
		if board[0][m] == 'O' {
			dfs(board, 0, m, 'A')
		}
		if board[nLen-1][m] == 'O' {
			dfs(board, nLen-1, m, 'A')
		}
	}
	for n := 0; n < nLen; n++ {
		if board[n][0] == 'O' {
			dfs(board, n, 0, 'A')
		}
		if board[n][mLen-1] == 'O' {
			dfs(board, n, mLen-1, 'A')
		}
	}
	for n := 0; n < nLen; n++ {
		for m := 0; m < mLen; m++ {
			if board[n][m] == 'O' {
				board[n][m] = 'X'
			}
			if board[n][m] == 'A' {
				board[n][m] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, n, m int, symbol byte) {
	if m < 0 || n < 0 || m >= len(board[0]) || n >= len(board) || board[n][m] == 'A' || board[n][m] == 'X' {
		return
	}
	board[n][m] = symbol
	dfs(board, n+1, m, symbol)
	dfs(board, n-1, m, symbol)
	dfs(board, n, m+1, symbol)
	dfs(board, n, m-1, symbol)
}
