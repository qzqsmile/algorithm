package main

func sortColors(nums []int)  {
	b := 0
	e := len(nums) - 1
	cur := 0
	for cur <= e{
		for nums[cur] == 2 && cur > b{
			tmp := nums[cur]
			nums[cur] = nums[b]
			nums[b] = tmp
			b += 1
		}
		for nums[cur] == 0 && cur < e{
			tmp := nums[cur]
			nums[cur] = nums[e]
			nums[e] = tmp
			e -= 1
		}
		cur += 1
	}
}

