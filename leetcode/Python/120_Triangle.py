class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        if not triangle:
            return 0
        res = triangle[0]
        for i in range(1,len(triangle)):
            tmp = res[:]
            res = [10000000]
            for j in range(0, len(tmp)):
                res[-1] = min(tmp[j] + triangle[i][j], res[-1])
                res.append(tmp[j] + triangle[i][j+1])
        return min(res)
