package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy_num := prices[0]
	max_profit := 0
	for i := 1; i < len(prices); i++ {
		if (prices[i] < buy_num) {
			buy_num = prices[i]
		} else {
			max_profit = max(max_profit, prices[i]-buy_num)
		}
	}
	return max_profit
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
