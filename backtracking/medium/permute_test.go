package medium

import "testing"

func TestPermute(t *testing.T) {
	data := []int{1, 2, 3}
	if len(Permute(data)) != 6 {
		t.Fail()
	}
}
