package binarySearch

func findPeakElement(nums []int) int {
	if len(nums) == 1 || nums[0] > nums[1]{
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2]{
		return len(nums)-1
	}
	b, e := 0, len(nums)-1
	for; b + 1 < e;{
		m := b + (e-b)/2
		if nums[m] > nums[m-1] && nums[m] > nums[m+1]{
			return m
		}else if nums[m] > nums[m-1]{
			b = m
		}else{
			e = m
		}
	}
	return -1
}

