/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	if len(lists) == 0{
		return nil
	}
	res = lists[0]
	for i := 1; i < len(lists); i++{
		res = merge(res, lists[i])
	}
	return res
}

func merge(l *ListNode, r *ListNode) *ListNode{
	t := &ListNode{-1, nil}
	h := t
	l_c, r_c := l, r
	for ;l_c != nil && r_c != nil;{
		if l_c.Val < r_c.Val{
			t.Next = l_c
			l_c = l_c.Next
			t = t.Next
		}else{
			t.Next = r_c
			r_c = r_c.Next
			t = t.Next
		}
	}

	if l_c != nil{
		t.Next = l_c
	}else{
		t.Next = r_c
	}
	return h.Next
}