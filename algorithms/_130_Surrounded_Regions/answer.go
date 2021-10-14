package _130_Surrounded_Regions

func solve(board [][]byte) {
	// 边界O染色
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1) && board[i][j] == 'O' {
				O2Y(i, j, &board)
			}
		}
	}
	// 剩余O变X
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	// Y变O
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}

// dfs方式O变Y
func O2Y(i, j int, board *[][]byte) {
	if i < 0 || i > len(*board)-1 || j < 0 || j > len((*board)[i])-1 || (*board)[i][j] != 'O' {
		return
	}
	(*board)[i][j] = 'Y'
	O2Y(i-1, j, board)
	O2Y(i+1, j, board)
	O2Y(i, j-1, board)
	O2Y(i, j+1, board)
}
