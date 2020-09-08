package main

import "fmt"

func quick_sort(start int, end int, q []int){
	if start >= end{
		return
	}
	pivot := q[(start+end)/2]
	fmt.Println(pivot)
	b, e := start, end
	for; b <= e; {
		for ; b <= e && q[b] < pivot;{
			b++
		}
		for ; b <= e && q[e] > pivot;{
			e--
		}
		if b <= e {
			q[b], q[e] = q[e], q[b]
			b++
			e--
		}
	}
	quick_sort(start, e, q)
	quick_sort(b, end, q)

}

func main(){
	s := []int{3,4,1,1,555555,1,1,1,23,41,1}
	quick_sort(0, len(s)-1, s)
	fmt.Println("s is %v", s)
}
