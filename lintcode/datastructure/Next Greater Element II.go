package main

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0{
		return []int{}
	}
	stk, res := []int{}, []int{}
	maxNum, maxIndex := nums[0], 0

	for i := 0; i < len(nums); i++{
		if nums[i] > maxNum{
			maxIndex = i
			maxNum = nums[i]
		}
	}
	part1 := []int{}
	for i := maxIndex; i >= 0; i--{
		for len(stk) > 0 && stk[len(stk)-1] <= nums[i]{
			stk = stk[0: len(stk)-1]
		}
		if len(stk) == 0{
			part1 = append(part1, -1)
		}else{
			part1 = append(part1, stk[len(stk)-1])
		}
		stk = append(stk, nums[i])
	}

	part2 := []int{}
	for i := len(nums)-1; i > maxIndex; i--{
		for len(stk) > 0 && stk[len(stk)-1] <= nums[i]{
			stk = stk[0: len(stk)-1]
		}
		if len(stk) == 0{
			part2 = append(part2, -1)
		}else{
			part2 = append(part2, stk[len(stk)-1])
		}
		stk = append(stk, nums[i])
	}
	reverse(part1)
	reverse(part2)
	res = append(part1, part2...)
	return res
}

func reverse(nums []int){
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1{
		nums[i], nums[j] = nums[j], nums[i]
	}
}
