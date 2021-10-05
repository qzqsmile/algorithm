package linklist

import "testing"

// type ListNode struct {
// 	Val int
// 	Next *ListNode
// }

func TestReverseBetween(t *testing.T){
	head := &ListNode{1, &ListNode{2, &ListNode{3,
		&ListNode{4, &ListNode{5, nil}}}}}
	got := reverseBetween(head,2,4)
	want := &ListNode{1, &ListNode{4, &ListNode{3,
		&ListNode{2, &ListNode{5, nil}}}}}
	for ;got != nil && want != nil;{
		if got.Val != want.Val{
			t.Errorf("Error here!")
		}
		got = got.Next
		want = want.Next
	}
}


func TestReverseNodesKGroup(t *testing.T){
	head := &ListNode{1, &ListNode{2, &ListNode{3,
		&ListNode{4, &ListNode{5, nil}}}}}
	got := reverseKGroup(head,2)
	want := &ListNode{2, &ListNode{1, &ListNode{4,
		&ListNode{3, &ListNode{5, nil}}}}}
	cur1 := got
	cur2 := want
	for ;cur1 != nil && cur2 != nil;{
		if cur1.Val != cur2.Val{
			t.Errorf("Error here!")
		}
		cur1 = cur1.Next
		cur2 = cur2.Next

	}
	if cur1 != nil || cur2 != nil {
		t.Errorf("Error here!")

	}
}

// [4,2,1,3]
func TestSortList(t *testing.T){
	input := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	wanted := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	output := sortList(input)
	c1, c2 := output, wanted

	for;c1 != nil && c2 != nil;{
		if c1.Val != c2.Val{
			t.Errorf("Sort List: output is %v, wanted is %v", c1.Val, c2.Val)
		}
		c1 = c1.Next
		c2 = c2.Next
	}
}
