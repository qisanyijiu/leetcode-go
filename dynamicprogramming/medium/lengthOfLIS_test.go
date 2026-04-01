package medium

import "testing"

func Test_LengthLTS(t *testing.T) {
	type Case struct {
		nums     []int
		excepted int
	}
	cases := []Case{
		{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, excepted: 4},
		{nums: []int{0, 1, 0, 3, 2, 3}, excepted: 4},
		{nums: []int{7, 7, 7, 7, 7, 7, 7}, excepted: 1},
	}
	for i, c := range cases {
		ret := lengthOfLIS(c.nums)
		if ret != c.excepted {
			t.Errorf("case[%d], excepted:%v, got:%v", i, c.excepted, ret)
		}
	}
}
