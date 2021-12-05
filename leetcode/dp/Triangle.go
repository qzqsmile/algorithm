func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0{
        return 0
    }
    dp1 := []int{triangle[0][0]}
    for i := 1; i < len(triangle); i++{
        dp2 := []int{}
        for j := 0; j < len(triangle[i]); j++{
            tmp := math.MaxInt64
            if j < len(dp1){
                tmp = triangle[i][j] + dp1[j]
            }
            if j-1 >= 0{
                tmp = min(tmp, triangle[i][j]+dp1[j-1])
            }
            dp2 = append(dp2, tmp)
        }
        // fmt.Println(dp2)
        dp1 = dp2
    }
    ans := math.MaxInt64
    for i := 0; i < len(dp1); i++{
        ans = min(ans, dp1[i])
    }
    return ans
}

func min(a int, b int) int{
    if a < b{
        return a
    }
    return b
}