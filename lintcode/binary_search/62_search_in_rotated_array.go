package main

/**
 * @param A: an integer rotated sorted array
 * @param target: an integer to be searched
 * @return: an integer
 */
func search (A []int, target int) int {
	// write your code here
	b, m, e := 0, 0, len(A)-1

	for ;b + 1 < e; {
		m = b + (e - b) / 2
		if A[m] == target{
			return m
		}else if A[m] < target {
			if A[m] > A[b] || A[e] >= target{
				b = m
			}else{
				e = m
			}
		}else{
			if A[m] <= A[b] || A[b] <= target{
				e = m
			}else{
				b = m
			}
		}
	}

	if A[b] == target{
		return b
	}
	if A[e] == target{
		return e
	}
	return -1
}

