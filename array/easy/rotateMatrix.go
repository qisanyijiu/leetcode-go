package easy

func RotateMatrix(matrix [][]int)  {
	n := len(matrix) / 2 + len(matrix) % 2
	for y := 0 ; y < n; y ++ {
		for x := 0; x < n ; x ++ {
			matrix[y][x], matrix[y][n-x-1], matrix[n-y-1][n-x-1], matrix[n-y-1][x] = matrix[n-y-1][x], matrix[y][x], matrix[y][n-x-1], matrix[n-y-1][n-x-1]
		}
	}
}

/**

1 2 3
4 5 6
7 8 9


 */
