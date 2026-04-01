package medium

import "testing"

func Test_MaxProduct(t *testing.T) {
	type Case struct {
		nums     []int
		expected int
	}
	cases := []Case{
		{nums: []int{2, 3, -2, 4}, expected: 6},
		{nums: []int{-2, 0, -1}, expected: 0},
	}
	for i, c := range cases {
		actual := maxProduct(c.nums)
		if actual != c.expected {
			t.Errorf("case %d expected %d, actual %d", i, c.expected, actual)
		}
	}
}
