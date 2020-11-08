package easy

func FirstUniqChar(s string) int {
	dict := [26]int{}
	for _, b := range([]byte(s)){
		dict[int(b) - int('a')] += 1
	}
	for i, b := range([]byte(s)){
		if dict[int(b) - int('a')] == 1 {
			return i
		}
	}
	return -1
}