"""
Definition of TreeNode:
class TreeNode:
    def __init__(self, val):
        this.val = val
        this.left, this.right = None, None
"""


class Solution:
    """
    @param: root: The root of the binary tree.
    @param: A: A TreeNode
    @param: B: A TreeNode
    @return: Return the LCA of the two nodes.
    """
    def lowestCommonAncestor3(self, root, A, B):
        # write your code here
        bA, bB, res = self.helper(root, A, B)
        return res


    def helper(self, root, A, B):
        if root is None:
            return False, False, root
        lp, lq, lc = self.helper(root.left, A, B)
        rp, rq, rc = self.helper(root.right, A, B)
        if lc:
            return True, True, lc
        if rc:
            return True, True, rc
        bp = lp or rp
        bq = lq or rq
        if root == A:
            bp = True
        if root == B:
            bq = True
        if bp and bq:
            return True, True, root
        return bp, bq, None