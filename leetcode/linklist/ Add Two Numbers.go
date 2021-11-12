/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    c1, c2 := l1, l2
    num := 0
    l := 0
    dummyNode := &ListNode{0, nil}
    c := dummyNode
    for;c1 != nil && c2 !=nil;{
        num = c1.Val + c2.Val + l 
        l = num / 10 
        num = num % 10
        c1 = c1.Next
        c2 = c2.Next
        c.Next = &ListNode{num, nil}
        c = c.Next
    }

    for;c1 != nil;{
        num = c1.Val + l 
        l = num / 10
        num = num % 10
        c1 = c1.Next
        c.Next = &ListNode{num, nil}
        c = c.Next
    }

    for;c2 != nil;{
        num = c2.Val + l 
        l = num / 10
        num = num % 10
        c2 = c2.Next
        c.Next = &ListNode{num, nil}
        c = c.Next
    }
    
    if l != 0{
        c.Next = &ListNode{l, nil}
    }
    return dummyNode.Next
}