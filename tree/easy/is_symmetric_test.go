package easy

import "testing"

func TestIsSymmetric(t *testing.T) {
	type Item struct {
		Tree        *TreeNode
		IsSymmetric bool
	}
	cases := []Item{
		{
			Tree:        &TreeNode{
				Val:   1,
				Left:  &TreeNode{
					Val:   2,
					Left:  &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			IsSymmetric: true,
		},
		{
			Tree:        &TreeNode{
				Val:   1,
				Left:  &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			IsSymmetric: false,
		},
	}
	for _, item := range cases {
		out := IsSymmetric(item.Tree)
		if out != item.IsSymmetric {
			t.Errorf("input: %v, except: %t, got: %t", item.Tree, item.IsSymmetric, out)
		}
	}
}
