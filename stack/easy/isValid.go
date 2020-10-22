package easy

func IsValid(s string) bool {
	stack := make([]byte, len(s))
	point := -1
	push := func(c byte) {
		point ++
		stack[point] = c
	}
	for i := 0; i < len(s); i ++ {
		c := s[i]
		if point == -1 && (c == ')' || c == ']' || c == '}') {
			return false
		}
		switch c {
		case '(', '{', '[':
			push(c)
		case ')':
			if stack[point] == '(' {
				point --
			} else {
				return false
			}
		case '}':
			if stack[point] == '{' {
				point --
			} else {
				return false
			}
		case ']':
			if stack[point] == '[' {
				point --
			} else {
				return false
			}
		default:
			break
		}
	}
	return point == -1
}