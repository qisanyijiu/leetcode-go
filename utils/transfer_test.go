package utils

import "testing"

func TestTransfer(t *testing.T) {
	type Item struct{
		Input int
		Except string
	}
	cases := []Item{
		Item{
			Input: 1,
			Except: "A",
		},
		Item{
			Input: 2,
			Except: "B",
		},
		Item{
			Input: 26,
			Except: "Z",
		},
		Item{
			Input: 27,
			Except: "AA",
		},
		Item{
			Input: 28,
			Except: "AB",
		},
		Item{
			Input: 52,
			Except: "AZ",
		},
		Item{
			Input: 53,
			Except: "BA",
		},
	}
	for _, item := range cases{
		out := Transfer(item.Input)
		if out != item.Except {
			t.Errorf("input: %d, except: %s, got: %s \n", item.Input, item.Except, out)
		}
	}
}

func TestTransfer2(t *testing.T) {
	Transfer()
}
