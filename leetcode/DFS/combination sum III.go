package DFS

func combinationSum3(k int, n int) [][]int {
	candidate := []int{}
	for i := 1; i < 10; i++{
		candidate = append(candidate, i)
	}
	res := [][]int{}
	dfscomination(candidate, n, k, 0, []int{}, &res)
	return res
}

func dfscomination(candidate []int, target int, k int, index int, tmp []int, res *[][]int){
	if target < 0 || len(tmp) > k{
		return
	}
	if target == 0 && len(tmp) == k{
		*res = append(*res, append([]int{}, tmp...))
	}
	for i := index; i < len(candidate); i++{
		tmp = append(tmp, candidate[i])
		dfscomination(candidate, target-candidate[i], k, i+1, tmp, res)
		tmp = tmp[0:len(tmp)-1]
	}
}
