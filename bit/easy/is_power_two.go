package easy

/**
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

去掉最右边的1之后是否为0

a = 4
bit a     : 1 0 0
bit a - 1 : 0 1 1
a & a - 1 : 0 0 0
 */
func IsPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n & (-n) == 0
}
