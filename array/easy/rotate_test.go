package easy

import "testing"

func TestRotate(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		output []int
	}{
		struct {
			nums   []int
			k      int
			output []int
		}{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 1, output: []int{7, 1, 2, 3, 4, 5, 6}},
		struct {
			nums   []int
			k      int
			output []int
		}{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 2, output: []int{6, 7, 1, 2, 3, 4, 5}},
		struct {
			nums   []int
			k      int
			output []int
		}{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, output: []int{5, 6, 7, 1, 2, 3, 4}},
		struct {
			nums   []int
			k      int
			output []int
		}{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 7, output: []int{1, 2, 3, 4, 5, 6, 7}},
		struct {
			nums   []int
			k      int
			output []int
		}{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 8, output: []int{7, 1, 2, 3, 4, 5, 6}},
	}
	for _, item := range cases {
		out := item.nums
		Rotate(out, item.k)
		if !ArrEqual(item.nums, item.output) {
			t.Errorf("Nums: %v, K: %d, Got: %v, Except: %v\n", item.nums, item.k, out, item.output)
		}
	}
}
