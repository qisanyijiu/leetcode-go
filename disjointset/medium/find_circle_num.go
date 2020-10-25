package medium

/**
班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。

给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。

 

示例 1：

输入：
[[1,1,0],
 [1,1,0],
 [0,0,1]]
输出：2
解释：已知学生 0 和学生 1 互为朋友，他们在一个朋友圈。
第2个学生自己在一个朋友圈。所以返回 2 。
示例 2：

输入：
[[1,1,0],
 [1,1,1],
 [0,1,1]]
输出：1
解释：已知学生 0 和学生 1 互为朋友，学生 1 和学生 2 互为朋友，所以学生 0 和学生 2 也是朋友，所以他们三个在一个朋友圈，返回 1 。

 */
func FindCircleNum(M [][]int) int {
	// 生成根节点记录表
	parent := make([]int, len(M))
	rank := make([]int,len(M))
	for i := range parent{
		parent[i] = -1
		rank[i] = 0
	}
	for i, row := range M{
		for j ,col := range row{
			if col == 1 && i != j {
				circleNumUnion(parent, i, j, rank)
			}
		}
	}
	count := 0
	for _, v := range parent {
		if v == - 1 {
			count ++
		}
	}
	return count
}

// 寻找根节点
func circleNumFind(parent []int, i int) int {
	root := i
	for parent[root] != - 1{
		root = parent[root]
	}
	return root
}

// 合并
func circleNumUnion(parent []int, x int, y int, rank []int){
	// 查找x的根
	xSet := circleNumFind(parent, x)
	// 查找y的根
	ySet := circleNumFind(parent, y)
	// 如果x的根不等于y的根合并，否则说明属于同一颗树，不需要合并
	if xSet == ySet {
		return
	}

	if rank[xSet] > rank[ySet] {
		parent[ySet] = xSet
	} else if rank[ySet] > rank[xSet]{
		parent[xSet] = ySet
	} else {
		parent[xSet] = ySet
		rank[ySet] ++
	}
}