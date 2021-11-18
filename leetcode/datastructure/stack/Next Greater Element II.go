package stack


func nextGreaterElements(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++{
		res = append(res, -1)
	}
	nums1 := append(append([] int{}, nums...), append([] int{}, nums...)...)
	stk := []int{}

	for i := 0; i < len(nums1); i++{
		for len(stk) > 0 && nums1[i] > nums1[stk[len(stk)-1]]{
			if stk[len(stk)-1] < len(res){
				res[stk[len(stk)-1]] = nums1[i]
			}
			stk = stk[0:len(stk)-1]
		}
		stk = append(stk, i)
	}

	return res
}


