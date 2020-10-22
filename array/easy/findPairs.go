package easy

func FindPairs(nums []int, k int) int {
	cnt := 0
	flag := map[int]int{}
	for _, v := range nums{
		flag[v]++
	}
	for i, v := range flag{
		_, ok := flag[i+k]
		if ok && k > 0 || k==0 && v > 1{
			cnt ++
		}
	}
	return cnt
}
