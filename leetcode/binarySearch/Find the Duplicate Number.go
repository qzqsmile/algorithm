package binarySearch

func findDuplicate(nums []int) int {
	b, e := 1, len(nums)-1
	for;b+1 < e;{
		m := b + (e-b)/2
		if isduplicate(nums, b, m){
			e = m
		}else{
			b = m
		}
	}
	if isduplicate(nums, b, b){
		return b
	}
	return e
}

func isduplicate(nums []int, b int, e int) bool{
	count := 0
	for i := 0; i < len(nums); i++{
		if nums[i] >= b && nums[i] <= e{
			count++
		}
	}
	if count > e-b+1{
		return true
	}
	return false
}