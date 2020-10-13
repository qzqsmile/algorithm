package main

import "fmt"

/**
 * @param nums: A set of numbers
 * @return: A list of lists
 */
func subsets(nums []int) [][]int {
	res := [][]int{}
	each := []int{}
	res = append(res, each)
	helper(0, nums, each, &res)
	return res
}

func helper(start int, nums []int, each []int,res *[][]int) {
	for i := start; i < len(nums); i++{
		each = append(each, nums[i])
		copyeach := append([]int{}, each...)
		*res = append(*res, copyeach)
		helper(i+1, nums, each, res)
		each = each[0: len(each)-1]
	}
}