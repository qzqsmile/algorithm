package binarySearch

//not binary search

func searchMatrix(matrix [][]int, target int) bool {
	r, c := 0, len(matrix[0])-1
	for; r < len(matrix) && c >= 0;{
		if matrix[r][c] < target{
			r += 1
		}else if matrix[r][c] > target {
			c -= 1
		}else{
			return true
		}
	}
	return false
}