func generateParenthesis(n int) []string {
    var tmp string
    res := []string{}
    dfs(0, 0, n, tmp, &res)
    return res
}

func dfs(left int, right int, n int, tmp string, res *[]string){
    if len(tmp) == 2*n{
        *res = append(*res, tmp)
        return 
    }
    if left == right{
        tmp += "("
        left++
        dfs(left, right, n, tmp, res)
        left--
        tmp = tmp[0:len(tmp)-1]
    }else{
        if left < n{
            tmp += "("
            left++
            dfs(left, right, n, tmp, res)
            left--
            tmp = tmp[0:len(tmp)-1]
        }
        if right < n{
            tmp += ")"
            right++
            dfs(left, right, n, tmp, res)
            right--
            tmp = tmp[0:len(tmp)-1]
        }
    }
} 