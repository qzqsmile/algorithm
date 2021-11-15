func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j := 0, 0
    sorted := []int{}
    for;i < m && j < n;{
        if nums1[i] < nums2[j]{
            sorted = append(sorted, nums1[i])
            i++
        }else{
            sorted = append(sorted, nums2[j])
            j++
        }
    }
    if i < m{
        sorted = append(sorted, nums1[i:m]...)
    }
    if j < n{
        sorted = append(sorted, nums2[j:]...)
    }
    copy(nums1, sorted)
}

//another solution

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j, c := m-1, n-1, len(nums1)-1
    for;i >= 0 && j >= 0;{
        if nums1[i] > nums2[j]{
            nums1[c] = nums1[i]
            i--
        }else{
            nums1[c] = nums2[j]
            j--
        }
        c--
    }
    for;j >= 0; j--{
        nums1[c] = nums2[j]
        c--
    }
}