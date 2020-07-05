package main
/**
 * @param matrix: matrix, a list of lists of integers
 * @param target: An integer
 * @return: a boolean, indicate whether matrix contains target
 */
func searchMatrix(matrix [][]int, target int) bool {
	// write your code here
	if len(matrix) == 0 || target < matrix[0][0] ||
		target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}

	b, m, e := 0, 0, len(matrix)-1
	for ; b+1 < e; {
		m = b + (e-b)/2
		if matrix[m][0] == target {
			return true
		} else if matrix[m][0] > target {
			e = m
		} else {
			b = m
		}
	}

	t := b
	if matrix[e][0] <= target {
		t = e
	}

	b, m, e = 0, 0, len(matrix[0])-1
	for ; b+1 < e; {
		m = b + (e-b)/2
		if matrix[t][m] == target {
			return true
		} else if matrix[t][m] > target {
			e = m
		} else {
			b = m
		}
	}
	if matrix[t][b] == target {
		return true
	}
	if matrix[t][e] == target {
		return true
	}

	return false
}

