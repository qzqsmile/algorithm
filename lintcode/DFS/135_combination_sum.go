package main

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	each := []int{}
	helper(0, candidates, target, 0, each, &res)
	return res
}

func helper(start int, candidates []int, target int, cur int, each []int, res *[][]int){
	if cur == target && len(each) > 0{
		copyeach := append([]int(nil), each...)
		*res = append(*res, copyeach)
		return ;
	}
	if cur > target{
		return ;
	}

	for i := start; i < len(candidates); i++{
		each = append(each, candidates[i])
		helper(i, candidates, target, cur+candidates[i], each, res)
		each = each[0: len(each)-1]
	}
}

