package binarySearch

func firstBadVersion(n int) int {
	b, e := 1, n
	for ;b + 1< e;{
		m := b + (e-b)/2
		if isBadVersion(m){
			e = m
		}else{
			b = m
		}
	}
	if isBadVersion(b) == true{
		return b
	}
	return e
}
