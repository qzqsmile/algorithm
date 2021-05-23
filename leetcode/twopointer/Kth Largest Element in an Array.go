package twopointer


func findKthLargest(nums []int, k int) int {
	return findKSmallest(nums, len(nums)-k+1)
}


func findKSmallest(nums []int, k int)int{
	b, e := 0, len(nums)-1
	slot := nums[0]
	for ; b < len(nums) && nums[b] < slot; b++{}
	for ; e >= 0 && nums[e] >= slot; e-- {}
	for ;b < e; {
		for ;b < e && nums[b] < slot; b++{}
		for ;b < e && nums[e] >= slot; e--{}
		if b < e {
			nums[b], nums[e] = nums[e], nums[b]
		}
	}
	if b == k-1{
		return slot
	}else if b >= k{
		return findKSmallest(nums[0:b], k)
	}else{
		if b != 0{
			return findKSmallest(nums[b:], k-b)
		}else{
			return findKSmallest(nums[1:], k-b-1)
		}
	}
}

