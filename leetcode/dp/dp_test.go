package dp

import "testing"

func TestLengthOfLIS(t *testing.T) {
	input := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	output := lengthOfLIS(input)
	wanted := 6
	if output != wanted {
		t.Errorf("wanted is %v, however output is %v", wanted, output)
	}
}
