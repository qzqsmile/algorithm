package main

import "fmt"

func combine(n int, k int) [][]int {
	res := dfs(1, n+1, k)
	return res
}

func dfs(b int, e int, k int)[][] int{
	var res [] [] int
	if e-b < k || k <=0 {
		var tmp [][] int
		return tmp
	}
	if e - b == k{
		var tmp[][] int
		var tmp1[] int
		for i := b; i < e; i++{
			tmp1 = append(tmp1, i)
		}
		tmp = append(tmp, tmp1)
		return tmp
	}
	s1 := dfs(b+1, e, k)
	s2 := dfs(b+1, e, k-1)
	for i := 0; i < len(s2); i++{
		res = append(res, append(s2[i], b))
	}
	for i := 0; i < len(s1); i++{
		res = append(res, s1[i])
	}
	return res
}

func main()  {
	fmt.Println(combine(4, 2))
}



