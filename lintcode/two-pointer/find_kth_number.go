package main


func findKthLargest(nums []int, k int) int {
	return myhelper(nums, 0, len(nums)-1, len(nums)-k+1)
}

func myhelper(nums []int, start int, end int, k int) int{
	povit := nums[(start+end)/2]
	k1 := helper1(nums, 0, len(nums)-1, povit)
	if k1+1 == k {
		return povit
	} else if k1+1 > k{
		return myhelper(nums, start, k1-1, k)
	} else{
		return myhelper(nums, k1, end, k-k1)
	}
}

func helper1(nums []int, start int, end int, povit int) int{
	b, e := start, end
	for;b < e; {
		for;b < e && nums[b] < povit; b++{
		}
		for;b < e && nums[e] >= povit; e--{
		}
		if b < e{
			nums[b], nums[e] = nums[e], nums[b]
			b++
			e--
		}
	}
	if nums[b] < povit{
		return b + 1
	}
	return b
}

func main(){
	s := []int{3,2,1,5,6,4}
	findKthLargest(s, 2)
}
