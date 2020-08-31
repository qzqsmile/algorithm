package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// write your code here
	res := [][]int{}
	each := []int{}
	sort.Ints(candidates)
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
		if i != start && candidates[i] == candidates[i-1]{
			continue
		}
		each = append(each, candidates[i])
		helper(i+1, candidates, target, cur+candidates[i], each, res)
		each = each[0: len(each)-1]
	}
}


