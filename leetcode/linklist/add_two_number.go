package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c1, c2 := l1, l2
	dummy := &ListNode{0, nil}
	cur := dummy
	h := 0
	for;c1 != nil && c2 != nil;{
		v := (c1.Val + c2.Val + h) % 10
		cur.Next = &ListNode{v, nil}
		cur = cur.Next
		h = (c1.Val + c2.Val + h) / 10
		c1 = c1.Next
		c2 = c2.Next
	}
	for;c1 != nil;{
		v := (c1.Val + h) % 10
		cur.Next = &ListNode{v, nil}
		cur = cur.Next
		h = (c1.Val + h) / 10
		c1 = c1.Next
	}
	for;c2 != nil;{
		v := (c2.Val + h) % 10
		cur.Next = &ListNode{v, nil}
		cur = cur.Next
		h = (c2.Val + h) / 10
		c2 = c2.Next
	}
	if h > 0{
		cur.Next = &ListNode{h, nil}
	}
	return dummy.Next
}