package main

func majorityElement(nums []int) int {
	res, count := nums[0], 1
	for _, v := range nums{
		if v == res{
			count++
		}else{
			count--
		}
		if count == 0{
			res = v
			count = 1
		}
	}
	return res
}
