package medium

import "testing"

func TestReverseBetween(t *testing.T) {
	type Item struct {
		Input  *ListNode
		M      int
		N      int
		Except *ListNode
	}

	cases := []Item{
		{
			Input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			M: 2,
			N: 4,
			Except: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	for _, item := range cases {
		out := ReverseBetween(item.Input, item.M, item.N)
		if !out.equal(item.Except) {
			t.Errorf("input: %v, m: %d, n: %d, except: %v, got: %v", item.Input.ToSlice(), item.M, item.N, item.Except.ToSlice(), out.ToSlice())
		}
	}
}
