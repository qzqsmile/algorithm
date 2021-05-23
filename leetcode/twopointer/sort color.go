package twopointer


// 3 step
// step1: 先搞好起始条件
// step2: 循环的条件设置一样，这样防止corner case出现
// step3: 对于这种3段指针，注意第三个指针的起始条件
func sortColors(nums []int)  {
	b, c, e := 0, 0,len(nums)-1

	for ;b < len(nums) && nums[b] == 0; b++{}
	for ;e >= 0 && nums[e] == 2; e--{}
	for c = b; c <= e && nums[c] == 1; c++{}

	for ;b < e && c <= e;{
		for;b < e && c <= e && nums[b] == 0; b++{}
		for;b < e && c <= e && nums[e] == 2; e--{}
		c = max(c, b)
		for;b < e && c <= e && nums[c] == 1; c++{}
		if b < e && c <= e {
			if nums[c] == 0{
				nums[b], nums[c] = nums[c], nums[b]
			}
			if nums[c] == 2{
				nums[e], nums[c] = nums[c], nums[e]
			}
		}
	}
}


