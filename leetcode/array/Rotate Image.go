func rotate(matrix [][]int)  {
    i, j := 0, len(matrix)-1
    for ;i < j;{
        matrix[i], matrix[j] = matrix[j], matrix[i]
        i++
        j--
    }

    for i := 0; i < len(matrix); i++{
        for j := i; j < len(matrix[0]); j++{
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    // fmt.Println(matrix)
}

// 0,0 -> 0,2
// 0，1   1,2