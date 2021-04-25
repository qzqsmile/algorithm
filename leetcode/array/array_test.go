package arraya

import "testing"

func TestSearch(t *testing.T){

	got := search([]int{1,3,4,5,6,7,8,12}, 7)
	want := 5

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}