package easy

/**
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	deep := 0
	if root != nil {
		deep += 1
		leftDeep := MaxDepth(root.Left)
		rightDepp := MaxDepth(root.Right)
		if leftDeep > rightDepp {
			deep += leftDeep
		}else{
			deep += rightDepp
		}
	}
	return deep
}
