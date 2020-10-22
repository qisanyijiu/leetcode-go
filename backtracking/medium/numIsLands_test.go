package medium

import "testing"

func TestNumsIsLands(t *testing.T){
	cases := []struct{
		grad [][]byte
		nums int
	}{
		struct {
			grad [][]byte
			nums int
		}{grad: [][]byte{
			[]byte("11110"),
			[]byte("11010"),
			[]byte("11000"),
			[]byte("00000"),
		}, nums: 1},
		struct {
			grad [][]byte
			nums int
		}{grad: [][]byte{
			[]byte("11000"),
			[]byte("11000"),
			[]byte("00100"),
			[]byte("00011"),
		}, nums: 3},
	}

	for _, item := range cases{
		res := numsIsLands(item.grad)
		if res != item.nums{
			t.Errorf("Grad: %v, Except: %d, Got: %d\n", item.grad, item.nums, res)
		}
	}
}