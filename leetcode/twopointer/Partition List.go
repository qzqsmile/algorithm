package twopointer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	d1, d2 := &ListNode{0, nil}, &ListNode{0, nil}
	c1, c2 := d1, d2
	c := head
	for; c != nil;{
		if c.Val < x{
			c1.Next = c
			c1 = c1.Next
		}else{
			c2.Next = c
			c2 = c2.Next
		}
		c = c.Next
	}
	c1.Next = d2.Next
	c2.Next = nil
	return d1.Next
}