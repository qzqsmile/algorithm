func intersection(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    res := []int{}
    i, j := 0, 0
    for;i < len(nums1) && j < len(nums2);{
        for;i < len(nums1) && j < len(nums2) && nums1[i] < nums2[j]; i++{}
        for;i < len(nums1) && j < len(nums2) && nums2[j] < nums1[i]; j++{}
        if i < len(nums1) && j < len(nums2) && nums1[i] == nums2[j]{
            res = append(res, nums1[i])
            val := nums1[i]
            for ;i < len(nums1) && nums1[i] == val; i++{}
        }
    }
    return res
}