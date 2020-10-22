package medium

func numsIsLands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0{
		return 0
	}
	cols := len(grid[0])
	if cols == 0{
		return 0
	}
	// 标记，回溯
	visited := make([]*[]bool, rows)
	for i := 0; i < rows; i ++ {
		tmp := make([]bool,cols)
		visited[i] = &tmp
	}
	res := 0
	for i := 0; i < rows; i ++ {
		for j := 0; j < cols; j ++ {
			// 是陆地并且没有被参观过的地方递归参观
			if grid[i][j] == '1' && !(*visited[i])[j]{
				dfsNumIsLands(grid, i, j, rows, cols, &visited)
				res ++
			}
		}
	}

	return res
}

func dfsNumIsLands(grid [][]byte, y, x int, rows, cols int, visited *[]*[]bool){
	(*(*visited)[y])[x] = true
	d := [2][2]int{
		[2]int{0,  1},
		[2]int{1,  0},
	}
	for i := 0; i < 2 ; i ++ {
		newy := y + d[i][1]
		newx := x + d[i][0]
		// 棋盘内参观过的所有地方全部标记
		if inArea(newx, newy, rows, cols) && !(*(*visited)[newy])[newx] && grid[newy][newx] == '1'{
			dfsNumIsLands(grid, newy, newx, rows, cols, visited)
		}
	}
}

func inArea(x, y int, rows, cols int) bool {
	return x >= 0 && y>=0 && x < cols && y < rows
}