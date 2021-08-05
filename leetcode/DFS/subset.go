package DFS

func subsets(nums []int) [][]int {
	var res [][]int
	var tmp []int
	res = append(res, []int{})
	subsetdfs(nums, 0, tmp, &res)
	return res
}

func subsetdfs(nums []int, index int, tmp []int, res *[][]int){
	if index >= len(nums){
		return
	}
	for i := index; i < len(nums); i++{
		tmp = append(tmp, nums[i])
		*res = append(*res, append([]int{}, tmp...))
		subsetdfs(nums, i+1, tmp, res)
		tmp = tmp[0:len(tmp)-1]
	}
}
