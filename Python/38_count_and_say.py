class Solution:
    def countAndSay(self, n: int) -> str:
        res = "1"
        while n > 1:
            i, j, tmp = 0, 0, ""
            while i < len(res):
                j = i + 1
                count = 1
                while j < len(res):
                    if res[j] == res[i]:
                        count += 1
                    else:
                        break
                    j += 1
                tmp += str(res[i]) + str(count)
                i = j
            res = tmp
            n -= 1
        return res

s = Solution()
res = s.countAndSay(3)
print(res)
