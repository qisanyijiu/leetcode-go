package easy

func RemoveDuplicates(S string) string {
	stack := []byte{}
	for i := 0; i < len(S); i ++ {
		c := S[i]
		index := len(stack) - 1
		if index < 0 || c != stack[index] {
			stack = append(stack, c)
		}else if c == stack[index] {
			stack = stack[:index]
		}
	}
	return string(stack)
}