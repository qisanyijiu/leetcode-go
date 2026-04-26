package medium

import "testing"

func Test_ContainsCycle(t *testing.T) {
	type TestCase struct {
		grid   [][]byte
		except bool
	}
	cases := []TestCase{
		TestCase{
			grid: [][]byte{
				[]byte{'a', 'a', 'a', 'a'},
				[]byte{'a', 'b', 'b', 'a'},
				[]byte{'a', 'b', 'b', 'a'},
				[]byte{'a', 'a', 'a', 'a'},
			},
			except: true,
		},
		TestCase{
			grid: [][]byte{
				[]byte{'a', 'a', 'a', 'b'},
				[]byte{'a', 'b', 'a', 'a'},
				[]byte{'a', 'a', 'b', 'a'},
				[]byte{'b', 'a', 'a', 'a'},
			},
			except: true,
		},
		TestCase{
			grid: [][]byte{
				[]byte{'a', 'a', 'a', 'b'},
				[]byte{'b', 'b', 'a', 'a'},
				[]byte{'a', 'a', 'b', 'a'},
				[]byte{'b', 'a', 'a', 'a'},
			},
			except: false,
		},
	}
	for idx, item := range cases {
		out := containsCycle(item.grid)
		if out != item.except {
			t.Errorf("test case [%d], except: %v, got: %v", idx, item.except, out)
		}
	}
}
