package twopointer

func sortArray(nums []int) []int {
	if len(nums) <= 1{
		return nums
	}
	k := nums[0]
	b, e := 0, len(nums)-1
	for; b < len(nums) && nums[b] <= k; b++{}
	for; e >= 0 && nums[e] >= k; e--{}
	for; b < e;{
		for;b < e && nums[b] <= k; b++{}
		for;b < e && nums[e] > k; e--{}
		if b < e{
			nums[b], nums[e] = nums[e], nums[b]
		}
	}
	sortArray(nums[0: b])
	sortArray(nums[b:])
	return nums
}

