package hard

import "testing"

func TestReverseKGroup(t *testing.T) {
	type Item struct {
		Input  *ListNode
		K int
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
			K: 2,
			Except: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  5,
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
			K: 3,
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
		out := ReverseKGroup(item.Input, item.K)
		if !out.equal(item.Except) {
			t.Errorf("input: %v, k: %d, except: %v, got: %v", item.Input.ToSlice(), item.K, item.Except.ToSlice(), out.ToSlice())
		}
	}
}
