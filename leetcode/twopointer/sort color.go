package twopointer


// 3 step
// step1: 先搞好起始条件
// step2: 循环的条件设置一样，这样防止corner case出现
// step3: 对于这种3段指针，注意第三个指针的起始条件

//三个变量就用三个指针
func sortColors(nums []int)  {
	left, right := 0, len(nums)-1
	mid := left

	for ;mid <= right;{
		if nums[mid] == 0{
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			//因为这里不可能有2所以可以++
			mid++
		}else if nums[mid] == 2{
			nums[right], nums[mid] = nums[mid], nums[right]
			right--
		}else{
			mid++
		}
	}
}

// 同向3指针
//func sortColors(nums []int)  {
//	b0, b1 := 0, 0
//	for c := 0; c < len(nums); c++{
//		if nums[c] == 0{
//			nums[b0], nums[c] = nums[c], nums[b0]
//			if b0 < b1{
//				nums[b1], nums[c] = nums[c], nums[b1]
//			}
//			b0++
//			b1++
//		}else if nums[c] == 1{
//			nums[b1], nums[c] = nums[c], nums[b1]
//			b1++
//		}
//	}
//}

