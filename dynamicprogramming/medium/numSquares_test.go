package medium

import "testing"

func Test_NumSquares(t *testing.T) {
	type Case struct {
		num      int
		expected int
	}
	cases := []Case{
		{num: 12, expected: 3},
		{num: 13, expected: 2},
	}
	for _, c := range cases {
		actual := NumSquares(c.num)
		if actual != c.expected {
			t.Errorf("NumSquares(%d) == %d, want %d", c.num, actual, c.expected)
		}
	}

}
