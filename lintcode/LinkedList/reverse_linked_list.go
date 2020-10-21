package main

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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// no recursion
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil{
//		return head
//	}
//	pre, cur := head, head.Next
//	pre.Next = nil
//	n := ListNode{0, head}
//
//	for cur != nil{
//		tmp := cur.Next
//		cur.Next = pre
//		pre = cur
//
//		n.Next = pre
//		cur = tmp
//	}
//	return n.Next
//}
