package binarySearch

func mySqrt(x int) int {
	b, e := 0, x
	for;b + 1 < e;{
		m := b + (e-b)/2
		if m * m > x{
			e = m
		}else{
			b = m
		}
	}
	if b * b == x{
		return b
	}
	if e * e == x{
		return e
	}
	return b
}

