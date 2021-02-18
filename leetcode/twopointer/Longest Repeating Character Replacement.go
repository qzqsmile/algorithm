package main

import (
	"math"
)

func characterReplacement(s string, k int) int {
	left, right, longest := 0, 0, 0
	win := make(map[byte]int)

	for right < len(s){
		c := s[right]
		right++
		win[c]++
		most := findMost(win)
		if right-left-win[most] <= k{
			longest = max(longest, right-left)
		}
		for right-left-win[most] > k{
			c := s[left]
			left++
			win[c]--
			most = findMost(win)
		}
	}
	return longest
}

func findMost(mp map[byte]int) byte{
	var most byte
	maxNum := math.MinInt64
	for k, v := range mp{
		if v > maxNum{
			most = k
			maxNum = v
		}
	}
	return most
}

func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}