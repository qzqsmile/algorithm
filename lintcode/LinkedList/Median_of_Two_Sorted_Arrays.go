package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if (l1+l2)%2 == 1 {
		return float64(findhelper(nums1, 0, nums2, 0, (l1+l2+1)/2))
	} else {
		return float64((findhelper(nums1, 0, nums2, 0, (l1+l2)/2) +
			findhelper(nums1, 0, nums2, 0, (l1+l2)/2+1))) / 2
	}
}

func findhelper(nums1 []int, index1 int, nums2 []int, index2 int, k int) int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	if index1 >= len(nums1) {
		return nums2[index2+k-1]
	}

	if index2 >= len(nums2) {
		return nums1[index1+k-1]
	}
	if k == 1 {
		return min(nums1[index1], nums2[index2])
	}

	key1, key2 := MaxInt, MaxInt
	if index1+k/2-1 < len(nums1) {
		key1 = nums1[index1+k/2-1]
	}
	if index2+k/2-1 < len(nums2) {
		key2 = nums2[index2+k/2-1]
	}
	if key1 < key2 {
		return findhelper(nums1, index1+k/2, nums2, index2, k-k/2)
	} else {
		return findhelper(nums1, index1, nums2, index2+k/2, k-k/2)
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
