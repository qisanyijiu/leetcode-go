package medium

import "testing"

func Test_UniquePaths(t *testing.T) {
	type Case struct {
		m, n     int
		excepted int
	}
	cases := []Case{
		{m: 3, n: 2, excepted: 3},
		{m: 7, n: 3, excepted: 28},
		{m: 3, n: 3, excepted: 6},
	}
	for _, c := range cases {
		got := UniquePaths(c.m, c.n)
		if got != c.excepted {
			t.Errorf("case (%d:%d): got %d, want %d", c.m, c.n, got, c.excepted)
		}
	}
}
