package main

/**
 * @param nums: The integer array.
 * @param target: Target to find.
 * @return: The first position of target. Position starts from 0.
 */
func binarySearch(nums []int, target int) int {
	// write your code here
	if len(nums) == 0 {
		return -1
	}
	begin, mid, end := 0, 0, len(nums)-1
	for ; begin+1 < end; {
		mid = begin + (end-begin)/2
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			begin = mid
		} else {
			end = mid
		}
	}

	if nums[begin] == target {
		return begin
	}
	if nums[end] == target {
		return end
	}
	return -1
}
