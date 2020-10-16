package main

func rotate(matrix [][]int) {
	helper(matrix, 0, len(matrix)-1)
}

func helper(matrix [][]int, r int, c int) {
	if r >= c {
		return
	}
	for i := r; i <= c-1; i++ {
		matrix[r][i], matrix[i][c], matrix[c][c-(i-r)], matrix[c-(i-r)][r] =
			matrix[c-(i-r)][r], matrix[r][i], matrix[i][c], matrix[c][c-(i-r)]
	}
	helper(matrix, r+1, c-1)
}

