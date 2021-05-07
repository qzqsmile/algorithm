package array

import "testing"

func TestHeap(t *testing.T){
	got := findKthLargest([]int{3,2,1,5,6,4}, 2)
	wanted := 5
	if got != wanted{
		t.Errorf("got %d, wanted %d", got, wanted)
	}
}