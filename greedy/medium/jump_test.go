package medium

import "testing"

func TestJump(t *testing.T) {
	type Case struct {
		Nums     []int
		Expected int
	}
	cases := []Case{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
	}
	for _, c := range cases {
		got := jump(c.Nums)
		if got != c.Expected {
			t.Errorf("canJump(%v) == %d, want %d", c.Nums, got, c.Expected)
		}
	}
}
