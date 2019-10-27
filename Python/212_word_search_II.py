class TrieNode:
    def __init__(self):
        self.children = collections.defaultdict(TrieNode)
        self.is_word = False

class Trie:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        current = self.root
        for e in word:
            current = current.children[e]
        current.is_word = True


class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        tire = Trie()
        root = tire.root
        res = []
        for each in words:
            tire.insert(each)
        for i in range(len(board)):
            for j in range(len(board[0])):
                self.dfs(i, j, board, root, "", res)
        return res

    def dfs(self, i, j, board, node, path, res):
        if node.is_word:
            res.append(path)
            node.is_word = False
        if i<0 or i >= len(board) or j >= len(board[0]) or j < 0:
            return
        tmp = board[i][j]
        node = node.children.get(tmp)
        if not node:
            return
        board[i][j] = '#'
        self.dfs(i+1, j, board, node, path+tmp, res)
        self.dfs(i-1, j, board, node, path+tmp, res)
        self.dfs(i, j-1, board, node, path+tmp, res)
        self.dfs(i, j+1, board, node, path+tmp, res)
        board[i][j] = tmp

