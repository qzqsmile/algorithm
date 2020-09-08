package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	h, e := helper(head)
	e.Next = nil
	return h
}

func helper(head *ListNode) (*ListNode, *ListNode){
	if head == nil || head.Next == nil{
		return head, head
	}
	h, e := helper(head.Next)
	e.Next = head
	return h, head
}
