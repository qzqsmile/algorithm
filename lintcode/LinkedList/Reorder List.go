package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	cc := 0
	if count%2 == 0 {
		cc = count / 2
	} else {
		cc = count/2 + 1
	}
	cur = head
	for cur != nil && cc > 1 {
		cur = cur.Next
		cc--
	}
	d1, d2 := &ListNode{0, head}, &ListNode{0, cur.Next}
	cur.Next = nil
	d2.Next = reverseLink(d2.Next)
	d := &ListNode{0, head}
	c := d
	c1, c2 := d1.Next, d2.Next
	count = 0
	for c1 != nil && c2 != nil {
		if count%2 == 0 {
			c.Next = c1
			c1 = c1.Next
		} else {
			c.Next = c2
			c2 = c2.Next
		}
		c = c.Next
		count++
	}
	if c1 != nil {
		c.Next = c1
	}
	if c2 != nil {
		c.Next = c2
	}
}

func reverseLink(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre, cur := head, head.Next
	pre.Next = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

