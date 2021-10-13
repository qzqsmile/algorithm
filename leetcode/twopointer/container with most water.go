package twopointer


func maxArea(height []int) int {
    b, e := 0, len(height)-1
    maxWater := -1
    for ;b <= e;{
        maxWater = max(maxWater, min(height[b], height[e])*(e-b))
        if height[b] < height[e]{
            b++
        }else{
            e--
        }
    }
    return maxWater
}

func max(a int, b int) int{
    if a  > b{
        return a
    }
    return b
}

func min(a int, b int) int{
    if a < b{
        return a
    }
    return b
}