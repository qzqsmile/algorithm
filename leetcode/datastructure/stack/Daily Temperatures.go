func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    stk := []int{}

    for i := len(temperatures)-1; i >= 0; i--{
        for;len(stk) > 0 && temperatures[i] >= temperatures[stk[len(stk)-1]];{
            stk = stk[0:len(stk)-1]
        }
        if len(stk) == 0{
            res[i] = 0
        }else{
            res[i] = stk[len(stk)-1] - i
        }
        stk = append(stk, i)
    }

    return res
}