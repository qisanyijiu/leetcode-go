package easy

import "testing"

func TestBackspaceCompare(t *testing.T) {
	cases := []struct{
		a string
		b string
		result bool
	}{
		struct {
			a      string
			b      string
			result bool
		}{a: "ab#c", b: "ad#c", result: true},
		struct {
			a      string
			b      string
			result bool
		}{a: "ab##", b: "c#d#", result: true},
		struct {
			a      string
			b      string
			result bool
		}{a: "a##c", b: "#a#c", result: true},
		struct {
			a      string
			b      string
			result bool
		}{a: "a#c", b: "b", result: false},
	}
	for _, item := range cases{
		res := BackspaceCompare(item.a, item.b)
		if res != item.result{
			t.Errorf("A: %s, B: %s, Except: %t, Got: %t", item.a, item.b, item.result, res)
		}
	}
}