package binarySearch

func mysearch(nums []int, target int) int {
	b, e := 0, len(nums)-1
	for; b + 1 < e;{
		m := b + (e - b)/2
		if nums[m] < target{
			b = m
		}else if nums[m] > target{
			e = m
		}else{
			return m
		}
	}
	if nums[b] == target{
		return b
	}
	if nums[e] == target{
		return e
	}
	return -1
}

