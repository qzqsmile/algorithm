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



// another solution easy to understand
// func largestRectangleArea(heights []int) int {
//     lToR, rToL := make([]int, len(heights)), make([]int, len(heights))
//     stk := []int{}
    
//     for i := 0; i < len(heights); i++{
//         for;len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]];{
//             stk = stk[0:len(stk)-1]
//         }
        
//         if len(stk) == 0{
//             lToR[i] = -1
//         }else{
//             lToR[i] = stk[len(stk)-1]
//         }
//         stk = append(stk, i)
//     }
//     stk = []int{}
//     for i := len(heights)-1; i >= 0; i--{
//         for;len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]];{
//             stk = stk[0:len(stk)-1]
//         }
//         if len(stk) == 0{
//             rToL[i] = -1
//         }else{
//             rToL[i] = stk[len(stk)-1]
//         }
//         stk = append(stk, i)
//     }
//     aera := 0
//     for i := 0; i < len(heights); i++{
//         left, right := -1, len(heights)
//         if lToR[i] != -1{
//             left = lToR[i]
//         }
//         if rToL[i] != -1{
//             right = rToL[i]
//         }
//         aera = max(aera, (right-left-1)*heights[i])
//         aera = max(aera, heights[i])
//     }
//     return aera
// }


// func max(a int, b int) int{
//     if a > b{
//         return a
//     }
//     return b
// }