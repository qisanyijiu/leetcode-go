package easy

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result = PostorderTraversal(root.Left)
	result = append(result, PostorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}
