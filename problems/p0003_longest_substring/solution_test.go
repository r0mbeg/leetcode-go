package p0003_longest_substring

import "testing"

func Test_longestSubstring_TableDriven(t *testing.T) {

	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"dvdf", 3},
		{"", 0},
	}

	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.s)
		if got != tt.want {
			t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
