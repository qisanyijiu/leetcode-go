package easy

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result = []int{root.Val}
	result = append(result, PreorderTraversal(root.Left)...)
	result = append(result, PreorderTraversal(root.Right)...)
	return result
}
