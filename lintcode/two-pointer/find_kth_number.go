package main


func findKthLargest(nums []int, k int) int {
	return helper1(nums, 0, len(nums)-1, len(nums)-k+1)
}

func helper1(nums []int, start int, end int, k int) int{
	b, e := start, end
	povit := nums[(start+end)/2]

	for;b <= e;{
		for;b <= e && nums[b] < povit; b++{
		}
		for;b <= e && nums[e] > povit; e--{
		}
		if b <= e{
			nums[b], nums[e] = nums[e], nums[b]
			b++
			e--
		}
	}

	if b+1 == k{
		return nums[b]
	}else if b + 1 >k{
		return helper1(nums, start, e, k)
	}else{
		return helper1(nums, b, end, k-b)
	}
}

func main(){
	s := []int{3,2,3,1,2,4,5,5,6}
	findKthLargest(s, 4)
}
