func maxProfit(prices []int) int {
	var profit, high int
	low := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < low || i == 0 {
			low = prices[i]
			high = 0
			for j := i + 1; j < len(prices); j++ {
				if prices[j] > high {
					high = prices[j]
					profit = max(profit, (prices[j] - prices[i]))
				}
			}
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
