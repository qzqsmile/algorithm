package DFS

func permute(nums []int) [][]int {
	res := [][]int{}
	tmp := []int{}
	// sort.Ints(nums)
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
	//here 这里copy一份是因为要保证子函数只是修改了自己的那部分，没有到调用者的那些变量
	//否则会造成程序编写的困难。
	copynum := append([]int{}, nums...)
	for i := index; i < len(copynum); i++{
		copynum[i], copynum[index] = copynum[index], copynum[i]
		tmp = append(tmp, copynum[index])
		dfs(copynum, index+1, tmp, res)
		tmp = tmp[0:len(tmp)-1]
	}
}
