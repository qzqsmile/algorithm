func partition(s string) [][]string {
	res := [][]string{}
	each := []string{}
	helper(0, s, each, &res)
	return res
}

func helper(start int, s string, each []string, res *[][]string){
	if start == len(s){
		copyeach := append([]string(nil), each...)
		*res = append(*res, copyeach)
	}

	for i := start; i < len(s); i++{
		if isPalindrome(s[start:i+1]){
			each = append(each, s[start:i+1])
			helper(i+1, s, each, res)
			each = each[0:len(each)-1]
		}
	}
}

func isPalindrome(s string) bool{
	if len(s) == 0{
		return true
	}
	i, j := 0, len(s)-1

	for; i < j;{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}