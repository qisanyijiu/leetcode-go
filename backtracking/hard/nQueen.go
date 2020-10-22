package hard

func nQueen(n int)[][]string{
	if n == 0 || n == 2 || n == 3{
		return [][]string{}
	}
	cols := make([]bool, n)
	// 记录\方向对角线的情况
	d1 := make([]bool, 2 * n)
	// 记录/防线对角线的情况
	d2 := make([]bool, 2 * n)

	board := make([]string, n)
	res := [][]string{}
	nQueenDfs(0, cols, d1, d2, board, &res)
	return res
}

func nQueenDfs(r int, cols, d1, d2 []bool, board []string, res *[][]string){
	if r == len(board){
		tmp := make([]string, len(board))
		copy(tmp, board)
		*res = append(*res, tmp)
		return
	}
	n := len(board)
	// 逐行列扫描
	for c := 0; c < len(board); c ++ {
		// 记录\方向对角线的情况
		id1 := r - c + n
		// 记录/方向对角线的情况
		id2 := 2 * n - r - c - 1
		// 确认列是否可以放入皇后
		if !cols[c] && !d1[id1] && !d2[id2] {
			// 构建一行
			b := make([]byte, n)
			for i := range b{
				b[i] = '.'
			}
			b[c] = 'Q'
			board[r] = string(b[:])
			// 标记占用
			cols[c], d1[id1], d2[id2] = true, true, true
			// 进入下一行
			nQueenDfs(r+1, cols, d1, d2, board, res)
			// 回溯，解除占用标记
			cols[c], d1[id1], d2[id2] = false, false, false
		}
	}
}