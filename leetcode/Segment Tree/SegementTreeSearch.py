"""
Definition of SegmentTreeNode:
class SegmentTreeNode:
    def __init__(self, start, end, max):
        self.start, self.end, self.max = start, end, max
        self.left, self.right = None, None
"""

class Solution:
    """
    @param root: The root of segment tree.
    @param start: start value.
    @param end: end value.
    @return: The maximum number in the interval [start, end]
    """
    def query(self, root, start, end):
        # write your code here
        if start == root.start and end == root.end:
            return root.max
        if root.left.end < start:
            return self.query(root.right, start, end)
        elif root.right.start > end:
            return self.query(root.left, start, end)
        else:
            return max(self.query(root.left, start, root.left.end),
                    self.query(root.right, root.right.start, end))