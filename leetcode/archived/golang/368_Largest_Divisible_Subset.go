package main

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	S := make(map[int][]int)
	S[-1] = []int{}
	sort.Ints(nums)
	for _, value := range nums {
		max := []int{}
		for k, v := range S {
			if value%k == 0 && len(v) > len(max) {
				max = v
			}
		}
		max = append(max, value)
		S[value] = max
	}
	m := []int{}
	for _, value := range S {
		if len(value) > len(m) {
			m = value
		}
	}
	return m
}


