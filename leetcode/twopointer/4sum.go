package twopointer

import "sort"

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    tmp := []int{}
    for i := 0; i < len(nums); i++{
        if i == 0 || nums[i] != nums[i-1]{
            tmp = append(tmp, nums[i])
            threeSum(nums[i+1:], target-nums[i], tmp ,&res)
            tmp = tmp[0:len(tmp)-1]
        }
    }
    return res
}

func threeSum(nums []int, target int, tmp []int, res *[][]int) {
    for i := 0; i < len(nums); i++{
        if i == 0 || nums[i] != nums[i-1]{
            tmp = append(tmp, nums[i])
            twoSum(nums[i+1:], target-nums[i], tmp, res)
            tmp = tmp[0:len(tmp)-1]
        }
    }
}

// -4 -1 -1 0 1 2
 
func twoSum(nums []int, target int, tmp []int, res *[][]int){
    b, e := 0, len(nums)-1
    for;b < e;{
        for;b < e && nums[b] + nums[e] < target; b++{}
        for;b < e && nums[b] + nums[e] > target; e--{}
        if b < e && nums[b]+nums[e] == target{
            *res = append(*res, append(tmp, []int{nums[b], nums[e]}...))
            b1 := nums[b]
            for ;b < e && nums[b] == b1; b++{}
        }
    }
}