class Solution:
    def minCut(self, s: str) -> int:
        dp = [i - 1 for i in range(0, len(s) + 1)]
        for i in range(0, len(s)):
            j = 0
            while i + j < len(s) and i - j >= 0 and s[i + j] == s[i - j]:
                dp[i + j + 1] = min(dp[i + j + 1], dp[i - j] + 1)
                j += 1

            j = 0
            while i + j < len(s) and i - j - 1 >= 0 and s[i + j] == s[i - j - 1]:
                dp[i + j + 1] = min(dp[i + j + 1], dp[i - j - 1] + 1)
                j += 1

        return dp[-1]

