package medium

import "testing"

func TestCanJump(t *testing.T) {
	type Case struct {
		Nums     []int
		Expected bool
	}
	cases := []Case{
		{[]int{1, 2, 3}, true},
		{[]int{3, 2, 1}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{2, 3, 1, 1, 4}, true},
	}
	for _, c := range cases {
		got := canJump(c.Nums)
		if got != c.Expected {
			t.Errorf("canJump(%v) == %t, want %t", c.Nums, got, c.Expected)
		}
	}
}
