package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	cur := node
	var pre *ListNode
	for cur.Next != nil{
		cur.Val = cur.Next.Val
		pre = cur
		cur = cur.Next
	}
	pre.Next = nil
}