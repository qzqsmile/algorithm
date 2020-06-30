# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def __init__(self):
        self.mp = {}

    def rob(self, root: TreeNode) -> int:
        if not root:
            return 0
        if not root.left and not root.right:
            return root.val
        if root in self.mp:
            return self.mp[root]
        l = self.rob(root.left)
        r = self.rob(root.right)
        l1, l2, r1, r2 = 0, 0, 0, 0
        if root.left:
            l1 = self.rob(root.left.left)
            l2 = self.rob(root.left.right)
        if root.right:
            r1 = self.rob(root.right.left)
            r2 = self.rob(root.right.right)
        self.mp[root] = max(l + r, root.val + l1 + l2 + r1 + r2)
        return self.mp[root]






