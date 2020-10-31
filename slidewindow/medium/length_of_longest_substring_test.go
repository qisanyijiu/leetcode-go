package medium

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct{
		input string
		out int
	}{
		{input: "abcabcbb", out: 3},
		{input: "bbbbb", out: 1},
		{input: "pwwkew", out: 3},
	}
	for _, item := range cases {
		res := LengthOfLongestSubstring(item.input)
		if res != item.out {
			t.Errorf("input: %s, except: %d, got: %d", item.input, item.out, res)
		}
	}
}