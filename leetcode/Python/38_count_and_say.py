class Solution:
    def candy(self, ratings):
        res, pre, t, i, j = 1, 1, 0, 0, 1
        while j < len(ratings):
            if ratings[j] > ratings[j - 1]:
                res += pre + 1
                pre = pre + 1
            elif ratings[j] == ratings[j - 1]:
                t = pre
                res += 1
                pre = 1
                i = j
            else:
                if pre > 1:
                    t = pre
                    res += 1
                    pre = 1
                    i = j
                elif pre == 1:
                    res += j - i + 1
                    if i > 0 and ratings[i - 1] != ratings[i] and j - i + 1 >= t:
                        res += 1
            j += 1
        return res

s = Solution()
res = s.candy([1,6,10,8,7,3,2])
print(res)
