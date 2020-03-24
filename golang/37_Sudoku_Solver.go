package main

func solveSudoku(board [][]byte) {
	slove(board)
}

func slove(board [][]byte) bool {
	row, col := findNextRowCol(board)
	if (row == -1 && col == -1) {
		return true
	}
	v := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < len(v); i++ {
		board[row][col] = v[i]
		if isSudoku(board, row, col) && slove(board) {
			return true
		}
		board[row][col] = '.'
	}
	return false
}

func isSudoku(board [][]byte, row int, col int) bool {
	if checkRow(board, row, board[row][col]) &&
		checkCol(board, col, board[row][col]) &&
		checkSqure(board, row, col, board[row][col]) {
		return true
	}
	return false
}

func findNextRowCol(board [][]byte) (int, int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func checkRow(board [][]byte, row int, v byte) bool {
	count := 0
	for i := 0; i < len(board); i++ {
		if board[row][i] == v {
			count++
		}
	}
	if count > 1 {
		return false
	}
	return true
}

func checkCol(board [][]byte, col int, v byte) bool {
	count := 0
	for i := 0; i < len(board); i++ {
		if board[i][col] == v {
			count++
		}
	}
	if count > 1 {
		return false
	}
	return true
}

func checkSqure(board [][]byte, row int, col int, v byte) bool {
	count := 0
	for i := (row / 3) * 3; i < (row/3+1)*3; i++ {
		for j := (col / 3) * 3; j < (col/3+1)*3; j++ {
			if board[i][j] == v {
				count++
			}
		}
	}
	if count > 1 {
		return false
	}
	return true
}
