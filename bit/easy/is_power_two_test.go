package easy

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	cases := []struct {
		num int
		out bool
	}{
		{num: 1, out: true},
		{num: 16, out: true},
		{num: 255, out: false},
	}
	for _, item := range cases {
		res := IsPowerOfTwo(item.num)
		if res != item.out {
			t.Errorf("input: %d, except: %t, got: %t", item.num, item.out, res)
		}
	}

}
