package easy

import "testing"

func TestMinArray(t *testing.T) {
	type Item struct {
		Nums   []int
		Except int
	}
	cases := []Item{
		{
			Nums:   []int{3, 4, 5, 1, 2},
			Except: 1,
		},
		{
			Nums:   []int{2, 2, 2, 0, 1},
			Except: 0,
		},
	}
	for _, item := range cases {
		out := MinArray(item.Nums)
		if out != item.Except {
			t.Errorf("nums: %v, except: %d, got: %d", item.Nums,item.Except, out)
		}
	}
}
