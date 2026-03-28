package medium

import "testing"

func Test_MinPathSum(t *testing.T) {
	type Case struct {
		grid     [][]int
		expected int
	}
	cases := []Case{
		{
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: 12,
		},
		{
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			expected: 7,
		},
	}
	for i, c := range cases {
		got := MinPathSum(c.grid)
		if got != c.expected {
			t.Errorf("grid%d, expected: %v, got: %v", i, c.expected, got)
		}
	}
}
