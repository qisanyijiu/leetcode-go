package medium

import "testing"

func TestThreeSumClosest(t *testing.T) {
	var data []int
	data = []int{-1,2,1,4}
	target := 1
	if threeSumClosest(data, target) != 2 {
		t.Fail()
	}
}

func BenchmarkThreeSumClosest(b *testing.B){
	var data []int
	data = []int{-1,2,1,4}
	b.ResetTimer()
	for i := 0; i< b.N; i ++ {
		threeSumClosest(data, 1)
	}
}