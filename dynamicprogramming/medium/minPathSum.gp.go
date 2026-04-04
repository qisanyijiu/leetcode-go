package medium

import "math"

func MinPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 && j > 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if i > 0 && j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[row-1][col-1]
}

func min(arr ...int) int {
	cur := math.MaxInt64
	for i := 0; i < len(arr); i++ {
		if cur > arr[i] {
			cur = arr[i]
		}
	}
	return cur
}
