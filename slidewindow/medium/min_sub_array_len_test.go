package medium

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	type Item struct {
		S      int
		Nums   []int
		Except int
	}
	cases := []Item{
		{
			S:    7,
			Nums: []int{2, 3, 1, 2, 4, 3},
			Except: 2,
		},
		{
			S: 15,
			Nums: []int{5,1,3,5,10,7,4,9,2,8},
			Except: 2,
		},
	}
	for _, item := range cases {
		res := MinSubArrayLen(item.S, item.Nums)
		if res != item.Except{
			t.Errorf("sum: %d, nums: %v, except: %d, got: %d", item.S, item.Nums, item.Except, res)
		}
	}
}
