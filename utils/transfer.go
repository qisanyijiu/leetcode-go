package utils

func Transfer(num int) string {
	var ans = []byte{}
	for num > 0 {
		num = num - 1
		tmp := num % 26
		ans = append([]byte{byte(int('A') + tmp)}, ans...)
		num = num / 26
	}
	return string(ans)
}

