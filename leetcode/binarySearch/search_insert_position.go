package binarySearch

func searchInsert(nums []int, target int) int {
	if len(nums) == 0{
		return 0
	}
	if target > nums[len(nums)-1]{
		return len(nums)
	}

	b, e := 0, len(nums)-1
	for;b + 1 < e;{
		m := b + (e-b)/2
		if nums[m] <= target{
			b = m
		}else {
			e = m
		}
	}
	if nums[b] >= target{
		return b
	}
	return e
}
