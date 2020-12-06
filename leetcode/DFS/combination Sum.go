package DFS


func combinationSum(candidates []int, target int) [][]int {
	res := [][]int {}
	each := []int{}
	sort.Ints(candidates)
	dfs(candidates, 0, target, each, &res)
	return res
}

func dfs(candidates []int, start int,  target int, each []int, res *[][]int ){
	if target == 0{
		copy := append([]int{}, each...)
		*res = append(*res, copy)
		return
	}
	for i := start; i < len(candidates); i++{
		if target < candidates[i]{
			return
		}
		each = append(each, candidates[i])
		dfs(candidates, i, target-candidates[i], each, res)
		each = each[0:len(each)-1]
	}
}
