package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	i, j := 0, len(nums)-1
	for ; i < j; {
		tmp := [][]int{}
		t := nums[i]
		twoSum(nums[i+1:], &tmp, -nums[i])
		for k := 0; k < len(tmp); k++ {
			res = append(res, append(tmp[k], nums[i]))
		}
		for ; i <= j && nums[i] == t; i++ {
		}
	}
	return res
}

func twoSum(nums []int, res *[][]int, target int) {
	if len(nums) < 2 {
		return
	}
	i, j := 0, len(nums)-1
	for ; i < j; {
		for ; i < j && nums[i]+nums[j] < target; i++ {
		}
		for ; i < j && nums[i]+nums[j] > target; j-- {
		}
		if i < j && nums[i]+nums[j] == target {
			*res = append(*res, []int{nums[i], nums[j]})
			t1, t2 := nums[i], nums[j]
			for ; i < j && t1 == nums[i]; i++ {
			}
			for ; i < j && t2 == nums[j]; j-- {
			}
		}
	}
}

