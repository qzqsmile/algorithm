func moveZeroes(nums []int)  {
    c1, c2 := 0, 0
    for;c2 < len(nums);{
        if nums[c2] != 0{
            nums[c1], nums[c2] = nums[c2], nums[c1]
            c1++
        }
        c2++
    }
}