package medium

import "testing"

func Test_LongestCommonSubsequence(t *testing.T) {
	type Case struct {
		s1, s2   string
		excepted int
	}
	cases := []Case{
		{
			"abcde",
			"ace",
			3,
		},
		{
			"abc",
			"abc",
			3,
		},
		{
			"abc",
			"def",
			0,
		},
	}
	for _, c := range cases {
		out := longestCommonSubsequence(c.s1, c.s2)
		if out != c.excepted {
			t.Errorf("LongestCommonSubsequence(%s, %s) = %d; want %d", c.s1, c.s2, out, c.excepted)
		}
	}
}
