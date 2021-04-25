package array

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j:= m-1, n-1
	cur := len(nums1)-1
	for; i >= 0 && j >= 0;{
		if nums1[i] > nums2[j]{
			nums1[cur] = nums1[i]
			i--
		}else{
			nums1[cur] = nums2[j]
			j--
		}
		cur--
	}
	for;j>=0;j--{
		nums1[cur] = nums2[j]
		cur--
	}
}