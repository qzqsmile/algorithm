package main

import (
	"fmt"
	"sort"
)

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3{
		return 0
	}
	sort.Ints(nums)
	closesetSum := nums[0] + nums[1] + nums[2]
	for i := 0 ; i < len(nums)-2; i++{
		c := twoSum(nums[i+1:], target-nums[i])
		if abs(target-closesetSum) > abs(target-nums[i]-c){
			closesetSum = nums[i]+c
		}
	}
	return closesetSum
}

func twoSum(nums []int, target int) int{
	i, j := 0, len(nums)-1
	closetNum := nums[0] + nums[1]
	for ; i < j; {
		if i < j && nums[i]+nums[j] < target{
			if abs(target - (nums[i] + nums[j])) < abs(target - closetNum){
				closetNum = nums[i] + nums[j]
			}
			i++
		}
		if i < j && nums[i]+nums[j] > target{
			if abs(target - (nums[i] + nums[j])) < abs(target - closetNum){
				closetNum = nums[i] + nums[j]
			}
			j--
		}
		if i < j && nums[i]+nums[j] == target{
			return target
		}
	}
	return closetNum
}

func abs(a int) int{
	if a  < 0{
		return -a
	}
	return a
}