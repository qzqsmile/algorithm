class Solution {
public:
    int minDistance(string word1, string word2) {
        int l1 = word1.size();
        int l2 = word2.size();
        
        vector<int> dp(l2+1, 0);
        for(int i = 1; i <= l2; i++)
            dp[i] = i;
        
        for(int i = 1; i <= l1; i++)
        {
            int prev = i;
            for(int j = 1; j <= l2; j++)
            {
                int cur;
                if(word1[i-1] == word2[j-1])
                {
                    cur = dp[j-1];
                }
                else
                {
                    cur = min(min(dp[j-1],prev),dp[j])+1;
                }
                dp[j-1] = prev;
                prev =cur;
            }
            dp[l2] = prev;
        }
        return dp[l2];
    }
};