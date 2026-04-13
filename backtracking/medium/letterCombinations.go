package medium

func letterCombinations(digits string) []string {
	var ans = make([]string, 0)
	dfsLetter(0, []byte(digits), "", &ans)
	return ans
}

func dfsLetter(depth int, data []byte, pre string, ans *[]string) {
	if depth == len(data) {
		*ans = append(*ans, pre)
		return
	}
	n2l := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	letters := n2l[data[depth]]
	for _, letter := range letters {
		dfsLetter(depth+1, data, pre+string(letter), ans)
	}
}
