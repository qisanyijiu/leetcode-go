package medium

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	uf := NewUnUnionFind(m * n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row > 0 && grid[row][col] == grid[row-1][col] {
				if !uf.FindAndUnite(row*n+col, (row-1)*n+col) {
					return true
				}
			}
			if col > 0 && grid[row][col] == grid[row][col-1] {
				if !uf.FindAndUnite(row*n+col, row*n+col-1) {
					return true
				}
			}
		}
	}
	return false
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	size := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1
	}
	uf.parent = parent
	uf.size = size
	return uf
}

func (u *UnionFind) FindSet(x int) int {
	// 如果x的父节点不是x，则递归查找x的父节点，每一步更新
	if u.parent[x] != x {
		u.parent[x] = u.FindSet(u.parent[x])
	}
	// 返回x的父节点
	return u.parent[x]
}

func (u *UnionFind) Unite(x, y int) {
	// 查找x和y的父节点
	x, y = u.FindSet(x), u.FindSet(y)
	// 将size较大的节点作为父节点
	if u.size[x] < u.size[y] {
		u.parent[x] = y
		// 更新x的size
		u.size[y] += u.size[x]
	}else{
		u.parent[y] = x
		u.size[x] += u.size[y]
	}
	
}

func (u *UnionFind) FindAndUnite(x, y int) bool {
	// 查找x和y的父节点
	parentX := u.FindSet(x)
	// 查找y的父节点
	parentY := u.FindSet(y)
	if parentX != parentY {
		// 将x和y合并
		u.Unite(parentX, parentY)
		// 返回true，表示x和y在同一个集合中
		return true
	}
	// 返回false，表示x和y不在同一个集合中
	return false
}
