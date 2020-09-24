package main

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, start int, end int, k int) int {
	pvoit := nums[start]
	b, e := start, end

	for ; b <= e; {
		for ; b <= e && nums[b] < pvoit; b++ {

		}
		for ; b <= e && nums[e] >= pvoit; e-- {

		}
		if b <= e {
			nums[b], nums[e] = nums[e], nums[b]
			b++
			e--
		}
	}

	if k > b {
		if nums[b] == pvoit {
			return quickSelect(nums, b+1, end, k)
		} else {
			return quickSelect(nums, b, end, k)
		}
	} else if k < b {
		return quickSelect(nums, start, b-1, k)
	} else {
		return pvoit
	}
}

func main() {

}
