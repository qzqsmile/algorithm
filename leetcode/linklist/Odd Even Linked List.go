/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func oddEvenList(head *ListNode) *ListNode {
    dummyNode1, dummyNode2 := &ListNode{0, nil}, &ListNode{0, nil}
    pre1, pre2 := dummyNode1, dummyNode2
    cur, count := head, 0
    
    for;cur != nil;{
        count++
        if count % 2 == 1{
            pre1.Next = cur 
            pre1 = cur
        }else{
            pre2.Next = cur 
            pre2 = cur
        }
        cur = cur.Next
    }
    pre1.Next = dummyNode2.Next 
    pre2.Next = nil
    return dummyNode1.Next
}