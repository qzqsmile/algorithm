package main

import "testing"

func Testlength(t *testing.T) {
	want := 3
	got := lengthOfLongestSubstring("abcabcbb")
	if got != 3 {
		t.Errorf("add was incorrect, got: %d, want: %d.", got, want)
	}
}
func TestMinimumWindow(t *testing.T){
	s := "ADOBECODEBANC"
	t1 := "ABC"
	wanted := "BANC"
	r := minWindow(s, t1)
	if r == wanted{
		t.Errorf("error was i")
	}
}

