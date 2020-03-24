func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m1, m2 := len(nums1), len(nums2)
	if ((m1+m2)%2 == 0) {
		t1 := (m1 + m2) / 2
		t2 := t1 + 1
		r1 := findTarget(nums1, nums2, t1)
		r2 := findTarget(nums1, nums2, t2)
		return float64(r1+r2) / 2.0
	} else {
		// 2th
		t := (m1+m2)/2 + 1
		return float64(findTarget(nums1, nums2, t))
	}
}

func findTarget(nums1 []int, nums2 []int, t int) int {
	b1, e1 := 0, len(nums1)
	b2, e2 := 0, len(nums2)
	for ; (b1 < e1) && (b2 < e2); {
		m1 := (b1 + e1) / 2
		i2 := findIndex(nums2, nums1[m1], b2, e2)
		//to exmaine
		if i2+m1+1 > t {
			e1 = m1
			e2 = i2
		} else if i2+m1+1 < t-1 {
			b1 = m1 + 1
			b2 = i2
		} else {
			if i2+m1+1 == t-1 {
				if i2 == b2 {
					b1 = m1 + 1
				} else {
					if nums2[i2-1] < nums1[b1] {
						b2 = i2
					} else {
						b1 = m1 + 1
					}
				}
			} else {
				if i2 == 0 {
					return nums1[m1]
				}
				return Max(nums1[m1], nums2[i2-1])
			}
		}
	}

	for ; b1 < e1; {
		m1 := (b1 + e1) / 2
		if m1+e2+1 > t {
			e1 = m1
		} else if (m1+e2+1 < t) {
			b1 = m1 + 1
		} else {
			if e2 != 0 {
				return Max(nums1[m1], nums2[e2-1])
			}
			return nums1[m1]
		}
	}

	for ; b2 < e2; {
		m2 := (b2 + e2) / 2
		if m2+e1+1 > t {
			e2 = m2
		} else if (m2+e1+1 < t) {
			b2 = m2 + 1
		} else {
			if e1 != 0 {
				return Max(nums2[m2], nums1[e1-1])
			}
			return nums2[m2]
		}
	}

	return Max(nums2[e2-1], nums1[e1-1])
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findIndex(nums []int, t int, begin int, end int) int {
	b, e := begin, end
	for ; b < e; {
		m := (b + e) / 2
		if (nums[m] > t) {
			e = m
		} else if (nums[m] < t) {
			b = m + 1
		} else {
			for i := m; i < end; i++ {
				if (nums[i] != t) {
					return i
				}
			}
			return end
		}
	}
	return b
}
