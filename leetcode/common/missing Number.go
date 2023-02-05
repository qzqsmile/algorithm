package main

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] < len(nums)-1 && nums[i] != i {
			tmp := nums[i]
			nums[i], nums[tmp] = nums[tmp], nums[i]
		}
	}
	// fmt.Printf("nums is %v", nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
