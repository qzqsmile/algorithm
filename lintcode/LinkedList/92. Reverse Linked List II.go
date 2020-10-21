package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n{
		return head
	}
	dn := &ListNode{0, head}
	dc := dn
	count := n - m
	for m > 1{
		m--
		dc = dc.Next
	}
	pre := dc.Next
	cur := pre.Next
	fn := pre

	for count > 0{
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
		count--
	}
	fn.Next = cur
	dc.Next = pre

	return dn.Next
}
