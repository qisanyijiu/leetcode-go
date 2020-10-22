package easy

func ReverseString(s []byte)  {
	length := len(s)
	for i := 0;i < len(s) / 2; i ++ {
		s[i], s[length - i - 1] = s[length - i - 1], s[i]
	}
}
