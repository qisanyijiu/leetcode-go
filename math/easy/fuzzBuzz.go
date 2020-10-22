package easy

import "strconv"

/**
写一个程序，输出从 1 到 n 数字的字符串表示。

1. 如果 n 是3的倍数，输出“Fizz”；

2. 如果 n 是5的倍数，输出“Buzz”；

3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

 */
func fizzBuzz(n int) []string {
	var rst []string
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			rst = append(rst, "FizzBuzz")
		} else if i%5 == 0 {
			rst = append(rst, "Buzz")
		} else if i%3 == 0 {
			rst = append(rst, "Fizz")
		} else {
			rst = append(rst, strconv.Itoa(i))
		}
	}
	return rst
}