package main

import "sort"

func nextPermutation(nums []int)  {
	if len(nums) == 0{
		return
	}
	max, max_i := nums[len(nums)-1], len(nums)-1

	for i := len(nums)-2; i >= 0; i--{
		if nums[i] < max{
			for j := i+1; j < len(nums); j++{
				if nums[j] > nums[i] && nums[j] <= max{
					max = nums[j]
					max_i = j
				}
			}
			nums[i], nums[max_i] = nums[max_i], nums[i]
			reverse(nums[i+1:])
			return
		}else{
			max = nums[i]
			max_i = i
		}
	}

	reverse(nums)
}

func reverse(nums []int){
	for i, j := 0, len(nums)-1; i < j;{
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func permute(nums []int) [][]int {
	n := len(nums)
	r := 1
	for n > 0{
		r = r * n
		n--
	}
	r--
	res := [][]int{}
	copynum := append([]int{}, nums...)
	res = append(res, copynum)

	for ;r > 0; r--{
		nextPermutation(nums)
		copynum = append([]int{}, nums...)
		res = append(res, copynum)
	}
	return res
}




