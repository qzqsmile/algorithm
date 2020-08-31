package main

func permute(nums []int) [][]int {
	res := [][]int{}
	each := []int{}

	helper(0, nums, each, &res)
	return res
}


func helper(start int, nums []int,each[]int, res *[][]int){
	if 0 == len(nums){
		copyeach := append([]int(nil), each...)
		*res = append(*res, copyeach)
		return
	}
	for i := start; i < len(nums); i++{
		each = append(each, nums[i])
		cnums := append([]int{}, nums[start:i]...)
		cnums = append(cnums, nums[i+1:len(nums)]...)
		helper(0, cnums, each, res)
		each = each[0: len(each)-1]
	}
}