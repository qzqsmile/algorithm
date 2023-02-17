package dp

import "sort"

//this problem can also be solved by dp
//but the complexity will case tl issue, dp n^2
//the key point of this binary search is the last element
//of f means that there is len(f) continuous numbers

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	f := []int{}
	for _, e := range envelopes {
		h := e[1]
		if i := findGreaterOrEuqalIndex(f, h); i != -1 {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}

func findGreaterOrEuqalIndex(nums []int, slot int) int {
	if len(nums) == 0 {
		return -1
	}
	b, e := 0, len(nums)-1
	for b+1 < e {
		m := b + (e-b)/2
		if nums[m] <= slot {
			b = m
		} else {
			e = m
		}
	}
	if nums[b] >= slot {
		return b
	}
	if nums[e] >= slot {
		return e
	}
	return -1
}
