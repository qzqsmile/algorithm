package binarySearch

import "testing"

func TestSearch(t *testing.T){

	got := search([]int{1,3,4,5,6,7,8,12}, 7)
	want := 5

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSearchRange(t *testing.T){
	got := searchRange([]int{1,3,4,6,6,6,7,8}, 6)
	want := []int{3, 5}
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("got %q, wanted %q", got, want)
	}}