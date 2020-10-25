package medium

import "testing"

func TestFindCircleNum(t *testing.T) {
	cases := []struct {
		Input [][]int
		Res   int
	}{
		{Input: [][]int{
			[]int{1,1,0},
			[]int{1,1,0},
			[]int{0,0,1},
		}, Res: 2},
		{Input: [][]int{
			[]int{1,1,0},
			[]int{1,1,1},
			[]int{0,0,1},
		}, Res: 1},
	}
	for _, item := range cases{
		out := FindCircleNum(item.Input)
		if out != item.Res {
			t.Errorf("input: %v, except: %d, got: %d", item.Input, item.Res, out)
		}
	}

}
