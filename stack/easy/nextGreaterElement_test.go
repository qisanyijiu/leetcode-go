package easy

import "testing"

func TestNextGreaterElement(t *testing.T) {
	cases := []struct{
		a, b, output []int
	}{
		struct{ a, b, output []int }{a: []int{4,1,2}, b: []int{1,3,4,2}, output: []int{-1,3,-1}},
		struct{ a, b, output []int }{a: []int{2,4}, b: []int{1,2,3,4}, output: []int{3,-1}},
	}
	for _, item := range cases{
		result := NextGreaterElement(item.a, item.b)
		if !equalArr(result, item.output){
			t.Errorf("A: %v, B: %v, Except: %v, Got: %v\n", item.a, item.b, item.output, result)
		}
	}
}

func equalArr(a, b []int)bool{
	if len(a) != len(b){
		return false
	}
	for i := 0; i < len(a); i ++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
