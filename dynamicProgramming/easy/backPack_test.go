package easy

import "testing"

func TestBackPack(t *testing.T){
	cases := []struct{
		m int
		products []int
		res int
	}{
		struct {
			m        int
			products []int
			res      int
		}{m: 10, products: []int{1, 2, 3, 4}, res: 10},
		struct {
			m        int
			products []int
			res      int
		}{m: 9, products: []int{1, 2, 3, 4, 5}, res: 9},
	}
	for _, item := range cases{
		result := backPack(item.m, item.products)
		if result != item.res{
			t.Errorf("M: %d, P: %v, Except: %d, Got: %d\n", item.m, item.products, item.res, result)
		}
	}
}
