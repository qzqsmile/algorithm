class Solution:
    def wordBreak(self, s: str, wordDict):
        res = []
        self.helper(res, "", s, set(wordDict))
        return res

    def helper(self, res, sen, s, words):
        if not s and sen:
            res.append(sen)
        else:
            for i in range(0, len(s)):
                if words and s[0:i + 1] in words:
                    tmp_sen = s[0:i + 1] if not sen else sen + " " + s[0:i + 1]
                    self.helper(res, tmp_sen, s[i + 1:], words)

s = Solution()
res = s.wordBreak("pineapplepenapple", ["apple", "pen", "applepen", "pine", "pineapple"])
print(res)

