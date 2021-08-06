package main


type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	//var h1, h2, l1, l2 *ListNode
	l1 := &ListNode{0, nil}
	l2 := &ListNode{0, nil}
	h1, h2 := l1, l2
	for ;head != nil;{
		if head.Val < x{
			l1.Next = head
			l1 = l1.Next
		}else{
			l2.Next = head
			l2 = l2.Next
		}
		head = head.Next
	}
	h1.Next = h2.Next
	return h1.Next
}

func main()  {

}