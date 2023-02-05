package main
//solutions:
// we need to adopt the nums as hashtable
// the concern case it index 0.
// the l may be stored in index 0, so we need to check it again

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++{
		j := i
		for nums[j] >= 0 && nums[j] != j &&  nums[j] < len(nums) && nums[nums[j]] != nums[j] {
			nums[j], nums[nums[j]] = nums[nums[j]], nums[j]
		}
	}
	l := 1
	for; l < len(nums); l++{
		if nums[l] != l{
			break
		}
	}
	if l == nums[0]{
		return l+1
	}
	return l
}