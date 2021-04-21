package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//对于链表，一定要注意表头表尾的处理，能置空的一定要置空，防止野指针的出现
func reorderList(head *ListNode) {
	dummy := &ListNode{0, nil}
	cur := head
	count := 0
	for ; cur != nil; {
		count++
		cur = cur.Next
	}
	cur = head
	c := 1
	for ; c < (count+1)/2; {
		c++
		cur = cur.Next
	}
	cur2 := reverseReorder(cur.Next)
	cur.Next = nil
	cur1 := head
	count = 0
	cur = dummy
	for ; cur1 != nil || cur2 != nil; {
		if count%2 == 0 {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		count++
		cur = cur.Next
	}
}

func reverseReorder(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for ; cur != nil; {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for ; cur != nil; {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}
