package easy

/**
108. 将有序数组转换为二叉搜索树

https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/

将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
 */

func SortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums) - 1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid - 1)
	root.Right = helper(nums, mid + 1, right)
	return root
}