class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        for i in range(0, len(board)):
            for j in range(0, len(board[0])):
                if board[i][j] == word[0] and self.search(board, i, j, word, 0):
                    return True
        return False

    def search(self, board, i, j, word, k):
        if k == len(word):
            return True
        if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]) or board[i][j] != word[k]:
            return False
        tmp, board[i][j] = board[i][j], '.'
        flag = self.search(board, i + 1, j, word, k + 1) or self.search(board, i - 1, j, word, k + 1) or self.search(
            board, i, j + 1, word, k + 1) or self.search(board, i, j - 1, word, k + 1)
        board[i][j] = tmp
        return flag

