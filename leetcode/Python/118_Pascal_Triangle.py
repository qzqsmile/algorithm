class Solution:
    def generate(self, numRows):
        pascal = [[1]*(i+1) for i in range(0, numRows)]
        for i in range(2, numRows):
            for j in range(1,i):
                pascal[i][j] = pascal[i-1][j-1] + pascal[i-1][j]
        return pascal
