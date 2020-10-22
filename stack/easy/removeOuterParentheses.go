package easy

import "bytes"

func RemoveOuterParentheses(S string) string {
	var res bytes.Buffer
	counter := 0
	for _, c := range S{
		// 写入条件
		// counter != 0
		// counter = 1且要写入的为)
		if counter != 0 && !(counter == 1 && c == ')'){
			res.WriteRune(c)
		}
		if c == '('{
			counter ++
		}else{
			counter --
		}
	}
	return res.String()
}