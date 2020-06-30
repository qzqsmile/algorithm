package main

import "fmt"

func maximumGap(nums []int) int {
	nums = radixSort(nums)
	res := 0
	for i := 1; i < len(nums); i++{
		res = max(res, nums[i]- nums[i-1])
	}
	return res
}

func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}

func radixSort(nums [] int) [] int{
	var one, zero []int
	var i uint
	for i = 0; i < 32; i++ {
		needle := 1
		needle <<=  i
		one, zero= nil, nil
		for j := 0; j < len(nums); j++{
			if needle & nums[j] != 0{
				one = append(one, nums[j])
			}else{
				zero = append(zero, nums[j])
			}
		}
		nums = append(zero, one...)
	}
	return nums
}

func main() {
	nums :=  []int{3,6,9,1}
	fmt.Println(maximumGap(nums))
}