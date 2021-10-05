package linklist

func sortList(head *ListNode) *ListNode {
	c := 0
	cur := head
	for cur != nil {
		c++
		cur = cur.Next
	}
	return mergeSort(head, c)
}

func mergeSort(head *ListNode, n int) *ListNode {
	if n <= 1 {
		return head
	}

	// #1
	c := n / 2
	cc := c
	var h1, h2 *ListNode
	h1 = head
	cur := head
	for ; c != 1; c-- {
		cur = cur.Next
	}
	h2 = cur.Next
	cur.Next = nil

	// #2
	nh1 := mergeSort(h1, cc)
	nh2 := mergeSort(h2, n-cc)

	//# merge
	dummy := &ListNode{0, nil}
	cur = dummy
	c1, c2 := nh1, nh2
	for c1 != nil && c2 != nil {
		if c1.Val < c2.Val {
			cur.Next = &ListNode{c1.Val, nil}
			c1 = c1.Next
		} else {
			cur.Next = &ListNode{c2.Val, nil}
			c2 = c2.Next
		}
		cur = cur.Next
	}

	if c1 != nil {
		cur.Next = c1
	}

	if c2 != nil {
		cur.Next = c2
	}
	return dummy.Next
}
