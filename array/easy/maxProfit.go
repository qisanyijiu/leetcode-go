package easy

func maxProfit(prices []int) int {
	profit := 0
	if len(prices) == 1 {
		return profit
	}
	for i := 1; i< len(prices); i ++ {
		if prices[i] - prices[i-1] > 0 {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}