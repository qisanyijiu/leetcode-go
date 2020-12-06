package medium

import "testing"

func TestMyAtoi2(t *testing.T) {
	type Item struct {
		Input string
		Except int
	}
	cases := []Item{
		{
			Input:  "42",
			Except: 42,
		},
		{
			Input:  "   -42",
			Except: -42,
		},
		{
			Input:  "4193 with words",
			Except: 4193,
		},
		{
			Input:  "words and 987",
			Except: 0,
		},
		{
			Input:  "-91283472332",
			Except:  -2147483648,
		},
		{
			Input:  "+1",
			Except:  1,
		},
	}
	for _, item := range cases{
		out := MyAtoi(item.Input)
		if out != item.Except {
			t.Errorf("input: %s, except: %d, got: %d", item.Input, item.Except, out)
		}
	}
}