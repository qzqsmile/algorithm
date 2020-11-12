package main

func searchInsert(nums []int, target int) int {
	A := nums
	if len(A) == 0{
		return 0
	}else if target > A[len(A)-1]{
		return len(A)
	}else if target < A[0]{
		return 0
	}
	b, e := 0, len(A)-1
	for;b+1<e;{
		mid := b + (e - b) /2
		if A[mid] > target{
			e = mid
		}else if A[mid] < target{
			b = mid
		}else{
			return mid
		}
	}

	if A[b] == target{
		return b
	}
	return b+1
}
