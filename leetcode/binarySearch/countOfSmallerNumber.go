package binarySearch

import (
	"sort"
)
/**
 * @param A: An integer array
 * @param queries: The query list
 * @return: The number of element in the array that are smaller that the given integer
 */

func countOfSmallerNumber (A []int, queries []int) []int {
	// write your code here
	sort.Ints(A)
	res := []int{}
	for i := 0; i  < len(queries); i++{
		v := findposition(A, queries[i])
		res = append(res, v)
	}
	return res
}

func findposition(A []int, target int) int{
	b, e := 0, len(A)-1
	for; b + 1 < e;{
		m := b + (e - b)/2
		if A[m] >= target{
			e = m
		}else{
			b = m
		}
	}
	if A[b] >= target{
		return b
	}
	return e
}