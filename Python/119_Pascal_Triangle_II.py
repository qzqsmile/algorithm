class Solution:
    def getRow(self, rowIndex: int):
        res = [[1], [1,1]]
        if rowIndex <= 1:
            return res[rowIndex]
        lastest_row = res[-1]
        for i in range(2, rowIndex+1):
            lastest_row = list(map(lambda x, y: x + y, [0]+lastest_row, lastest_row+[0]))
        return lastest_row

