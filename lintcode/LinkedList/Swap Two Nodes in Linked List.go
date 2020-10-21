package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * @param head: a ListNode
 * @param v1: An integer
 * @param v2: An integer
 * @return: a new head of singly-linked list
 */
func swapNodes (head *ListNode, v1 int, v2 int) *ListNode {
	// write your code here
	d := &ListNode{0, head}
	p, c := d, d.Next
	var p1, p2 *ListNode
	p1, p2= nil, nil
	for c != nil{
		if c.Val == v1{
			p1 = p
		}
		if c.Val == v2{
			p2 = p
		}
		p = p.Next
		c = c.Next
	}
	if head == nil || p1 == nil || p2 == nil {
		return head
	}

	p1.Next, p2.Next = p2.Next, p1.Next
	p1.Next.Next, p2.Next.Next = p2.Next.Next, p1.Next.Next
	return d.Next
}
