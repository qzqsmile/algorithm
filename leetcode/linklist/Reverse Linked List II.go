/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    dummyNode := &ListNode{-1, head}
    fast, slow := dummyNode, dummyNode
    for ;fast != slow || (fast == dummyNode);{
        if fast == nil || fast.Next == nil{
            return nil
        }
        fast = fast.Next.Next
        slow = slow.Next 
    }
    
    cur1, cur2 := head, fast.Next
    for;cur1 != cur2;{
        cur1 = cur1.Next
        cur2 = cur2.Next
    }
    return cur1
}