package easy

import "testing"

func TestInorderTraversal(t *testing.T) {
	cases := []struct {
		root *TreeNode
		res  []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			res: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, item := range cases {
		out := InorderTraversal(item.root)
		if !ArrEqual(out, item.res) {
			t.Errorf("input: %v, except: %v, got: %v", item.root, item.res, out)
		}
	}
}

func ArrEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
