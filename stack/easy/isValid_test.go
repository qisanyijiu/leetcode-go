package easy

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct{
		input string
		res bool
	}{
		struct {
			input string
			res   bool
		}{input: "()", res: true},
		struct {
			input string
			res   bool
		}{input: "()[]{}", res: true},
		struct {
			input string
			res   bool
		}{input: "(]", res: false},
		struct {
			input string
			res   bool
		}{input: "([)]", res: false},
		struct {
			input string
			res   bool
		}{input: "{[]}", res: true},
	}

	for _, item := range cases{
		result := IsValid(item.input)
		if result != item.res{
			t.Errorf("Input: %s, Except: %t, Got: %t\n", item.input, item.res, result)
		}
	}
}
