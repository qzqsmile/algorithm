package linklist


func reverseKGroup(head *ListNode, k int) *ListNode {
	d := &ListNode{0, head}
	cur := d.Next
	count := 0
	for;cur != nil;{
		count++
		cur = cur.Next
	}
	times := count / k
	for cur := d;times > 0; times--{
		cur.Next = reversekgroup(cur.Next, k)
		mk := k
		for;mk > 0; mk--{
			cur = cur.Next
		}
	}
	return d.Next
}

func reversekgroup(head *ListNode, k int) *ListNode{
	var pre *ListNode
	cur := head

	for ;k > 0; k--{
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	head.Next = cur
	return pre
}