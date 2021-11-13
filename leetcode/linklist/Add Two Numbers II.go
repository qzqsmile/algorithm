/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    rl1 := reverse(l1)
    rl2 := reverse(l2)
    newHead := addTwo(rl1, rl2)
    return reverse(newHead)
}

func addTwo(l1 *ListNode, l2 *ListNode) *ListNode{
    c1, c2 := l1, l2 
    sum, carryBit := 0, 0
    dummyNode := &ListNode{0, nil}
    pre := dummyNode

    for;c1 != nil && c2 != nil;{
        sum = carryBit + c1.Val + c2.Val
        pre.Next = &ListNode{sum % 10, nil}
        pre = pre.Next
        carryBit = sum / 10 
        c1 = c1.Next
        c2 = c2.Next
    }

    for;c1 != nil;{
        sum = carryBit + c1.Val
        pre.Next = &ListNode{sum % 10, nil}
        pre = pre.Next
        carryBit = sum / 10
        c1 = c1.Next
    }

    for;c2 != nil;{
        sum = carryBit + c2.Val
        pre.Next = &ListNode{sum % 10, nil}
        pre = pre.Next
        carryBit = sum / 10
        c2 = c2.Next
    }
    if carryBit > 0{
        pre.Next = &ListNode{carryBit, nil}
    }
    return dummyNode.Next
}

func reverse(head *ListNode) *ListNode{
    var pre *ListNode
    cur := head 
    for;cur != nil;{
        tmp := cur.Next
        cur.Next = pre 
        pre = cur 
        cur = tmp
    }
    return pre
}