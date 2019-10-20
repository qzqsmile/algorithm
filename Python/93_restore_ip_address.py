class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        res = []
        self.helper(s, 0, 0, res)
        return res

    def helper(self, s, i, c, res):
        if i >= len(s) or c > 3:
            return
        elif c == 3 and (int(s[i:]) <= 255 and (s[i] != '0' or len(s) - i == 1)):
            res.append(s)
        else:
            for j in range(i + 1, len(s)):
                if int(s[i:j]) <= 255 and (s[i] != '0' or j - i == 1):
                    self.helper(s[0:j] + "." + s[j:], j + 1, c + 1, res)
                else:
                    break
