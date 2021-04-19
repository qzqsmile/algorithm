package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//dummy Node在需要返回头指针时，使用

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{0, nil}
	dummy2 := &ListNode{0, nil}
	c1, c2 := dummy1, dummy2
	c := head
	for;c != nil;{
		if c.Val < x{
			c1.Next = c
			c1 = c
		}else{
			c2.Next = c
			c2 = c
		}
		c = c.Next
	}
	c1.Next = dummy2.Next
	c2.Next = nil
	return dummy1.Next
}