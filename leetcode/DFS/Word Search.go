func exist(board [][]byte, word string) bool {
    flags := make([][]bool, len(board))
    for i := 0; i < len(board); i++{
        flags[i] = make([]bool, len(board[0]))
    }

    for i := 0; i < len(board); i++{
        for j := 0; j < len(board[0]); j++{
            contain := dfs(board, i, j, 0, word, flags)
            if contain{
                return true
            }
        }
    }
    return false
}


func dfs(board [][]byte, i int, j int, index int, word string,flags [][]bool) bool{
    if index >= len(word){
        return true
    }
    if i < 0 ||i >= len(board) || j < 0 || j >= len(board[0]) || word[index] != board[i][j] || flags[i][j] == true{
        return false
    }
    flags[i][j] = true
    isContain := dfs(board, i+1, j, index+1, word, flags) ||
    dfs(board, i-1, j, index+1, word, flags) ||
    dfs(board, i, j+1, index+1, word,flags) ||
    dfs(board, i, j-1, index+1, word, flags)
    flags[i][j] = false
    return isContain
}