package easy

import "testing"

func TestMajorityElement(t *testing.T) {
	type Item struct {
		Nums []int
		Except int
	}
	cases := []Item{
		{
			Nums:   []int{3,2,3},
			Except: 3,
		},
		{
			Nums:   []int{2,2,1,1,1,2,2},
			Except: 2,
		},
	}
	for _, item := range cases{
		out := MajorityElement(item.Nums)
		if out != item.Except {
			t.Errorf("input: %v, except: %d, got: %d", item.Nums, item.Except, out)
		}
	}
}