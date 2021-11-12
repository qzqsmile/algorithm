/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    cur := head 
    n := 0
    for;cur != nil;{
        n++
        cur = cur.Next
    }

    h := (n+1) / 2
    cur = head
    for;h > 1;{
        h--
        cur = cur.Next
    }
    
    tail := cur
    newhead := reverse(tail.Next)

    cur = head 
    rc := newhead
    for;rc != nil;{
        if rc.Val != cur.Val{
            return false
        }
        rc = rc.Next
        cur = cur.Next
    }

    tail.Next = reverse(newhead)

    return true
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