package easy

import "testing"

func Test_DecodeString(t *testing.T) {
	type Case struct {
		input  string
		expect string
	}
	cases := []Case{
		{"abc", "abc"},
		{"3[a2[c]]", "accaccacc"},
		{"abc3[cd]xyz", "abccdcdcdxyz"},
	}
	for _, c := range cases {
		got := decodeString(c.input)
		if got != c.expect {
			t.Errorf("input: %v, expect: %v, got: %v", c.input, c.expect, got)
		}
	}
}
