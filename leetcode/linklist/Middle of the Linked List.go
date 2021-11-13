/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func middleNode(head *ListNode) *ListNode {
    cur := head 
    count := 0
    for;cur != nil;{
        count++
        cur = cur.Next
    }

    return findKth(head, (count+2)/2)
}

func findKth(head *ListNode, k int) *ListNode{
    cur := head 
    for;k > 1;{
        k--
        cur = cur.Next
    }
    return cur
}