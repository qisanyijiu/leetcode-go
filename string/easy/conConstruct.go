package easy

func conConstruct(ransomNote string, magazine string) bool {
	mc := getCount(magazine)
	for _, b := range ransomNote{
		mc[b-'a'] --
		if mc[b-'a'] < 0 {
			return false
		}
	}
	return true
}

/**
 构建杂志字母频次表
 */
func getCount(s string) []int {
	res := make([]int, 26)
	for i := range s{
		res[s[i]-'a'] ++
	}
	return res
}
