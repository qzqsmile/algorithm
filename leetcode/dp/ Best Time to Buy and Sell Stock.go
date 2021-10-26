package dp

func maxProfit(prices []int) int {
	lowest := prices[0]
	maxPro := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > lowest {
			maxPro = max(prices[i]-lowest, maxPro)
		} else {
			lowest = prices[i]
		}
	}

	return maxPro
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
