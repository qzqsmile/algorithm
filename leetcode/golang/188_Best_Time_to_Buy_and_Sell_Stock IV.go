package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	k = min(k, len(prices)/2)
	dp1 := make([]int, len(prices))
	dp2 := make([]int, len(prices))
	max_profit := 0
	for i := 1; i <= k; i++ {
		tmp := dp1[0] - prices[0]
		for j := 1; j < len(prices); j++ {
			tmp = max(dp1[j]-prices[j], tmp)
			dp2[j] = max(dp2[j-1], prices[j]+tmp)
			max_profit = max(max_profit, dp2[j])
		}
		copy(dp1, dp2)
	}
	return max_profit
}

func main() {

}
