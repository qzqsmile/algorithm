func missingNumber(nums []int) int {
    for i := 0; i < len(nums); i++{
        // tmp := nums[i]
        for;i != nums[i] && nums[i] < len(nums);{
            tmp := nums[i]
            nums[tmp], nums[i] = nums[i], nums[tmp]
        }
    }
    for i := 0; i < len(nums); i++{
        if i != nums[i]{
            return i
        }
    }

    return len(nums)
}