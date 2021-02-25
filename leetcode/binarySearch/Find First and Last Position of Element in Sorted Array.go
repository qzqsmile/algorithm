package binarySearch

func searchRange(nums []int, target int) []int {
	if len(nums) == 0{
		return []int {-1, -1}
	}
	b, e := fristPosition(nums, target), lastPosition(nums, target)
	return []int{b, e}
}

func fristPosition(nums []int, target int) int{
	b, e := 0, len(nums)-1
	for; b+1<e;{
		m := b + (e-b)/2
		if nums[m] < target{
			b = m
		}else if nums[m] > target{
			e = m
		}else{
			e = m
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

func lastPosition(nums []int, target int) int{
	b, e := 0, len(nums)-1
	for; b+1<e;{
		m := b + (e-b)/2
		if nums[m] < target{
			b = m
		}else if nums[m] > target{
			e = m
		}else{
			b = m
		}
	}

	if nums[e] == target{
		return e
	}
	if nums[b] == target{
		return b
	}
	return -1
}
