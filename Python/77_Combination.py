class Solution:
    def combine(self, n: int, k: int):
        def helper(b, e, k):
            res = []
            if e - b < k or k <= 0:
                return [[]]
            if e - b == k:
                return [list(range(b, e))]
            s1 = helper(b + 1, e, k)
            s2 = helper(b + 1, e, k - 1)
            for each in s2:
                res.append([b]+each)
            res += s1
            return res

        res = helper(1, n + 1, k)
        return res



