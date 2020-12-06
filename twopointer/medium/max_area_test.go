package medium

import "testing"

func TestMaxArea(t *testing.T) {
	type Item struct {
		Input []int
		Except int
	}
	cases := []Item{
		{
			Input:  []int{1,8,6,2,5,4,8,3,7},
			Except: 49,
		},
		{
			Input:  []int{1,1},
			Except: 1,
		},
		{
			Input:  []int{4,3,2,1,4},
			Except: 16,
		},
		{
			Input:  []int{1,2,1},
			Except: 2,
		},
	}
	for _, item := range cases{
		out := MaxArea(item.Input)
		if out != item.Except {
			t.Errorf("height: %v, except: %d, got: %d", item.Input, item.Except, out)
		}
	}
}
