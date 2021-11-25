func spiralOrder(matrix [][]int) []int {
    res := []int{}
    if len(matrix) == 0{
        return res
    }
    helper(matrix, 0, 0, len(matrix)-1, len(matrix[0])-1, &res)
    return res
}

func helper(matrix [][]int, r1 int, c1 int,  r2 int, c2 int, res *[]int){
    if r2 < r1 || c2 < c1{
        return 
    }
    if r1 == r2{
        for i := c1; i <= c2; i++{
            *res = append(*res, matrix[r1][i])
        }
        return 
    }
    if c1 == c2{
        for i := r1; i <= r2; i++{
            *res = append(*res, matrix[i][c1])
        }
        return 
    }
    for i := c1; i < c2; i++{
        *res = append(*res, matrix[r1][i])
    }

    for i := r1; i < r2; i++{
        *res = append(*res, matrix[i][c2])
    }

    for i := c2; i > c1; i--{
        *res = append(*res, matrix[r2][i])
    }

    for i := r2; i > r1; i--{
        *res = append(*res, matrix[i][c1])
    }

    helper(matrix, r1+1, c1+1, r2-1, c2-1, res)
}