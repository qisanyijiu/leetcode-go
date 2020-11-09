package easy

import "testing"

func TestFindMagicIndex(t *testing.T) {
	type Item struct {
		nums  []int
		index int
	}
	cases := []Item{
		{
			nums:  []int{0, 2, 3, 4, 5},
			index: 0,
		},
		{
			nums:  []int{1, 1, 1},
			index: 1,
		},
		{
			nums:  []int{1, 1, 1, 4, 4, 4, 4},
			index: 1,
		},
		{
			nums:  []int{1, 2, 3, 4, 5, 6, 7},
			index: -1,
		},
	}
	for _, item := range cases {
		out := FindMagicIndex(item.nums)
		if out != item.index {
			t.Errorf("nums: %d, except: %d, got: %d", item.nums, item.index, out)
		}
	}
}
