class Solution:
    def searchMatrix(self, matrix, target):
        if not matrix:
            return False
        row = 0
        while row + 1 < len(matrix):
            if matrix[row + 1][0] <= target:
                row += 1
            else:
                break

        b, e = 0, len(matrix[0]) - 1
        cur = (b + e) // 2
        while b <= e:
            if matrix[row][cur] == target:
                return True
            elif matrix[row][cur] > target:
                e = cur - 1
            else:
                b = cur + 1
            cur = (b + e) // 2
        return False
