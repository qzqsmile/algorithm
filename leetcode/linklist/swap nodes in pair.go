package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy1 := &ListNode{0, nil}
	dummy2 := &ListNode{0, nil}
	count := 0
	cur := head
	c1, c2 := dummy1, dummy2
	for ; cur != nil; {
		if count%2 == 0 {
			c1.Next = cur
			c1 = c1.Next
		} else {
			c2.Next = cur
			c2 = c2.Next
		}
		cur = cur.Next
		count++
	}
	c1.Next = nil
	c2.Next = nil
	count = 0
	c1, c2 = dummy1.Next, dummy2.Next
	cur = dummy2
	for ; c1 != nil && c2 != nil; {
		if count%2 == 0 {
			cur.Next = c2
			cur = c2
			c2 = c2.Next
		} else {
			cur.Next = c1
			cur = c1
			c1 = c1.Next
		}
		count++
	}
	if c1 != nil {
		cur.Next = c1
	}
	if c2 != nil {
		cur.Next = c2
	}
	return dummy2.Next
}
