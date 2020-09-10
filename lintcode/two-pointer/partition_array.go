package main

func partitionArray (nums []int, k int) int {
	// write your code here
	if len(nums) == 0{
		return 0
	}
	i, j := 0, len(nums)-1
	for;i < j;{
		for ;i < j && nums[i] < k; i++{
		}
		for ;i < j && nums[j] >= k; j--{

		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if i > j{
		if nums[j] < k{
			return j
		}
	}
	if nums[i] < k{
		return i+1
	}
	return i
}


func main()  {
	partitionArray([]int{}, 9)
}