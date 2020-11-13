package hard

/**
37. 解数独

https://leetcode-cn.com/problems/sudoku-solver/

编写一个程序，通过填充空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/
func SolveSudoku(board [][]byte) {
	sudoKuBacktracking(board, 0, 0)
}

func sudoKuBacktracking(board [][]byte, i, j int) bool {
	m, n := 9, 9
	if j == n {
		// 穷举到最后一列的话就换到下一行重新开始。
		return sudoKuBacktracking(board, i+1, 0)
	}
	if i == m {
		return true
	}
	// 如果该位置是预设的数字，不用我们操心
	if board[i][j] != '.' {
		return sudoKuBacktracking(board, i, j+1)
	}

	for ch := '1'; ch <= '9'; ch++ {
		if !sudoKuIsValid(board, i, j, byte(ch)) {
			continue
		}
		board[i][j] = byte(ch)
		if sudoKuBacktracking(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func sudoKuIsValid(board [][]byte, r, c int, n byte) bool {
	for i := 0; i < 9; i++ {
		// 判断行是否存在重复
		if board[r][i] == n {
			return false
		}
		// 判断列是否存在
		if board[i][c] == n {
			return false
		}
		// 判断3 * 3的格子内是否存在
		var cr = (r/3)*3 + i/3
		var cc = (c/3)*3 + i%3
		if board[cr][cc] == n {
			return false
		}
	}
	return true
}
