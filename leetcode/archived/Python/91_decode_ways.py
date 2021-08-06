class Solution:
    def numDecodings(self, s: str) -> int:
        dp = [1, 1] if s and s[0] != '0' else [1, 0]
        for i in range(1, len(s)):
            if s[i] == '0' and (s[i - 1] == '0' or int(s[i - 1:i + 1]) > 26):
                dp.append(0)
            elif s[i - 1] == '0' or int(s[i - 1:i + 1]) > 26:
                dp.append(dp[-1])
            elif s[i] == '0':
                dp.append(dp[-2])
            else:
                dp.append(dp[-1] + dp[-2])
        return dp[-1]

