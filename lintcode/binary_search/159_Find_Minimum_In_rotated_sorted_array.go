package main

/**
 * @param nums: a rotated sorted array
 * @return: the minimum number in the array
 */
func findMin (nums []int) int {
	// write your code here
	b, m, e := 0, 0, len(nums)-1
	t := nums[len(nums)-1]

	for ;b + 1 < e;{
		m = b + (e - b) / 2
		if nums[m] <= t{
			e = m
		}else{
			b = m
		}
	}
	if nums[b] < nums[e]{
		return nums[b]
	}
	return nums[e]
}

