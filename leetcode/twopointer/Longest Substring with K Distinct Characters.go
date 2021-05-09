package twopointer

import (
	"math"
)

func LongestSubstringKDistinct(str string, k int) int{
	left, right := 0, 0
	win := make(map[byte]int)
	longest := math.MinInt64
	for right < len(str){
		c := str[right]
		win[c]++
		right++
		if len(win) <= k{
			longest = max(longest, right-left)
		}else{
			for len(win) > k{
				c := str[left]
				left++
				delete(win, c)
			}
		}
	}
	return longest
}
