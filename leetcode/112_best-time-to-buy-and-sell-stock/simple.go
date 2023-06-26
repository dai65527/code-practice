func maxProfit(prices []int) int {
	min := 1000000000000
	profit := 0
	for _, price := range prices {
		if price < min {
			min = price
		}
		if profit < price-min {
			profit = price - min
		}
	}
	return profit
}
