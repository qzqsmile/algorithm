package main

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if (board[i][j] == 'O') {
				copy_board := mycopy(board)
				if (isSurrounded(copy_board, i, j)) {
					setRegions(board, i, j)
				}
			}
		}
	}
}

func setRegions(board [][]byte, i int, j int) {
	if i < 0 || i >= len(board) || j < 0 || j > len(board[0]) {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = 'X'
		setRegions(board, i+1, j)
		setRegions(board, i-1, j)
		setRegions(board, i, j+1)
		setRegions(board, i, j-1)
	}
}

func isSurrounded(board [][] byte, i int, j int) bool {
	if board[i][j] == 'O' {
		board[i][j] = 'X'
		if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1) {
			return false
		}
		up := isSurrounded(board, i+1, j)
		down := isSurrounded(board, i-1, j)
		right := isSurrounded(board, i, j+1)
		left := isSurrounded(board, i, j-1)
		return up && down && right && left
	}
	return true
}

func mycopy(board [][]byte) [][]byte {
	duplicate := make([][]byte, len(board))
	for i := range board {
		duplicate[i] = make([]byte, len(board[i]))
		copy(duplicate[i], board[i])
	}
	return duplicate
}

func main() {

}
