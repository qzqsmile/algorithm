package datastructure

import "testing"

func TestSearch(t *testing.T){

	got := heapSort([]int{3,4,1,4,1,2,3})
	want := []int{1,1,2,3,3,4,4}

	if len(got) != len(want){
		t.Errorf("got %q, wanted %q", got, want)
	}
	for i := 0; i < len(got); i++{
		if got[i] != want[i]{
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
}
