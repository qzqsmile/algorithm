class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int low , max = 0;
        if(!prices.size())
            return 0;
        low = prices[0];
        for(int i = 0; i < prices.size(); i++)
        {
            if(prices[i] <= low)
                low = prices[i];
            else
                max = (prices[i] - low) > max ? (prices[i]-low) : max; 
        }
        
        return max;
    }
};