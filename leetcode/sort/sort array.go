func sortArray(nums []int) []int {
    heapSort(nums)
    return nums
}

func heapSort(nums []int){
    index := len(nums) / 2
    for i := index; i >= 0; i--{
        adjustNode(nums, i)
    }
    
    tmp := []int{}
    for ;len(nums) > 0;{
        tmp = append(tmp, nums[0])
        nums[0] = nums[len(nums)-1]
        nums = nums[0:len(nums)-1]
        adjustNode(nums, 0)
    }
    nums = append(nums, tmp...)
}

func adjustNode(nums []int, n int){
    l, r := 2*n+1, 2*n+2
    minIndex := n 
    if l < len(nums) && nums[minIndex] > nums[l]{
        minIndex = l
    }
    if r < len(nums) && nums[minIndex] > nums[r]{
        minIndex = r 
    }
	// attention 1
    if minIndex != n{
        nums[minIndex], nums[n] = nums[n], nums[minIndex]
        adjustNode(nums, minIndex)
    }
}



func quicksort(nums []int, b int, e int){
    if b >= e{
        return 
    }
    b1, e1 := b, e
    for;b1 <= e1;{
        for;b1 <= e1 && nums[b1] < nums[e]; b1++{}
        for;b1 <= e1 && nums[e1] >= nums[e]; e1--{}
        if b1 <= e1{
            nums[b1], nums[e1] = nums[e1], nums[b1]
        }
    }
	//attention
    nums[b1], nums[e] = nums[e], nums[b1]

    quicksort(nums, b, b1-1)
    quicksort(nums, b1+1, e)
}

func mergesort(nums []int, b int, e int){
    if b >= e{
        return 
    }
    
    m := b + (e-b)/2
	//not mergesort(nums, b, m-1)
    mergesort(nums, b, m)
    mergesort(nums, m+1, e)
    tmp := []int{}
    
    i, j := b, m+1
    for;i <= m && j <= e;{
        if nums[i] < nums[j]{
            tmp = append(tmp, nums[i])
            i++
        }else{
            tmp = append(tmp, nums[j])
            j++
        }
    }

    if i <= m{
        tmp = append(tmp, nums[i:m+1]...)
    }
    if j <= e{
        tmp = append(tmp, nums[j:e+1]...)
    }

    for i := 0; i < len(tmp); i++{
        nums[i+b] = tmp[i]
    }
}