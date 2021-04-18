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
	}
}
