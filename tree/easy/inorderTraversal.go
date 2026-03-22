package easy

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result = InorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, InorderTraversal(root.Right)...)
	return result
}
