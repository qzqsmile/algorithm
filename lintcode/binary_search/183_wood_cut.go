package main

/**
 * @param L: Given n pieces of wood with length L[i]
 * @param k: An integer
 * @return: The maximum length of the small pieces
 */
func woodCut(L []int, k int) int {
	// write your code here
	if len(L) == 0 {
		return 0
	}
	maxNum := L[0]
	sum := 0
	for _, v := range L {
		if v > maxNum {
			maxNum = v
		}
		sum += v
	}
	if sum < k {
		return 0
	}

	b, e := 1, maxNum

	for ; b+1 < e; {
		m := b + (e-b)/2
		c := woodNum(L, m)
		if c >= k {
			b = m
		} else {
			e = m
		}
	}
	if woodNum(L, e) >= k {
		return e
	}
	return b
}

func woodNum(L []int, m int) int {
	count := 0
	for _, v := range L {
		count += v / m
	}
	return count
}
