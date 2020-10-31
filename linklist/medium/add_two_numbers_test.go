package medium

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		numA *ListNode
		numB *ListNode
		res  *ListNode
	}{
		{
			numA: &ListNode{
				Val:  1,
				Next: nil,
			},
			numB: &ListNode{
				Val:  1,
				Next: nil,
			},
			res: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
		{
			numA: &ListNode{
				Val:  9,
				Next: nil,
			},
			numB: &ListNode{
				Val:  1,
				Next: nil,
			},
			res: &ListNode{
				Val:  0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
		{
			numA: &ListNode{
				Val:  0,
				Next: &ListNode{
					Val:  9,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
			numB: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			res: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{
						Val:  0,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, item := range cases {
		out := AddTwoNumbers(item.numA, item.numB)
		if !item.res.equal(out) {
			outArr := out.ToSlice()
			exceptArr := item.res.ToSlice()
			t.Errorf("except: %v, got: %v", exceptArr, outArr)
		}

	}

}
