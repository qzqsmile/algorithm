package main

import "fmt"

func search(nums []int, target int) int {
	var b, e int
	b, e = 0, len(nums)-1
	for b <= e {
		m := (b + e) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			if nums[m] >= nums[b] || nums[b] > target {
				b = m + 1
			} else {
				e = m - 1
			}
		} else {
			if nums[m] >= nums[b] && nums[b] > target {
				b = m + 1
			} else {
				e = m - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	res := search(nums, 0)
	fmt.Println(res)
}
