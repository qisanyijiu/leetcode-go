package medium

import (
	"strconv"
)

/**
93. 复原IP地址

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。
 */
const SEG_COUNT = 4

func RestoreIpAddresses(s string) []string {
	var (
		segments = make([]int, SEG_COUNT)
		ans = []string{}
	)
	IPDfs(s, 0, 0, &segments ,&ans)
	return ans
}

func IPDfs(s string, segId, segStart int, segments *[]int,ans *[]string) {
	// 如果找到了 4 段 IP 地址并且遍历完了字符串，那么就是一种答案
	if segId == SEG_COUNT {
		if segStart == len(s) {
			ipAddr := ""
			for i := 0; i < SEG_COUNT; i++ {
				ipAddr += strconv.Itoa((*segments)[i])
				if i != SEG_COUNT - 1 {
					ipAddr += "."
				}
			}
			*ans = append(*ans, ipAddr)
		}
		return
	}

	// 如果还没有找到 4 段 IP 地址就已经遍历完了字符串，那么提前回溯
	if segStart == len(s) {
		return
	}
	// 由于不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
	if s[segStart] == '0' {
		(*segments)[segId] = 0
		IPDfs(s, segId + 1, segStart + 1, segments, ans)
	}
	// 一般情况，枚举每一种可能性并递归
	addr := 0
	for segEnd := segStart; segEnd < len(s); segEnd++ {
		addr = addr * 10 + int(s[segEnd] - '0')
		if addr > 0 && addr <= 0xFF {
			(*segments)[segId] = addr
			IPDfs(s, segId + 1, segEnd + 1, segments, ans)
		} else {
			break
		}
	}
}