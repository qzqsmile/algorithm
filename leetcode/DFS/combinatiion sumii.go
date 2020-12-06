package DFS

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int {}
	each := []int{}
	sort.Ints(candidates)
	dfs(candidates, 0, target, each, &res)
	return res
}

//这里注意减枝的情形，可以将其想像成树状结构。dfs即是遍历根节点为 candidates[start]的所有组合
//为了不产生重复的结果，注意减枝。
func dfs(candidates []int, start int,  target int, each []int, res *[][]int ){
	if target == 0{
		copy := append([]int{}, each...)
		*res = append(*res, copy)
		return
	}
	for i := start; i < len(candidates); i++{
		if (i != start && candidates[i] == candidates[i - 1]) {
			continue;
		}
		if target < candidates[i]{
			return
		}
		each = append(each, candidates[i])
		dfs(candidates, i+1, target-candidates[i], each, res)
		each = each[0:len(each)-1]
	}
}
