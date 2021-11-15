func maxAreaOfIsland(grid [][]int) int {
    maxAera := 0
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            aera := 0
            dfs(grid, i, j, &aera)
            maxAera = max(maxAera, aera)
        }
    }
    return maxAera
}

func dfs(grid [][]int, i int, j int, aera *int){
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j]==0{
        return 
    }
    grid[i][j] = 0
    *aera += 1
    dfs(grid, i-1, j, aera)
    dfs(grid, i+1, j, aera)
    dfs(grid, i, j-1, aera)
    dfs(grid, i, j+1, aera)
}

func max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}
