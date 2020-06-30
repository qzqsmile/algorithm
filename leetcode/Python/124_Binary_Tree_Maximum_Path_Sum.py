class Solution:
    def maxPathSum(self, root):
        self.maxvalue = float('-inf')
        def dfs(root):
            if not root: return 0
            left = max(0, dfs(root.left))
            right = max(0, dfs(root.right))
            self.maxvalue = max(self.maxvalue, left + right + root.val)
            return max(left,right)+root.val
        dfs(root)
        return self.maxvalue