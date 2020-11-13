package hard

import (
	"testing"
)

func TestNQueen(t *testing.T){
	cases := []struct{
		n, r int
	}{
		struct{ n, r int }{n: 1, r: 1},
		struct{ n, r int }{n: 2, r: 0},
		struct{ n, r int }{n: 3, r: 0},
		struct{ n, r int }{n: 4, r: 2},
		struct{ n, r int }{n: 5, r: 10},
		struct{ n, r int }{n: 6, r: 4},
		struct{ n, r int }{n: 7, r: 40},
		struct{ n, r int }{n: 8, r: 92},
		struct{ n, r int }{n: 9, r: 352},
		struct{ n, r int }{n: 10, r: 724},
	}
	for _, item := range cases{
		result := len(nQueen(item.n))
		if result != item.r {
			t.Errorf("N: %d, Except: %d, Got: %d\n", item.n, item.r, result)
		}
	}
}