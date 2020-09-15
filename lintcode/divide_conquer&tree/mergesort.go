package main

import "fmt"

func mergesort(start int, end int, q *[]int){
	if start >= end{
		return
	}
	mid := start + (end - start) / 2
	mergesort(start, mid, q)
	mergesort(mid+1, end, q)
	res := []int{}
	b, e := start, mid+1
	for; b <= mid && e <= end;{
		if (*q)[b] < (*q)[e]{
			res = append(res, (*q)[b])
			b++
		}else{
			res = append(res, (*q)[e])
			e++
		}
	}
	for ; b <= mid; b++{
		res = append(res, (*q)[b])
	}

	for; e <= end; e++{
		res = append(res, (*q)[e])
	}

	b = start
	for i, _ := range res{
		(*q)[b] = res[i]
		b++
	}
	fmt.Printf("res is %v mid is %v \n", res, mid)
}

func main(){
	s := []int{5, 2, 3, 1}
	mergesort(0, len(s)-1, &s)
	fmt.Println("s is %v", s)
}
