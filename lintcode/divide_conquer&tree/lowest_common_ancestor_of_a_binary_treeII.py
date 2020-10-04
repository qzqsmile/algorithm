"""
Definition of ParentTreeNode:
class ParentTreeNode:
    def __init__(self, val):
        self.val = val
        self.parent, self.left, self.right = None, None, None
"""


class Solution:
    """
    @param: root: The root of the tree
    @param: A: node in the tree
    @param: B: node in the tree
    @return: The lowest common ancestor of A and B
    """
    def lowestCommonAncestorII(self, root, A, B):
        # write your code here
        path1 = self.helper(A)
        path2 = self.helper(B)
        i = 1
        while i < len(path1) and i < len(path2) and path1[i] == path2[i]:
            i += 1
        return path1[i-1]

    def helper(self, root):
        path = []
        c = root
        while c is not None:
            path.append(c)
            c = c.parent
        path = path[::-1]
        return path