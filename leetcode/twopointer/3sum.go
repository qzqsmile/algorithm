package twopointer


func threeSum(nums []int) [][]int {
	res := [][]int{}
	for i := 0; i < len(nums); i++{
		tres := twoSum(nums[i+1:], -nums[i])
		for j := 0; j < len(tres); j++{
			t := append([]int{nums[i]}, tres[j]...)
			res = append(res, t)
		}
	}
	return res
}

func twoSum(nums []int, k int) [][]int{
	mp := make(map[int]int)
	res := [][]int{}
	for i := 0; i < len(nums); i++{
		mp[nums[i]] = i
	}
	for i := 0; i < len(nums); i++{
		if index, ok := mp[k-nums[i]]; ok && index != i && index > i{
			res = append(res, []int{nums[i], nums[mp[k-nums[i]]]})
		}
	}
	return res
}
