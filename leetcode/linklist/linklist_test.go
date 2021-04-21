package linklist

import "testing"

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
