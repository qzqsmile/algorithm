/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    dummyNode := &ListNode{-1, head}
    cur, count := head, 0

    for;cur != nil;{
        count++
        cur = cur.Next
    }

    times := count / k 
    preNode := dummyNode
    for;times > 0;{
        preNode.Next = reverse(preNode.Next, k)
        tmpk := k 
        for ;tmpk > 0; tmpk--{
            preNode = preNode.Next
        }
        times--
    }
    return dummyNode.Next
}

func reverse(head *ListNode, k int) *ListNode{
    var pre *ListNode  
    cur := head 
    for;k > 0;{
        tmp := cur.Next 
        cur.Next = pre 
        pre = cur 
        cur = tmp
        k--
    }
    head.Next = cur
    return pre
}
