package easy

/*
*
101. 对称二叉树

https://leetcode-cn.com/problems/symmetric-tree/

给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

	   1
	  / \
	 2   2
	/ \ / \

3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

	  1
	 / \
	2   2
	 \   \
	 3    3
*/
func IsSymmetric(root *TreeNode) bool {
	left, right := root, root
	q := []*TreeNode{}
	q = append(q, left)
	q = append(q, right)
	for len(q) > 0 {
		// 取出左右两个节点
		left, right = q[0], q[1]
		// 删除两个节点
		q = q[2:]
		// 如果两个节点都为空，则继续
		if left == nil && right == nil {
			continue
		}
		// 如果两个节点有一个为空，则返回false
		if left == nil || right == nil {
			return false
		}
		// 如果两个节点值不相等，则返回false
		if left.Val != right.Val {
			return false
		}
		// 将两个节点的左右子节点加入队列
		q = append(q, left.Left)
		q = append(q, right.Right)
		// 将两个节点的右子节点加入队列
		q = append(q, left.Right)
		q = append(q, right.Left)
	}
	return true
}
