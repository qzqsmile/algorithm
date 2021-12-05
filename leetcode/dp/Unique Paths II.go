func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid) == 0{
        return 0
    }
    dp1 := make([]int, len(obstacleGrid[0])+1)
    for i := 0; i < len(obstacleGrid[0]); i++{
        if obstacleGrid[0][i] != 1{
            dp1[i+1] = 1
        }else{
            break
        }
    }
    for i := 1; i < len(obstacleGrid); i++{
        dp2 := []int{0}
        for j := 1; j < len(dp1); j++{
            if obstacleGrid[i][j-1] == 1{
                dp2 = append(dp2, 0)
            }else{
                dp2 = append(dp2, dp2[len(dp2)-1]+dp1[j])
            }
        }
        dp1 = dp2
    }

    return dp1[len(dp1)-1]
}