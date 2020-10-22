package easy

import (
	"strconv"
)

func CalPoints(ops []string) int {
	stack := make([]int, len(ops))
	point := -1
	for _, op := range ops{
		switch op {
		case "C":
			point--
		case "D":
			point++
			stack[point] = stack[point-1] * 2
		case "+":
			point++
			stack[point] = stack[point-1] + stack[point-2]
		default:
			points, _ := strconv.Atoi(op)
			point++
			stack[point] = points
		}
	}
	sum := 0
	for i := 0; i <= point; i ++ {
		sum += stack[i]
	}
	return sum
}