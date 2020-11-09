package medium

import "testing"

func TestThreeSumClosest(t *testing.T) {
	type Item struct {
		Nums []int
		Target int
		Res int
	}
	cases := []Item{
		{
			Nums: []int{-1,2,1,-4},
			Target: 1,
			Res: 2,
		},
		{
			Nums: []int{0,2,1,-3},
			Target: 1,
			Res: 0,
		},
	}
	for _, item := range cases{
		out := ThreeSumClosest(item.Nums,item.Target)
		if out != item.Res {
			t.Errorf("nums: %v, target: %d, except: %d, got: %d", item.Nums, item.Target, item.Res, out)
		}
	}
}