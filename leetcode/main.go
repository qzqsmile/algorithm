package main

import "fmt"

func main() {
	nums := []int{1,2,3}
	res := permute(nums)
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	res := [][]int{}
	tmp := []int{}
	dfs(nums, 0, tmp, &res)
	return res
}

func dfs(nums []int, index int, tmp []int, res *[][]int){
	if index > len(nums){
		return
	}
	if index == len(nums){
		*res = append(*res, append([]int{}, tmp...))
	}
	for i := index; i < len(nums); i++{
		nums[i], nums[index] = nums[index], nums[i]
		tmp = append(tmp, nums[index])
		dfs(nums, index+1, tmp, res)
		tmp = tmp[0:len(tmp)-1]
	}
}