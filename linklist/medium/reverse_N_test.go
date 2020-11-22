package medium

import "testing"

func TestReverseN(t *testing.T) {
	type Item struct {
		Input  *ListNode
		N int
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
			N: 5,
			Except: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
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
			N: 3,
			Except: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
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
		},
	}
	for _, item := range cases {
		out := ReverseN(item.Input, item.N)
		if !out.equal(item.Except) {
			t.Errorf("input: %v, n: %d, except: %v, got: %v", item.Input.ToSlice(), item.N, item.Except.ToSlice(), out.ToSlice())
		}
	}
}