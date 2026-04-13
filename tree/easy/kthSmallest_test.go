package easy

import "testing"

func Test_KthSmallest(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	got := kthSmallest(tree, 1)
	if got != 1 {
		t.Errorf("got %d, want %d", got, 1)
	}
}
