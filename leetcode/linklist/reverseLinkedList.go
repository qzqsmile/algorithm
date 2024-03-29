package linklist


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode1 struct{
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var pre,  l *ListNode

	dummy := &ListNode{0, head}
	count := right - left
	c := 0
	p := dummy

	for;c != left-1; c++{
		p = p.Next
	}
	l = p.Next
	cur := p.Next

	for;count >= 0; count--{
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	p.Next = pre
	l.Next = cur

	return dummy.Next
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    h, _ := helper(head)
    return h
}


func helper(head *ListNode) (*ListNode, *ListNode){
    if head == nil || head.Next == nil{
        return head, head
    }
    h,  t := helper(head.Next)
    head.Next = nil 
    t.Next = head 
    return h, t.Next
}