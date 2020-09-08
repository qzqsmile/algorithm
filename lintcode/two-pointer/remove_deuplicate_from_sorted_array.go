package two_pointer

func removeDuplicates(nums []int) int {
	i, j := 0, 1

	for ; j < len(nums);{
		for; j<len(nums) && nums[i] == nums[j];j++{
		}
		if j < len(nums){
			nums[i+1], nums[j] = nums[j], nums[i+1]
			i++
			j++
		}
	}
	return i+1
}
