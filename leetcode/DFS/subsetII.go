package DFS

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{[]int{}}
	var tmp []int
	dfs(nums, 0, tmp, &res)
	return res
}

func dfs(nums []int, index int, tmp []int, res *[][]int){
	if index > len(nums){
		return
	}
	for i := index; i < len(nums); i++{
		if i != index && nums[i] == nums[i-1]{
			continue
		}
		tmp = append(tmp, nums[i])
		*res = append(*res, append([]int{}, tmp...))
		dfs(nums, i+1, tmp, res)
		tmp = tmp[0:len(tmp)-1]
	}
}