package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode1 struct{
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	l1 := &ListNode{0, nil}
	l2 := &ListNode{0, nil}
	h1, h2 := l1, l2
	cur := head
	for ;cur != nil;{
		if cur.Val < x{
			h1.Next = cur
			h1 = h1.Next
			h1.Next = nil
		}else{
			h2.Next = cur
			h2 = h2.Next
			h2.Next = nil
		}
		cur = cur.Next
	}
	h2.Next = nil
	h1.Next = l2.Next
	return l1.Next
}

func main()  {
	//[1,4,3,2,5,2]
	//3
	data := &ListNode{1, &ListNode{
		4, &ListNode{3,&ListNode{2,&ListNode{5,
			&ListNode{2, nil}}}},
	}}
	partition(data, 3)
}