package DFS

func combine(n int, k int) [][]int {
	res := [][]int{}
	tmp := []int{}
	dfs(1, n, k, tmp, &res)
	return res
}

func dfs(start int, end int, k int, tmp []int, res *[][]int){
	if len(tmp) == k{
		*res = append(*res, append([]int{}, tmp...))
	}
	if len(tmp) > k{
		return
	}

	for i := start; i <= end; i++{
		tmp = append(tmp, i)
		dfs(i+1, end, k, tmp, res)
		tmp = tmp[0:len(tmp)-1]
	}
}



