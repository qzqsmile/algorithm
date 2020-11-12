package main
//这题我觉的关键是要先利用一些if else把一些特殊的corner case给去掉

/**
 * @param nums: a mountain sequence which increase firstly and then decrease
 * @return: then mountain top
 */
func mountainSequence (nums []int) int {
	// write your code here
	if len(nums) == 1{
		return nums[0]
	}
	if len(nums) == 2{
		if nums[0] > nums[1]{
			return nums[0]
		}
		return nums[1]
	}

	b, e := 1, len(nums)-2
	for; b + 1 < e ;{
		m := (e-b)/2 + b
		if nums[m] > nums[m-1] && nums[m] > nums[m+1]{
			return nums[m]
		}else if nums[m] > nums[m-1]{
			b = m
		}else{
			e = m
		}
	}
	if nums[b] > nums[b-1] && nums[b] > nums[b+1]{
		return nums[b]
	}
	if nums[e] > nums[e-1] && nums[e] > nums[e+1]{
		return nums[e]
	}
	if nums[0] > nums[len(nums)-1]{
		return nums[0]
	}else{
		return nums[len(nums)-1]
	}
}
