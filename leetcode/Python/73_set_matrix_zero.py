class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        if not matrix:
            return matrix
        fr, fc = False, False
        for i in range(0, len(matrix[0])):
            if matrix[0][i] == 0:
                fr = True
        for i in range(0, len(matrix)):
            if matrix[i][0] == 0:
                fc = True
        for i in range(0, len(matrix)):
            for j in range(0, len(matrix[0])):
                if matrix[i][j] == 0:
                    matrix[0][j] = 0
                    matrix[i][0] = 0

        for i in range(1, len(matrix)):
            for j in range(1, len(matrix[0])):
                if matrix[0][j] == 0 or matrix[i][0] == 0:
                    matrix[i][j] = 0
        if fr:
            for i in range(0, len(matrix[0])):
                matrix[0][i] = 0
        if fc:
            for i in range(0, len(matrix)):
                matrix[i][0] = 0
