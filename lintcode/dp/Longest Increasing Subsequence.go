package main

func lengthOfLIS(nums []int) int {
	// write your code here
	if len(nums) == 0{
		return 0
	}
	fn := []int{}
	for i := 0; i < len(nums); i++{
		fn = append(fn, 1)
	}

	for i := 0; i < len(nums); i++{
		for j := 0; j < i; j++{
			if nums[j] < nums[i]{
				fn[i] = max(fn[i], fn[j]+1)
			}
		}
	}

	maxCount := fn[0]
	for _, v := range fn{
		maxCount = max(maxCount, v)
	}
	return maxCount
}

func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}


