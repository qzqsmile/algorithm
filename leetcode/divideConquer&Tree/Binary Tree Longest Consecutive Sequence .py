"""
Definition of TreeNode:
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left, self.right = None, None
"""

class Solution:
    """
    @param root: the root of binary tree
    @return: the length of the longest consecutive sequence path
    """
    def longestConsecutive(self, root):
        # write your code here
        self.max = 0
        self.helper(root)
        return self.max

    def helper(self, root):
        if root is None:
            return 0
        if root.left is None and root.right == None:
            self.max = max(self.max, 1)
            return 1
        l = self.helper(root.left)
        r = self.helper(root.right)
        rl, rr = 1, 1
        if root.left and root.val == root.left.val-1:
            rl = l + 1
            self.max = max(self.max, rl)
        if root.right and root.val == root.right.val-1:
            rr = r + 1
            self.max = max(self.max, rr)
        return max(rl, rr)