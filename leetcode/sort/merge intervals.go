func merge(intervals [][]int) [][]int {
    if len(intervals) == 0{
        return [][]int{}
    }
    mergesort(intervals, 0, len(intervals)-1)
    res := [][]int{intervals[0]}

    for i := 1; i < len(intervals); i++{
        if intervals[i][0] <= res[len(res)-1][1]{
            res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1]) 
        }else{
            res = append(res, intervals[i])
        }
    }

    return res
}

func max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}


// func mergesort(intervals [][]int, b int, e int){
//     if b >= e{
//         return 
//     }
//     m := b + (e-b) / 2
//     mergesort(intervals, b, m)
//     mergesort(intervals, m+1, e)
//     tmp := [][]int{}
//     b1, m1 := b, m+1
//     for;b1 <= m && m1 <= e;{
//         if intervals[b1][0] < intervals[m1][0]{
//             tmp = append(tmp, intervals[b1])
//             b1++
//         }else{
//             tmp = append(tmp, intervals[m1])
//             m1++
//         }
//     }
//     for;b1 <= m; b1++{
//         tmp = append(tmp, intervals[b1])
//     }
//     for;m1 <= e;m1++{
//         tmp = append(tmp, intervals[m1])
//     }
//     for i := 0; i < len(tmp); i++{
//         intervals[b+i] = tmp[i]
//     }
// }

// func quicksort(intervals [][]int, b int, e int){
//     if b >= e{
//         return 
//     }
//     m := intervals[e][0]
//     b1, e1 := b, e
//     for;b1 <= e1;{
//         for;b1 <= e1 && intervals[b1][0] < m; b1++{}
//         for;b1 <= e1 && intervals[e1][0] >= m; e1--{}
//         if b1 <= e1{
//             intervals[b1], intervals[e1] = intervals[e1], intervals[b1]
//             // b1++
//             // e1--
//         }
//     }
//     intervals[b1], intervals[e] = intervals[e], intervals[b1]
//     quicksort(intervals, b, b1-1)
//     quicksort(intervals, b1+1, e)
// }