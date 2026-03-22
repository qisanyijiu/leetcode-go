package medium

import "math"

func isValidBST(root *TreeNode) bool {

	var helper func(root *TreeNode, lower, upper int) bool
	helper = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if root.Val <= lower || root.Val >= upper {
			return false
		}
		return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
	}
	return helper(root, math.MinInt, math.MaxInt)
}
