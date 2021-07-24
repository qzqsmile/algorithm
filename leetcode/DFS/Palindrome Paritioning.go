package DFS

func partition(s string) [][]string {
	res := [][]string{}
	tmp := []string{}
	dfs(s, 0, tmp, &res)
	return res
}

func mydfs(s string, index int, tmp []string, res *[][]string){
	if index > len(s){
		return
	}
	if index == len(s){
		*res = append(*res, append([]string{}, tmp...))
	}
	for i := index; i < len(s); i++{
		if isPalindrome(s[index:i+1]){
			tmp = append(tmp, s[index: i+1])
			dfs(s, i+1, tmp, res)
			tmp = tmp[0: len(tmp)-1]
		}
	}
}

func isPalindrome(s string) bool{
	i, j := 0, len(s)-1
	for ;i < j;{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}


