package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	n1, n2 := 0, 0
	c1, c2 := headA, headB
	for c1 != nil {
		n1++
		c1 = c1.Next
	}
	for c2 != nil {
		n2++
		c2 = c2.Next
	}
	nc1, nc2 := headA, headB
	if n2 > n1 {
		d := n2 - n1
		for ; d > 0; d-- {
			nc2 = nc2.Next
		}
	} else {
		d := n1 - n2
		for ; d > 0; d-- {
			nc1 = nc1.Next
		}
	}
	for nc1 != nil && nc2 != nil {
		if nc1 == nc2 {
			return nc1
		}
		nc1 = nc1.Next
		nc2 = nc2.Next
	}
	//should never reach to this
	return nil
}
