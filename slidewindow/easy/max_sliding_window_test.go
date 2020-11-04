package easy

import (
	"leetcode/utils"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	type Item struct {
		Nums []int
		K int
		Res []int
	}
	cases := []Item{
		{
			Nums: []int{},
			K: 3,
			Res: []int{},
		},
	}
	for _,item := range cases {
		out := MaxSlidingWindow(item.Nums, item.K)
		if !utils.EqualIntSlice(out, item.Res) {
			t.Errorf("nums: %v, k: %d, except: %v, got: %v", item.Nums, item.K, item.Res, out)
		}
	}
}