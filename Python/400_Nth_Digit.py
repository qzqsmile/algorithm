class Solution:
    def findNthDigit(self, n: int) -> int:
        if n < 10:
            return n
        lower, up, c = 0, 0, 0
        while up < n:
            lower = up
            up = up + 9 * pow(10, c) * (c + 1)
            c += 1
        t = (n - lower) // c + pow(10, c - 1)
        m = (n - lower) % c
        return int(str(t - 1)[-1]) if m == 0 else int(str(t)[m - 1])

