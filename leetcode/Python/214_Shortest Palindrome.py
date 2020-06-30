class Solution:
    def shortestPalindrome(self, s: str) -> str:
        l = len(s)
        minstr = s + s
        for i in range(0, l):
            if i <= l // 2:
                if s[0:i] == "" or s[0:i] == s[i + 1:2 * i + 1][::-1]:
                    tmp = s[i + 1:l][::-1] + s[i] + s[i + 1:l]
                    if len(tmp) < len(minstr):
                        minstr = tmp
                if i != l // 2 and s[0:i + 1] == s[i + 1:2 * i + 2][::-1]:
                    tmp = s[i + 1:l][::-1] + s[i + 1:l]
                    if len(tmp) < len(minstr):
                        minstr = tmp
        return minstr

