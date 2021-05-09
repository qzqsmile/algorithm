package twopointer

import (
	"math"
)

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	win := make(map[byte]int)
	for _, v := range t{
		need[byte(v)]++
	}
	left, right, match := 0, 0, 0
	start, end := 0, 0
	min := math.MaxInt64
	for right < len(s){
		c := s[right]
		right++
		if need[c] != 0{
			win[c]++
			if win[c] == need[c]{
				match++
			}
		}

		for match == len(need){
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c := s[left]
			left++
			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}

