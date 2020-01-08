func maxProfit(prices []int) int {
    b, e, maxprofit := 0, 1, 0
    for ; e < len(prices); {
        if (prices[e] > prices[b]){
            maxprofit += prices[e] - prices[b]
        }
        b, e = e, e+1
    }
    return maxprofit
}