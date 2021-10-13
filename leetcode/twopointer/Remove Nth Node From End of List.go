package twopointer
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    b, e := dummy, dummy
    for;e != nil;{
        e = e.Next
        if n < 0{
            b = b.Next 
        }
        n--
    }
    b.Next = b.Next.Next
    return dummy.Next
}