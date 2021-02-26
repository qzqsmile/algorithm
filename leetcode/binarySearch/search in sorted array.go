package binarySearch


//这题主要主要注意一点
//在写出二分模板后，只要加入一些判断条件，确认其是什么时候成立
//其它的都是其它情况，不用在详细些判断条件

func search1(nums []int, target int) int {
	b, e := 0, len(nums)-1
	for; b + 1 < e; {
		m := b + (e - b) /2
		if nums[m] < target {
			if nums[m] > nums[b] || target <= nums[e]{
				b = m // when this case is true
			}else{
				e = m
			}
		}else if nums[m] > target{
			if nums[m] < nums[e] || target >= nums[b]{
				e = m
			}else{
				b = m
			}
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


