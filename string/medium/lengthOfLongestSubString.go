package medium

func lengthOfLongestSubString(s string) int {
	// 初始化位置表
	location := make([]int, 256)
	for i := range location{
		location[i] = -1
	}
	maxLen, left := 0, 0
	for i := 0; i <  len(s); i ++ {
		if location[s[i]] >= left{
			// 说明是第二次出现，应该移动左指针了
			left = location[s[i]] + 1
		}else if i + 1 - left > maxLen{
			// 字串长度大于最大长度，更新最大长度
			maxLen = i + 1 - left
		}
		// 更新位置表
		location[s[i]] = i
	}
	return maxLen
}