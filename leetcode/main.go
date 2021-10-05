package main

import (
	"fmt"
	"sort"
)

func main(){
	r := threeSumClosest([]int{-1,2,1,-4}, 1)
	fmt.Println(r)
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	cur := dummy
	c1, c2 := l1, l2
	for ; c1 != nil && c2 != nil; {
		if c1.Val < c2.Val {
			cur.Next = c1
			c1 = c1.Next
		} else {
			cur.Next = c2
			c2 = c2.Next
		}
		cur = cur.Next
	}
	if c1 != nil {
		cur.Next = c1
	}

	if c2 != nil {
		cur.Next = c2
	}
	return dummy.Next
}