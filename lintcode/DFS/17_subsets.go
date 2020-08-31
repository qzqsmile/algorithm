package main

/**
 * @param nums: A set of numbers
 * @return: A list of lists
 */
func subsets(nums []int) [][]int {
	res := [][]int{}
	each := []int{}
	helper(0, nums, each, &res)
	return res
}

func helper(start int, nums []int, each []int,res *[][]int) {
	if start > len(nums){
		return
	}
	copyeach := append([]int(nil), each...)
	*res = append(*res, copyeach)
	for i := start; i < len(nums); i++{
		each = append(each, nums[i])
		helper(i+1, nums, each, res)
		each = each[0:len(each)-1]
	}
}
