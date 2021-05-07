package datastructure

import "fmt"

func heapSort(nums[] int) []int{
	makeHeap(nums)
	var sortedNums []int
	for ;len(nums) > 0;{
		sortedNums = append(sortedNums, nums[0])
		nums[0] = nums[len(nums)-1]
		nums = nums[0:len(nums)-1]
		adjustHeap(nums, 0)
	}
	return sortedNums
}

func makeHeap(nums[] int){
	for i := len(nums)/2; i >= 0; i--{
		adjustHeap(nums, i)
	}
}

func adjustHeap(nums[]int, i int){
	l, r := 2 * i + 1, 2 * i +2
	min := i
	if len(nums) > l && nums[l] < nums[min]{
		min = l
	}
	if len(nums) > r && nums[r] < nums[min]{
		min = r
	}
	if min != i{
		nums[min], nums[i] = nums[i], nums[min]
		adjustHeap(nums, min)
	}
}


func main(){
	nums := []int{2,3,4,5,6}
	sortedNums := heapSort(nums)
	fmt.Println(sortedNums)
}