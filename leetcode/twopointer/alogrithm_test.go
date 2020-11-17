package main

import "testing"

func Testlength(t *testing.T) {
	want := 3
	got := lengthOfLongestSubstring("abcabcbb")
	if got != 3 {
		t.Errorf("add was incorrect, got: %d, want: %d.", got, want)
	}
}


