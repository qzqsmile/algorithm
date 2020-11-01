package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//merge version
func sortList(head *ListNode) *ListNode {
	cur := head
	count := 0
	for cur != nil{
		count++
		cur = cur.Next
	}
	if count == 0 || count == 1{
		return head
	}

	l := head
	cur = head
	half := count / 2
	var r *ListNode
	for half > 1{
		cur = cur.Next
		half--
	}
	r = cur.Next
	cur.Next = nil
	l = sortList(l)
	r = sortList(r)

	//merge
	dummyNode := &ListNode{0, nil}
	cur = dummyNode
	for l != nil && r != nil{
		if l.Val < r.Val{
			cur.Next = l
			l = l.Next
		}else{
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}

	if l != nil{
		cur.Next = l
	}
	if r != nil{
		cur.Next = r
	}
	return dummyNode.Next
}


//quick sort versiin

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	dummyNode1 := &ListNode{0, nil}
	dummyNode2 := &ListNode{0, nil}
	cur1, cur2 := dummyNode1, dummyNode2
	povit := head.Val
	cur := head.Next

	for cur != nil{
		if cur.Val <= povit{
			cur1.Next = cur
			cur1 = cur1.Next
		}else{
			cur2.Next = cur
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	cur1.Next = nil
	cur2.Next = nil
	l := sortList(dummyNode1.Next)
	r := sortList(dummyNode2.Next)
	if l != nil{
		cl := l
		for cl.Next != nil{
			cl = cl.Next
		}
		cl.Next = head
		head.Next = r
		return l
	}else{
		head.Next = r
		return head
	}
}