package easy

import "testing"

func TestMinimumTotal(t *testing.T){
	cases := []struct{
		arr [][]int
		res int
	}{
		struct {
			arr [][]int
			res int
		}{arr: [][]int{
			[]int{2},
			[]int{3, 4},
			[]int{6, 5, 7},
			[]int{4, 1, 8, 3},
		}, res: 11},
	}
	for _, item := range cases{
		result := minimumTotal(item.arr)
		if result != item.res{
			t.Errorf("Except: %d, Got: %d\n", item.res, result)
		}
	}
}