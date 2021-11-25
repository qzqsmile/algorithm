
func findMaxLength(nums []int) int {
    sum := 0
    mp := make(map[int]int)
    maxLength := 0
    mp[0] = -1
    for i := 0; i < len(nums); i++{
        if nums[i] == 0{
            sum -= 1
        }else{
            sum += 1
        }
        if _, ok := mp[sum]; !ok{
            mp[sum] = i
        }else{
            maxLength = max(maxLength, i-mp[sum])
        }
    }
    return maxLength
}

func max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}