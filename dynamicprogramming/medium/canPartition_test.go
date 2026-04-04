package medium

import "testing"

func TestCanPartition(t *testing.T) {
	type Case struct {
		nums   []int
		expect bool
	}
	cases := []Case{
		{
			nums: []int{
				1,
				5,
				11,
				5,
			},
			expect: true,
		},
		{
			nums: []int{
				1,
				2,
				3,
				5},
			expect: false,
		},
	}
	for _, c := range cases {
		got := canPartition(c.nums)
		if got != c.expect {
			t.Errorf("canPartition(%v) == %t, want %t", c.nums, got, c.expect)
		}
	}
}
