package twopointer

// 三个指针，中间的指针是辅助的作用
func removeDuplicates(nums []int) int {
	c, s, l := 0, 0, 0
	for ;l < len(nums);{
		for;l < len(nums) && nums[l] == nums[s]; l++{}
		if s < len(nums){
			nums[c] = nums[s]
			c++
			s = l
		}
	}
	return c+1
}