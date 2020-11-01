"""
Definition of TreeNode:
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left, self.right = None, None
"""

class Solution:
    """
    @param root: root of a tree
    @return: head node of a doubly linked list
    """
    def treeToDoublyList(self, root):
        # Write your code here.
        head,tail = self.helper(root)
        head.left = tail
        tail.right = head
        return head


    def helper(self, root):
        if root is None:
            return root, root
        if root.left is None and root.right is None:
            return root, root

        lh, lt = self.helper(root.left)
        rh, rt = self.helper(root.right)
        h, t = lh, rt
        if lt is not None:
            lt.right = root
            root.left = lt
        else:
            h = root

        if rh is not None:
            root.right = rh
            rh.left = root
        else:
            t = root
        return h, t
