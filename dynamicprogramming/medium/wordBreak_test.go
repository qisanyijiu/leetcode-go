package medium

import "testing"

func Test_WordBreak(t *testing.T) {
	type Case struct {
		s        string
		wordDict []string
		expected bool
	}
	cases := []Case{
		{
			"applepenapple",
			[]string{"apple", "pen"},
			true,
		},
		{
			"leetcode",
			[]string{"leet", "code"},
			true,
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			false,
		},
	}
	for i, c := range cases {
		got := wordBreak(c.s, c.wordDict)
		if got != c.expected {
			t.Errorf("case %d: got %t, want %t", i, got, c.expected)
		}
	}
}
