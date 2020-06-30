class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        b, e, max_profit = 0, 1, 0
        while e < len(prices):
            if prices[e]-prices[b] > 0:
                max_profit += prices[e]-prices[b]
            b, e = e, e+1
        return max_profit