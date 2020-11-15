package medium

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	freq := make([]int, 26)
	maxFreq := 0
	for _, task := range tasks{
		index := int(task)-int('A')
		freq[index] += 1
		if freq[index] > maxFreq {
			maxFreq = freq[index]
		}
	}
	count := 0
	for i := 0; i < len(freq); i ++ {
		if freq[i] == maxFreq {
			count ++
		}
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	return max(len(tasks), (maxFreq-1) * (n+1) + count)
}