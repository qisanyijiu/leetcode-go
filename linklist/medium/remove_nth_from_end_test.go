package medium

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		input  *ListNode
		n      int
		output *ListNode
	}{
		struct {
			input  *ListNode
			n      int
			output *ListNode
		}{
			n: 2,
			input: &ListNode{
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
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
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
		struct {
			input  *ListNode
			n      int
			output *ListNode
		}{
			n: 3,
			input: &ListNode{
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
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
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
		struct {
			input  *ListNode
			n      int
			output *ListNode
		}{
			n: 5,
			input: &ListNode{
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
			output: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		struct {
			input  *ListNode
			n      int
			output *ListNode
		}{
			n: 6,
			input: &ListNode{
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
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{
							Val:  4,
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

	for i, item := range cases {
		if i < 2 {
			continue
		}
		res := removeNthFromEnd(item.input, item.n)
		if !res.equal(item.output) {
			t.Errorf("except: %v, got: %v", item.output.ToSlice(), res.ToSlice())
		}

	}
}
