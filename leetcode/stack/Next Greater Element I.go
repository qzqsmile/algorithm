package stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	stk := []int{}

	for i := 0; i < len(nums2); i++{
		if len(stk) == 0 || nums2[stk[len(stk)-1]] >= nums2[i]{
			stk = append(stk, i)
		}else{
			for;len(stk) > 0 && nums2[stk[len(stk)-1]] < nums2[i];{
				t := stk[len(stk)-1]
				stk = stk[0:len(stk)-1]
				mp[nums2[t]] = nums2[i]
			}
			stk = append(stk, i)
		}
	}

	var res []int

	for i := 0; i < len(nums1); i++{
		if val, ok := mp[nums1[i]]; ok{
			res = append(res, val)
		}else{
			res = append(res, -1)
		}
	}

	return res
}
