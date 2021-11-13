/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode)  {
    count := 0
    cur := head
    for;cur != nil;{
        cur = cur.Next 
        count++
    }

    firstTail := findKth(head, (count+1)/2)
    secondHead := firstTail.Next
    firstTail.Next = nil 
    secondHead = reverse(secondHead)

    cur1, cur2 := head, secondHead
    for;cur1 != nil && cur2 != nil;{
        tmp1 := cur1.Next
        tmp2 := cur2.Next
        cur1.Next = cur2 
        cur2.Next = tmp1 
        cur1 = tmp1 
        cur2 = tmp2
    }
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

func findKth(head *ListNode, k int) *ListNode{
    cur := head
    for;cur != nil;{
        k--
        if k == 0{
            return cur
        }
        cur = cur.Next
    }
    return nil
}

