class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        num = "123456789"
        for i in range(0, len(board)):
            for j in range(0, len(board)):
                if board[i][j] != '.':
                    flag = self.checkrow(board, i, j) and self.checkcol(board, i, j) and self.checkeach(board, i, j)
                    if not flag:
                        return False
        return True

    def checkrow(self, board, i, j):
        tmp, count = board[i][j], 0
        for each in board[i]:
            if each == tmp:
                count += 1
        return True if count == 1 else False

    def checkcol(self, board, i, j):
        tmp, count = board[i][j], 0
        for r in range(0, 9):
            if tmp == board[r][j]:
                count += 1
        return True if count == 1 else False

    def checkeach(self, board, i, j):
        tmp, count = board[i][j], 0
        r, c = i // 3, j // 3
        for r1 in range(r * 3, (r + 1) * 3):
            for c1 in range(c * 3, (c + 1) * 3):
                if tmp == board[r1][c1]:
                    count += 1
        return True if count == 1 else False



