// O(n^2)
func maximumSwap(num int) int {
    nums := []int{}
    for;num > 0;{
        nums = append(nums, num % 10)
        num = num / 10
    }
    i, j := 0, len(nums)-1
    for;i < j;{
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }

    for i := 0; i < len(nums); i++{
        maxTmp := nums[i]
        maxIndex := i
        for j := i; j < len(nums); j++{
            if nums[j] >= maxTmp{
                maxTmp = nums[j]
                maxIndex = j
            }
        }
        if maxTmp != nums[i]{
            nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
            break
        }
    }

    res := 0
    for i := 0; i < len(nums); i++{
        res = 10 * res + nums[i]
    }
    return res
}
