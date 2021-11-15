func removeDuplicates(nums []int) int {
    c, s, f := 0, 0, 0
    for;f < len(nums);{
        nums[c] = nums[s]
        c++
        for;f < len(nums) && nums[s] == nums[f];f++{}
        s = f
    }
    return c
}