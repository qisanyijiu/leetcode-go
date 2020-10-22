package easy

import "testing"

func TestFindPairs(t *testing.T) {
	cases := []struct{
		nums []int
		k int
		res int
	}{
		struct {
			nums []int
			k    int
			res  int
		}{nums: []int{3, 1, 4, 1, 5}, k: 2, res: 2},
		struct {
			nums []int
			k    int
			res  int
		}{nums: []int{1, 2, 3, 4, 5}, k: 1, res: 4},
		struct {
			nums []int
			k    int
			res  int
		}{nums: []int{1, 3, 1, 5, 4}, k: 0, res: 1},
	}

	for _,item:=range cases{
		out := FindPairs(item.nums, item.k)
		if out != item.res{
			t.Errorf("Arr :%v, K: %d, Got: %d, Except: %d\n", item.nums, item.k, out, item.res)
		}
	}
}