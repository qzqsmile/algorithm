# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

#mirror Traversal

class Solution:
    def recoverTree(self, root) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        cur = root
        p1, c1, p2 = None, None, None
        while cur:
            if not cur.left:
                if p2:
                    if cur.val < p2.val:
                        if not p1:
                            p1, c1 = p2, cur
                        else:
                            c1 = cur
                p2 = cur
                cur = cur.right
            else:
                cur1 = cur.left
                while cur1.right and cur1.right != cur:
                    cur1 = cur1.right
                if not cur1.right:
                    cur1.right = cur
                    cur = cur.left
                else:
                    cur1.right = None
                    if p2:
                        if cur.val < p2.val:
                            if not p1:
                                p1, c1 = p2, cur
                            else:
                                c1 = cur
                    p2 = cur
                    cur = cur.right
        p1.val, c1.val = c1.val, p1.val




