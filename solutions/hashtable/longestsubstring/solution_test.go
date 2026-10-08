package longestsubstring

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
		{"dvdf", 3},
		{"abba", 2},
	}

	for _, test := range tests {
		fmt.Printf("Running test case: input=%q, want=%d\n", test.input, test.want)
		if got := lengthOfLongestSubstring(test.input); got != test.want {
			t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}
