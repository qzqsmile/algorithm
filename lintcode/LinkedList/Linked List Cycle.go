package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * @param head: The first node of linked list.
 * @return: True if it has a cycle, or false
 */
func hasCycle (head *ListNode) bool {
	// write your code here
	if head == nil{
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != fast{
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow == fast{
		return true
	}else{
		return false
	}
}
