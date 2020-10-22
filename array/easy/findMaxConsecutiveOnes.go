package easy

func FindMaxConsecutiveOnes(nums []int) int {
	max := 0
	tmp := 0
	for _, num := range nums{
		if num == 1{
			tmp += 1
		}else{
			if tmp > max{
				max = tmp
			}
			tmp = 0
		}
	}
	if tmp > max{
		max = tmp
	}
	return max
}