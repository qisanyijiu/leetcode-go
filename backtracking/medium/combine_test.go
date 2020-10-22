package medium

import "testing"

func TestCombine(t *testing.T){
	cases := []struct{
		n, k, l int
	}{
		struct{ n, k, l int }{n: 4, k: 2, l: 6},
		struct{ n, k, l int }{n: 4, k: 3, l: 4},
	}
	for _, item :=range cases{
		result := len(combine(item.n, item.k))
		if result != item.l{
			t.Errorf("N: %d, K: %d, Except: %d, Got: %d\n", item.n, item.k, item.l, item.k)
		}
	}
}
