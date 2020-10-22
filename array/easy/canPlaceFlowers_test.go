package easy

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	cases := []struct{
		nums []int
		n int
		res bool
	}{
		struct {
			nums []int
			n    int
			res  bool
		}{nums: []int{1,0,0,0,1}, n: 1, res: true},
		struct {
			nums []int
			n    int
			res  bool
		}{nums: []int{1,0,0,0,1}, n: 2, res: false},
		struct {
			nums []int
			n    int
			res  bool
		}{nums: []int{0, 1, 0, 0, 1, 0, 0, 1, 0}, n: 1, res: false},
	}
	for _, item := range cases{
		out := CanPlaceFlowers(item.nums,item.n)
		if out != item.res {
			t.Errorf("Arr: %v, N: %d, Got: %t, Except: %t\n", item.nums, item.n, out, item.res)
		}
	}
}
