package main


func sortColors2 (colors *[]int, k int)  {
	// write your code here
	helper(colors, 0, len(*colors)-1, 1, k)
}

func helper(colors *[]int, start int, end int, s1 int, e1 int){
	if start >= end || s1 >= e1{
		return
	}
	p := (s1 + e1) / 2
	b, e := start, end
	for;b <= e;{
		for ;b <= e && (*colors)[b] <= p;{
			b++
		}
		for ;b <= e && (*colors)[e] > p;{
			e--
		}
		if b <= e {
			(*colors)[b], (*colors)[e] = (*colors)[e], (*colors)[b]
			b++
			e--
		}
	}
	helper(colors, start, e, s1, p)
	helper(colors, b, end, p+1, e1)
}