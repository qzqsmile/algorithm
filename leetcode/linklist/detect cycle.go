package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if !hasCycle(head){
		return nil
	}
	cur1, cur2 := head, head
	for;cur2 != nil && cur2.Next != nil;{
		cur1 = cur1.Next
		cur2 = cur2.Next.Next
		if cur1 == cur2{
			break
		}
	}
	cur1 = head
	for;cur1 != cur2;{
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return cur1
}

func hasCycle(head *ListNode) bool {
	cur1, cur2 := head, head
	for;cur2 != nil && cur2.Next != nil;{
		cur1 = cur1.Next
		cur2 = cur2.Next.Next
		if cur1 == cur2{
			return true
		}
	}
	return false
}
