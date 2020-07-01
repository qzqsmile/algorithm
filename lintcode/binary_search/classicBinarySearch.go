package main

import "fmt"

/**
 * @param nums: An integer array sorted in ascending order
 * @param target: An integer
 * @return: An integer
 */
func findPosition (nums []int, target int) int {
	// write your code here
	if len(nums) == 0 {
		return -1
	}
	begin, end := 0, len(nums)-1
	mid := 0
	for;begin+1 < end;{
		mid = begin + (end-begin)/2
		if nums[mid] == target{
			end = mid
		}else if nums[mid] < target{
			begin = mid
		}else{
			end = mid
		}
	}

	if nums[begin] == target{
		return begin
	}
	if nums[end] == target{
		return end
	}
	return -1
}


func main()  {
	res := findPosition([]int{1,2,2,4,5,5}, 2)
	fmt.Println(res)
}
