func nextPermutation(nums []int)  {
    lowestj, lowesti := -1, -1
    for i := len(nums)-1; i >= 0; i--{
        for j := i; j >= 0; j--{
            if nums[j] < nums[i] && j > lowestj{
                lowestj = j 
                lowesti = i
            }
        }
    }
    if lowestj != -1{
        nums[lowesti], nums[lowestj] = nums[lowestj], nums[lowesti]
        sort.Ints(nums[lowestj+1:])
        return 
    }
    reverse(nums)
    return 
}