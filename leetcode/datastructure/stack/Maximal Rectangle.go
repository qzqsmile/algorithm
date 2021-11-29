func maximalRectangle(matrix [][]byte) int {
    aera := 0
    for r := len(matrix)-1; r >= 0; r--{
        height := []int{}
        calHeight(matrix, r, &height)
        aera = max(aera, largestRectangleArea(height))
    }
    return aera
}

func calHeight(matrix [][]byte, row int, height *[]int){
    
    for i := 0; i < len(matrix[0]); i++{
        h := 0
        for j := row; j >= 0; j--{
            if matrix[j][i] == '1'{
                h++
            }else{
                break
            }
        }
        *height = append(*height, h)
    }
}


func largestRectangleArea(heights []int) int {
    lToR, rToL := make([]int, len(heights)), make([]int, len(heights))
    stk := []int{}
    
    for i := 0; i < len(heights); i++{
        for;len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]];{
            stk = stk[0:len(stk)-1]     
        }
        if len(stk) == 0{
            lToR[i] = -1
        }else{
            lToR[i] = stk[len(stk)-1]
        }
        stk = append(stk, i)
    }
    stk = []int{}
    for i := len(heights)-1; i >= 0; i--{
        for;len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]];{
            stk = stk[0:len(stk)-1]     
        }
        if len(stk) == 0{
            rToL[i] = -1
        }else{
            rToL[i] = stk[len(stk)-1]
        }
        stk = append(stk, i)
    }
    aera := 0
    for i := 0; i < len(heights); i++{
        left, right := lToR[i], rToL[i]
        if right == -1{
            right = len(heights)
        }
        aera = max(aera, heights[i] * (right-left-1))
        // aera = max(aera, heights[i])
    }
    return aera
}


func max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}