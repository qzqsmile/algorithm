package array

func findKthLargest(nums []int, k int) int {
	heap := []int{}
	heap = append(heap, nums[0:k]...)
	makeHeap(heap)
	for i := k; i < len(nums); i++{
		if nums[i] > heap[0]{
			heap[0] = nums[i]
			adjustHeap(heap, 0)
		}
	}
	return heap[0]
}

func makeHeap(nums []int){
	for i := len(nums) / 2; i >= 0; i--{
		adjustHeap(nums, i)
	}
}

func adjustHeap(nums []int,  i int){
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
		adjustHeap(nums, min)
	}
}
