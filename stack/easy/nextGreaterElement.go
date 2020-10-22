package easy

/**
递减栈加hash表
 */
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, len(nums2))
	point := -1
	nextDict := map[int]int{}
	for _, item := range nums2{
		// 空栈直接压栈
		if point == -1 || item < stack[point]{
			point ++
			stack[point] = item
		} else {
			for point > -1 && item > stack[point] {
				top := stack[point]
				nextDict[top] = item
				point --
			}
			point++
			stack[point] = item
		}
	}
	out := make([]int, len(nums1))
	for index, num := range nums1{
		if v, ok := nextDict[num]; ok {
			out[index] = v
		} else {
			out[index] = -1
		}
	}
	return out
}
