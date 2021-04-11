package divideConquer_Tree

import "strconv"

//这题使用divide&conquer很容易想到，但是有些细节的处理需要注意
//1. golang是否能进行string到int的转换使用strconv.Atoi返回error是否为null进行判断

func diffWaysToCompute(expression string) []int {
	if isDigit(expression){
		num, _ := strconv.Atoi(expression)
		return []int{num}
	}

	var res []int
	for index, tmpc := range(expression){
		tmp := string(tmpc)
		if tmp == "+" || tmp == "-" || tmp == "*" {
			left := diffWaysToCompute(expression[:index])
			right := diffWaysToCompute(expression[index+1:])
			for _, l := range left{
				for _, r := range right{
					if tmp == "+"{
						res = append(res, l + r)
					}else if tmp == "-"{
						res = append(res, l-r)
					}else if tmp == "/"{
						res = append(res, l/r)
					}else{
						res = append(res, l*r)
					}
				}
			}
		}
	}

	return res

}

func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}
