package medium

import "testing"

func TestThreeSum(t *testing.T) {
	sum := func(nums []int) int {
		ans := 0
		for _, num := range nums{
			ans += num
		}
		return ans
	}
	var data []int
	data = []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(data)
	if len(result) != 2 {
		t.Fail()
	}
	for _, record := range result{
		if sum(record) != 0{
			t.Fail()
		}
	}
}

func BenchmarkThreeSum(b *testing.B){
	var data []int
	data = []int{-1, 0, 1, 2, -1, -4}
	b.ResetTimer()
	for i := 0; i< b.N; i ++ {
		threeSum(data)
	}
}