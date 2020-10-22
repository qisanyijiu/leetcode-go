package medium

func combine(n, k int) [][]int{
	res := [][]int{}
	if n <= 0 || k <= 0 {
		return res
	}
	generateCombine(n, k, 1, []int{}, &res)
	return res
}

func generateCombine(n, k, start int, c []int, res *[][]int){
	if len(c) == k{
		*res = append(*res, c)
		return
	}
	for i := start; i <= n ; i ++ {
		c = append(append([]int{}, c...), i)
		generateCombine(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}