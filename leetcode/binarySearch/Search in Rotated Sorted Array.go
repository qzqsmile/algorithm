package binarySearch

// 这题的使用二分法是很容易想到的。二分法的使用有一个前提即是排好序的数组
// 所以这里使用nums[m] > nums[b]的条件判断m点落在哪个连续的空间，然后进行二分。
func search(nums []int, target int) int {
	b, e := 0, len(nums)-1
	for; b + 1 < e; {
		m := b + (e-b)/2
		if nums[m] > nums[b]{
			//因为这段数字是连续的，如果不在这段区间，那么就在另一段区间
			if target >= nums[b] && target <= nums[m]{
				e = m
			}else{
				b = m
			}
		}else{
			//同理
			if target >= nums[m] && target <= nums[e]{
				b = m
			}else{
				e = m
			}
		}
	}
	if nums[b] == target{
		return b
	}
	if nums[e] == target{
		return e
	}
	return -1
}


// another solution， low一点但是很好想
// 先找到最小值，再分别binary search

// func search(nums []int, target int) int {
//     b, e := 0, len(nums)-1
//     for;b + 1 < e;{
//         m := b + (e-b)/2
//         if nums[m] > nums[b]{
//             if nums[b] < nums[e]{
//                 e = m
//             }else{
//                 b = m
//             }
//         }else {
//             e = m
//         }
//     }
//     index := -1
//     if nums[b] < nums[e]{
//         index = b
//     }else{
//         index = e
//     }
//     l := binarySearch(nums[:index], target)
//     r := binarySearch(nums[index:], target)
//     if l != -1{
//         return l 
//     }
//     if r != -1{
//         return r+index
//     }
//     return -1
// }

// func binarySearch(nums []int, target int) int{
//     if len(nums) == 0{
//         return -1
//     }
//     b, e := 0, len(nums)-1
//     for;b + 1 < e;{
//         m := b + (e-b)/2
//         if nums[m] < target{
//             b = m
//         }else if nums[m] > target{
//             e = m
//         }else{
//             return m
//         }
//     }
//     if nums[b] == target{
//         return b
//     }
//     if nums[e] == target{
//         return e
//     }
//     return -1
// }