package binarySearch

func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0{
		return false
	}
	row := findrow(matrix, target)
	b, e := 0, len(matrix[0])-1
	for;b+1 < e;{
		m := b + (e-b)/2
		if matrix[row][m] < target{
			b = m
		}else if matrix[row][m] > target{
			e = m
		}else{
			return true
		}
	}
	if matrix[row][b] == target{
		return true
	}
	if matrix[row][e] == target{
		return true
	}
	return false
}

func findrow(matrix [][]int, target int) int{
	b, e := 0, len(matrix)-1
	for;b+1<e;{
		m := b + (e - b)/2
		if matrix[m][0] <= target{
			b = m
		}else{
			e = m
		}
	}
	if matrix[e][0] <= target{
		return e
	}
	return b
}
