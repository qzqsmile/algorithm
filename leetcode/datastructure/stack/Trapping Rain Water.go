func trap(height []int) int {
    stk := []int{}
    leftRight := []int{}
    rightLeft := []int{}
    for i := 0; i < len(height); i++{
        for;len(stk) > 0 && height[i] > height[stk[len(stk)-1]];{
            stk = stk[0:len(stk)-1]
        }
        if len(stk) == 0{
            leftRight = append(leftRight, -1)
        }else{
            leftRight = append(leftRight, stk[len(stk)-1])
        }
        stk = append(stk, i)
    }
    stk = []int{}
    for i := len(height)-1; i >= 0; i--{
        for;len(stk) > 0 && height[i] >= height[stk[len(stk)-1]];{
            stk = stk[0:len(stk)-1]
        }
        if len(stk) == 0{
            rightLeft = append(rightLeft, -1)
        }else{
            rightLeft = append(rightLeft, stk[len(stk)-1])
        }
        stk = append(stk, i)
    }
    reverse(rightLeft)
    ans := 0
    for i := 0; i < len(leftRight); i++{
        if leftRight[i] != - 1 && rightLeft[i] != -1{
            fmt.Println(leftRight[i], rightLeft[i])
            curWidth := rightLeft[i] - leftRight[i]-1
            curHeight := min(height[rightLeft[i]], height[leftRight[i]]) - height[i]
            ans += curHeight * curWidth
        }
    }

    return ans
}

func min(a int, b int) int{
    if a < b{
        return a
    }
    return b
}
func reverse(nums []int){
    i, j := 0, len(nums)-1
    for;i < j;{
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}