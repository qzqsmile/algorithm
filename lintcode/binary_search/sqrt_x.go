package main

func mySqrt(x int) int {
	if x == 0{
		return 0
	}
	b, e := 1, x
	for ;b + 1 < e;{
		mid := b + (e - b)/2
		if mid*mid == x{
			return mid
		}else if mid*mid > x{
			e = mid
		}else{
			b = mid
		}
	}
	if e * e < x{
		return e
	}
	return b
}
