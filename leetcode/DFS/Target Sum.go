func findTargetSumWays(nums []int, target int) int {
    if len(nums) == 0{
        return 0
    }
    if len(nums) == 1 {
        e1, e2 := 0, 0
        if target == nums[0]{
            e1 = 1
        }
        if target == -nums[0]{
            e2 = 1
        }
        return e1 + e2
    }
    n := findTargetSumWays(nums[1:], target+nums[0])
    p := findTargetSumWays(nums[1:], target-nums[0])
    return n+p
}
