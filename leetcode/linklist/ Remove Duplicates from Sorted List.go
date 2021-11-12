/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    c := head 
    
    for;c != nil;{
        nextNode := c.Next
        for;nextNode != nil && nextNode.Val == c.Val;{
            nextNode = nextNode.Next
        }
        c.Next = nextNode
        c = nextNode
    }

    return head
}