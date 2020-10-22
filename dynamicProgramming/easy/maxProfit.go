package easy

func maxProfit(prices []int) int {
	minPrice := 2147483647
	maxProfit := 0
	for _, val := range prices {
		if val < minPrice {
			minPrice = val
		} else if val-minPrice > maxProfit {
			maxProfit = val - minPrice
		}
	}
	return maxProfit
}