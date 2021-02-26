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
	}
}

func TestSearchMatrix(t *testing.T){
	matrix := [][]int{
		{1,2,3,4,5},
		{6,7,8,9,10},
		{11,12,13,14,15},
		{16,17,18,19,20},
		{21,22,23,24,25}}
	target := 15
	flag := searchMatrix1(matrix, target)
	if !flag{
		t.Errorf("should be true")
	}
}
