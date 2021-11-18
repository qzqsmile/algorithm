package stack

func largestRectangleArea(heights []int) int {
	stk := []int{-1}
	heights = append(heights, -1)
	area := -1
	for i := 0; i < len(heights); i++{
		if len(stk) == 1 || heights[i] >= heights[stk[len(stk)-1]]{
			stk = append(stk, i)
		}else{
			for;len(stk) > 1 && heights[i] < heights[stk[len(stk)-1]];{
				t := stk[len(stk)-1]
				stk = stk[0: len(stk)-1]
				area = max(area, heights[t] * (i - stk[len(stk)-1]-1))
			}
			stk = append(stk, i)
		}
	}
	return area
}
