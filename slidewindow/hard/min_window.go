package hard

func MinWindow(s string, t string) string {
	var need = map[byte]int{}
	var window = map[byte]int{}
	for _, c := range []byte(t) {
		need[c] += 1
	}
	left, right := 0, 0
	valid := 0
	start, length := 0, len(s) + 1

	for right < len(s) {
		c := s[right]
		right ++
		if _, ok := need[c]; ok {
			window[c] ++
			if window[c] == need[c] {
				valid ++
			}
		}

		for valid == len(need) {
			if right - left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left ++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid --
				}
				window[d] --
			}
		}
	}

	if  length == len(s) + 1 {
		return ""
	}else{
		return s[start:start+length]
	}
}