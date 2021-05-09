package twopointer


type ListNode struct {
 Val int
 Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	d := &ListNode{0, head}
	c, s, l := d, head, head
	for;l != nil;{
		for; l != nil && s != nil && l.Val == s.Val;{
			l = l.Next
		}
		if s != nil{
			c.Next = s
			c = s
			s = l
		}
	}
	c.Next = nil
	return d.Next
}
