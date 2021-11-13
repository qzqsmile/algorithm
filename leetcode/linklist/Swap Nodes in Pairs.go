/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
    dummyNode := &ListNode{0, head}
    pre := dummyNode
    for ;pre != nil && pre.Next != nil && pre.Next.Next != nil;{
        node1 := pre.Next
        node2 := pre.Next.Next
        tmp := node2.Next

        pre.Next = node2 
        node2.Next = node1
        node1.Next = tmp 
        
        pre = node1
    }
    return dummyNode.Next
}