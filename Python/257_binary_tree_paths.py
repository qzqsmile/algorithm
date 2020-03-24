# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def binaryTreePaths(self, root: TreeNode) -> List[str]:
        if not root:
            return []
        if not root.left and not root.right:
            return [str(root.val)]
        l = self.binaryTreePaths(root.left)
        r = self.binaryTreePaths(root.right)
        res = []
        for i in l:
            res.append(str(root.val)+"->"+i)
        for i in r:
            res.append(str(root.val)+"->"+i)
        return res
        
        