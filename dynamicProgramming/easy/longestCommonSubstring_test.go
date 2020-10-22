package easy

import "testing"

func TestLongestCommonSubstring(t *testing.T){
	cases := []struct{
		a, b string
		res int
	}{
		struct {
			a, b string
			res  int
		}{a: "abcd", b: "ebcf", res: 2},
		struct {
			a, b string
			res  int
		}{a: "abcd", b: "ghjk", res: 0},
		struct {
			a, b string
			res  int
		}{a: "", b: "ebcf", res: 0},
		struct {
			a, b string
			res  int
		}{a: "fuck", b: "", res: 0},
		struct {
			a, b string
			res  int
		}{a: "fuck", b: "fuckyou", res: 4},
	}
	for _, item := range cases{
		result := longestCommonSubstring(item.a, item.b)
		if result != item.res{
			t.Errorf("A: %s, B: %s, Except: %d, Got: %d\n", item.a, item.b, item.res, result)
		}
	}
}
