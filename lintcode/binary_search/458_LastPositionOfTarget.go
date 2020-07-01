package main


/**
 * @param nums: An integer array sorted in ascending order
 * @param target: An integer
 * @return: An integer
 */
func lastPosition (nums []int, target int) int {
	// write your code here
	if len(nums) == 0 {
		return -1
	}
	begin, mid, end := 0, 0, len(nums)-1
	for ; begin+1 < end; {
		mid = begin + (end-begin)/2
		if nums[mid] == target {
			begin = mid
		} else if nums[mid] < target {
			begin = mid
		} else {
			end = mid
		}
	}

	if nums[end] == target {
		return end
	}

	if nums[begin] == target {
		return begin
	}

	return -1
}

