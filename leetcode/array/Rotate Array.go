func rotate(nums []int, k int)  {
    k = k % len(nums)
    rotateRange(nums, 0, len(nums)-1)
    rotateRange(nums, k, len(nums)-1)
    rotateRange(nums, 0, k-1)
}

func rotateRange(nums []int, b int, e int){
    b1, e1 := b, e
    for;b1 <= e1;{
        nums[b1], nums[e1] = nums[e1], nums[b1]
        b1++
        e1--
    }
}
