package binarySearch

//感觉这题没啥意思，因为最差的条件只能是O(n),换句话说你遍历一遍，总能AC


func findMin(nums []int) int {
	b, e := 0, len(nums)-1
	for; b + 1 <e;{
		m := b + (e-b)/2
		//这题的关键在这里判断m位于何处
		if nums[m] > nums[e]{
			b = m + 1
		}else if nums[m] < nums[e]{
			e = m
		}else{
			//如何相等则减1
			e--
		}
	}
	if nums[b] < nums[e]{
		return nums[b]
	}
	return nums[e]
}
