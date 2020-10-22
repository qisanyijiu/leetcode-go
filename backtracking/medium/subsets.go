package medium

func subsets(nums []int) [][]int {
	res := make([][]int, 1)
	for _, v := range nums {
		for _, r := range res {
			res = append(res, append([]int{v}, r...))
		}
	}
	return res
}