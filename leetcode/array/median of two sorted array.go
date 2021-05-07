package array

//几点需要注意，1首先把找median转换为找第k个的问题
// 2. 通过设置k1, k2为math.Int64这种guard point来简化代码
// 3. k-k/2 != k/2因为有个除以舍去的问题
//
//
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)%2 == 0{
		r1 := findkth(nums1, nums2, (m+n)/2+1)
		r2 := findkth(nums1, nums2, (m+n)/2)
		return float64(r1+r2)/ 2
	}else{
		return float64(findkth(nums1, nums2, (m+n)/2+1))
	}
}

func findkth(nums1 []int, nums2 []int, k int) int{
	if len(nums1) == 0{
		return nums2[k-1]
	}
	if len(nums2) == 0{
		return nums1[k-1]
	}
	if k == 1{
		return min(nums1[0], nums2[0])
	}
	k1 := math.MaxInt32
	k2 := math.MaxInt32
	if len(nums1) > k/2-1{
		k1 = nums1[k/2-1]
	}
	if len(nums2) > k/2-1{
		k2 = nums2[k/2-1]
	}
	if k1 > k2{
		return findkth(nums1, nums2[k/2:], k-k/2)
	}else{
		return findkth(nums1[k/2:], nums2, k-k/2)
	}
}

func min(a int, b int) int{
	if a < b{
		return a
	}
	return b
}


