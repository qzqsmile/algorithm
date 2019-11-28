class Solution:
    def maxProfit(self, prices):
        if not prices:
            return 0
        min_num, max_stock = prices[0], 0
        for i in range(1, len(prices)):
            if prices[i] < min_num:
                min_num = prices[i]
            else:
                max_stock = max(max_stock, prices[i] - min_num)

        return max_stock
