class Solution:
    def findAllConcatenatedWordsInADict(self, words: List[str]) -> List[str]:
        res, s = [], set(words)
        for w in words:
            if self.isConcatedWords(w, 0, len(w), s):
                res.append(w)
        return res

    def isConcatedWords(self, w, b, e, s):
        for i in range(b + 1, e + 1):
            if w[b:i] in s and ((i == e and b != 0) or self.isConcatedWords(w, i, e, s)):
                return True
        return False