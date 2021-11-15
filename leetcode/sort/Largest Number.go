package sort

import "sort"

func largestNumber(nums []int) string {
    quickSort(nums, 0, len(nums)-1)
    res := ""

    if len(nums) > 0&& nums[len(nums)-1] == 0{
        return "0"
    }
    
    for i := len(nums)-1; i >= 0; i--{
        res += strconv.Itoa(nums[i])
    }
    return res
}

func quickSort(nums []int, b int, e int){
    if b >= e{
        return 
    }
    b1, e1 := b, e 
    for;b1 <= e1;{
        for;b1 <= e1 && (!firstBiggerThanSecond(nums[b1], nums[e])); b1++{}
        for;b1 <= e1 && firstBiggerThanSecond(nums[e1], nums[e]) ; e1--{}
        if b1 <= e1{
            nums[b1], nums[e1] = nums[e1], nums[b1]
        }
    }
    nums[b1], nums[e] = nums[e], nums[b1]
    quickSort(nums, b, b1-1)
    quickSort(nums, b1+1, e)
}

func firstBiggerThanSecond(a int, b int) bool{
    stra := strconv.Itoa(a)
    strb := strconv.Itoa(b)
    first := stra + strb
    second := strb + stra

    if first >= second {
        return true
    }
    return false
}


////////////////
func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        x, y := nums[i], nums[j]
        sx, sy := 10, 10
        for sx <= x {
            sx *= 10
        }
        for sy <= y {
            sy *= 10
        }
        return sy*x+y > sx*y+x
    })
    if nums[0] == 0 {
        return "0"
    }
    ans := []byte{}
    for _, x := range nums {
        ans = append(ans, strconv.Itoa(x)...)
    }
    return string(ans)
}
