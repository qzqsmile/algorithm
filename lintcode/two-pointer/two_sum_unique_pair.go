package main

/**
 * @param nums: an array of integer
 * @param target: An integer
 * @return: An integer
 */
import "sort"

func twoSum6 (nums []int, target int) int {
	// write your code here
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	count := 0
	for; i < j;{
		for ;i < j && nums[i] + nums[j] < target;i++{
		}
		for ;i < j && nums[i] + nums[j] > target;j--{
		}
		for ;i < j && nums[i] + nums[j] == target;{
			t1, t2 := nums[i], nums[j]
			for; i<j && nums[i] == t1; i++{
			}
			for; i<j && nums[j] == t2; j--{
			}
			count += 1
		}
	}
	return count
}

