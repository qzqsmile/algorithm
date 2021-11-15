/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func sortList(head *ListNode) *ListNode {
    count, cur := 0, head
    for;cur != nil;{
        count++
        cur = cur.Next
    }
    newHead := mergeSort(head, 1, count)
    return newHead
}

func mergeSort(head *ListNode, left int, right int)*ListNode{
    if left >= right{
        return head
    }
    m := left + (right-left)/2
    leftTail := findKth(head, m)
    rightHead := leftTail.Next
    leftTail.Next = nil 

    leftHead := mergeSort(head, left, m)
    rightHead = mergeSort(rightHead, 1, right-m)

    dummyNode := &ListNode{0, nil}
    pre := dummyNode
    c1, c2 := leftHead, rightHead
    for;c1 != nil && c2 != nil;{
        if c1.Val < c2.Val{
            pre.Next = c1 
            c1 = c1.Next
        }else{
            pre.Next = c2
            c2 = c2.Next
        }
        pre = pre.Next
    }

    if c1 != nil{
        pre.Next = c1
    }
    
    if c2 != nil{
        pre.Next = c2
    }

    return dummyNode.Next
}


func findKth(head *ListNode, k int) *ListNode{
    cur := head
    for;k > 1; k--{
        cur = cur.Next
    }
    return cur
}

