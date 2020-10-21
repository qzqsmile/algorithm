package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	c1, c2 := l1, l2
	h := &ListNode{0, nil}
	c := h
	for c1 != nil && c2 != nil {
		if c1.Val < c2.Val {
			c.Next = c1
			c1 = c1.Next
		} else {
			c.Next = c2
			c2 = c2.Next
		}
		c = c.Next
	}
	if c1 != nil {
		c.Next = c1
	}
	if c2 != nil {
		c.Next = c2
	}
	return h.Next
}

func main() {

}
