package main

import (
	"fmt"
	"strconv"
)

func addOperators(num string, target int) []string {
	var res [] string
	if len(num) == 0 && target == 0{
		return res
	}
	dfs1(num, 0, 0, 0, target, "", &res)
	return res
}

func dfs1(num string, cur int, p int, pre int, target int, path string, res *[]string){
	if p == len(num) && cur == target{
		*res = append(*res, path)
		return
	}
	for i := p; i < len(num); i++{
		if num[p] == '0' && i != p{
			break
		}
		tmp, _ := strconv.Atoi(num[p:i+1])
		if p == 0 {
			dfs1(num, cur+tmp, i+1, pre+tmp, target, path+strconv.Itoa(tmp), res)
		} else{
			dfs1(num, cur+tmp, i+1, tmp, target, path+"+"+strconv.Itoa(tmp), res)
			dfs1(num, cur-tmp, i+1, -tmp, target, path+"-"+strconv.Itoa(tmp), res)
			dfs1(num, cur-pre+pre*tmp, i+1, pre*tmp, target, path+"*"+strconv.Itoa(tmp), res)
		}
	}
}

func main()  {
	res := addOperators("105", 5)
	fmt.Println(res)
}




