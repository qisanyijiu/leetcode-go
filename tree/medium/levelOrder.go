package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	var result = [][]int{}
	if root == nil {
		return result
	}

	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var level = []int{}
		var size = len(queue)
		var level_queue = []*TreeNode{}
		for i := 0; i < size; i++ {
			var node = queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				level_queue = append(level_queue, node.Left)
			}
			if node.Right != nil {
				level_queue = append(level_queue, node.Right)
			}
		}
		queue = level_queue
		result = append(result, level)
	}
	return result
}
