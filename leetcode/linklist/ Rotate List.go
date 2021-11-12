/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil{
        return head
    }
    count := 0
    cur := head 
    for;cur != nil;{
        count++
        cur = cur.Next
    }

    k = k % count
    b, e := head, head
    for;e.Next != nil;{
        e = e.Next
        k--
        if k < 0{
            b = b.Next
        }
    }
    e.Next = head 
    newhead := b.Next
    b.Next = nil
    return newhead
}