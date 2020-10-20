package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int] int)
	stk := []int{}
	for i := len(nums2)-1; i >= 0; i--{
		for len(stk) > 0 && stk[len(stk)-1] < nums2[i]{
			stk = stk[0: len(stk)-1]
		}
		if len(stk) == 0{
			mp[nums2[i]] = -1
		}else{
			mp[nums2[i]] = stk[len(stk)-1]
		}
		stk = append(stk, nums2[i])
	}

	res := []int{}
	for i := 0; i < len(nums1);  i++{
		res = append(res, mp[nums1[i]])
	}
	return res
}

