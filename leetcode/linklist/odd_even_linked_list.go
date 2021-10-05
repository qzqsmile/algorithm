package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy1 := &ListNode{0, head}
	dummy2 := &ListNode{0, nil}
	pre, cur2 := dummy1, dummy2
	cur1 := head
	c := 0
	for cur1 != nil {
		c++
		if c%2 == 0 {
			pre.Next = cur1.Next
			cur1.Next = nil

			cur2.Next = cur1
			cur2 = cur2.Next

			cur1 = pre.Next
		} else {
			cur1 = cur1.Next
			pre = pre.Next
		}
	}

	pre.Next = dummy2.Next

	return dummy1.Next
}
