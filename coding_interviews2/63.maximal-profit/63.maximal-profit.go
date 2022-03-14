package problem63

// 剑指 Offer 63. 股票的最大利润

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxDiff := max(prices[1]-prices[0], 0)
	minPrice := min(prices[0], prices[1])
	for i := 2; i < len(prices); i++ {
		if prices[i]-minPrice > maxDiff {
			maxDiff = prices[i] - minPrice
		}
		minPrice = min(minPrice, prices[i])
	}
	return maxDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
