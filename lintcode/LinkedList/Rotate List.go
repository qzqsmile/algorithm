package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil{
		return head
	}
	cur := head
	count := 1
	for cur.Next != nil{
		cur = cur.Next
		count++
	}
	cur.Next = head
	k = k % count
	s := count - k
	d := &ListNode{0, head}
	pre := d
	for s > 0{
		s--
		pre = pre.Next
	}
	res := pre.Next
	pre.Next = nil
	return res
}
