package easy

func HammingWeight(num uint32) int {
	var sum = 0
	for num != 0 {
		num &= (num-1)
		sum ++
	}
	return sum
}