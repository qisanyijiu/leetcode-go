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
	if u.parent[x] != x {
		u.parent[x] = u.FindSet(u.parent[x])
	}
	return u.parent[x]
}

func (u *UnionFind) Unite(x, y int) {
	x, y = u.FindSet(x), u.FindSet(y)
	if u.size[x] < u.size[y] {
		x, y = y, x
	}
	u.parent[y] = x
	u.size[x] += u.size[y]
}

func (u *UnionFind) FindAndUnite(x, y int) bool {
	parentX := u.FindSet(x)
	parentY := u.FindSet(y)
	if parentX != parentY {
		u.Unite(parentX, parentY)
		return true
	}
	return false
}
