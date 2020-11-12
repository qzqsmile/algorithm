package datastructure
//stack

func dailyTemperatures(T []int) []int {
	res := []int{}
	if len(T) == 0{
		return res
	}

	stk := []int{}
	stk = append(stk, 0)
	res = append(res, 0)
	for i := 1; i < len(T); i++{
		res = append(res, 0)
		for ;len(stk)>=1 && T[i] > T[stk[len(stk)-1]];{
			res[stk[len(stk)-1]] = i-stk[len(stk)-1]
			stk = stk[0:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return res
}
