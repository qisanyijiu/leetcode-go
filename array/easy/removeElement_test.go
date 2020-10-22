package easy

import "testing"

func TestRemoveElement(t *testing.T) {
	cases := []struct{
		nums []int
		val, res int
	}{
		struct {
			nums     []int
			val, res int
		}{nums: []int{3,2,2,3}, val: 3, res: 2},
		struct {
			nums     []int
			val, res int
		}{nums: []int{0,1,2,2,3,0,4,2}, val: 2, res: 5},
	}
	for _, item := range cases{
		result := RemoveElement(item.nums, item.val)
		if result != item.res{
			t.Errorf("Nums: %v, Val: %d, Except: %d, Got: %d\n", item.nums, item.val, item.res, result)
		}
	}
}
