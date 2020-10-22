package easy

func PlusOne(digits []int) []int {
	ten := 1
	index := len(digits) - 1
 	for ten == 1 && index >= 0 {
 		tmp := digits[index]
		digits[index] = (tmp + ten) % 10
 		ten = (tmp + 1) / 10
 		index --
	}
	if ten == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
