class Solution:
    def candy(self, ratings: List[int]) -> int:
        res, pre, t, i, j = 1, 1, 0, 0, 1
        while j < len(ratings):
            if ratings[j] > ratings[j - 1]:
                pre += 1
                res += pre
            elif ratings[j] == ratings[j - 1]:
                t, i, pre = pre, j, 1
                res += 1
            else:
                if pre > 1:
                    t, i, pre = pre, j, 1
                    res += 1
                elif pre == 1:
                    res += j - i + 1
                    if i > 0 and ratings[i - 1] != ratings[i] and j - i + 1 >= t:
                        res += 1
            j += 1
        return res



