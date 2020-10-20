package main

/**
 * @param nums: an integer array
 * @param k: An integer
 * @return: the top k largest numbers in array
 */
/**
 * @param nums: an integer array
 * @param k: An integer
 * @return: the top k largest numbers in array
 */
func topk (nums []int, k int) []int {
	// write your code here
	topK := []int{}
	topK = append(topK, nums[0: k]...)
	makeHeap(topK)
	for i := k; i < len(nums); i++{
		if nums[i] > topK[0]{
			topK[0] = nums[i]
			heapSort(topK, 0)
		}
	}
	res := []int{}
	for i := len(topK)-1; i >= 0; i--{
		res = append(res, topK[0])
		topK[0] = topK[i]
		topK = topK[0:i]
		heapSort(topK, 0)
	}
	reverse(res)
	return res
}

func reverse(nums []int){
	i, j := 0, len(nums)-1
	for i < j{
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func makeHeap(nums []int){
	for i := len(nums) / 2; i >= 0; i--{
		heapSort(nums, i)
	}
}

func heapSort(nums []int,  i int){
	l := 2 * i + 1
	r := 2 * i + 2
	min := i
	if l < len(nums) && nums[l] < nums[min]{
		min = l
	}
	if r < len(nums) && nums[r] < nums[min]{
		min = r
	}
	if i != min{
		nums[min], nums[i] = nums[i], nums[min]
		heapSort(nums, min)
	}
}
