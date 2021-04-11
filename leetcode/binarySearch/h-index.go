package binarySearch


func hIndex(citations []int) int {
	b, e := 0, citations[0]
	for i := 0 ; i < len(citations); i++{
		if citations[i] > e{
			e = citations[i]
		}
		if citations[i] < b{
			b = citations[i]
		}
	}

	for;b+1 < e;{
		m := b + (e-b)/2
		if ifmeet(citations, m){
			b = m
		}else{
			e = m
		}
	}
	if ifmeet(citations, e){
		return e
	}
	return b
}

func ifmeet(citations []int, h int) bool{
	c := 0
	for i := 0; i < len(citations); i++{
		if citations[i] >= h{
			c++
		}
	}
	if c >= h{
		return true
	}
	return false
}
