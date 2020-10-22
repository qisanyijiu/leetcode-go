package easy

func BackspaceCompare(S string, T string) bool {

	sa := gen(S)
	st := gen(T)
	// 判断是否相等
	if  sa != st{
		return false
	}
	return true
}

func gen(S string) string {
	stack := make([]byte, len(S))
	point := -1
	for i := 0; i < len(S); i ++ {
		c := S[i]
		if c != '#'{
			point++
			stack[point] = c
		}else{
			if point > -1{
				point--
			}
		}
	}
	return string(stack[:point+1])
}