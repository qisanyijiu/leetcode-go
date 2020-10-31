package medium

import "testing"

func TestLongestPalindromeDP(t *testing.T) {
	cases := []struct{
		input string
		output string
	}{
		{input: "cabad", output: "aba"},
		{input: "abad", output: "aba"},
		{input: "cbbd", output: "bb"},
	}
	for _, item := range cases{
		res := LongestPalindromeDP(item.input)
		if res != item.output {
			t.Errorf("input: %s, except: %s, got: %s", item.input, item.output, res)
		}
	}
}

func TestLongestPalindromeMID(t *testing.T) {
	cases := []struct{
		input string
		output string
	}{
		{input: "cabad", output: "aba"},
		{input: "abad", output: "aba"},
		{input: "cbbd", output: "bb"},
	}
	for _, item := range cases{
		res := LongestPalindromeMID(item.input)
		if res != item.output {
			t.Errorf("input: %s, except: %s, got: %s", item.input, item.output, res)
		}
	}
}

func TestLongestPalindromeManacher(t *testing.T) {
	cases := []struct{
		input string
		output string
	}{
		{input: "cabad", output: "aba"},
		{input: "abad", output: "aba"},
		{input: "cbbd", output: "bb"},
	}
	for _, item := range cases{
		res := LongestPalindromeManacher(item.input)
		if res != item.output {
			t.Errorf("input: %s, except: %s, got: %s", item.input, item.output, res)
		}
	}
}