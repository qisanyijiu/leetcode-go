package easy

func DiameterOfBinaryTree(root *TreeNode) int {
	var result = 1
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		result = max(result, left+right+1)
		return max(left, right) + 1
	}
	depth(root)
	return result - 1
}
