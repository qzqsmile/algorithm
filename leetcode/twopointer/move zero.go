package twopointer

func moveZeroes(nums []int)  {
	b, c := 0, 0
	for;b < len(nums);{
		for;b < len(nums) && nums[b] == 0; b++{}
		if b < len(nums) && c < len(nums){
			nums[b], nums[c] = nums[c], nums[b]
			b++
			c++
		}
	}
}
