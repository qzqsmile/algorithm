class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        str_store = [str(c) for c in range(1, n + 1)]
        f, res = [1] * n, ""
        for i in range(1, n):
            f[i] = f[i - 1] * i
        k -= 1
        for i in range(n, 0, -1):
            j = k // f[i - 1]
            k %= f[i - 1]
            res += str_store[j]
            str_store.remove(str_store[j])
        return res
