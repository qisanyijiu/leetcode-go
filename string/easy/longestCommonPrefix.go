package easy

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}
	if len(strs) == 1 {
		return strs[0]
	}
	length := len(strs[0])
	for _, str := range strs{
		if len(str) < length {
			length = len(str)
		}
	}
	if length == 0 {
		return prefix
	}
	Outer:
	for i := 0 ; i < length; i ++ {
		preByte := []byte(strs[0])[i]
		for j := 1; j < len(strs); j ++ {
			if []byte(strs[j])[i] != preByte {
				break Outer
			}
		}
		prefix += string(preByte)
	}
	return prefix
}