"""
Definition of SegmentTreeNode:
class SegmentTreeNode:
    def __init__(self, start, end, max):
        self.start, self.end, self.max = start, end, max
        self.left, self.right = None, None
"""

class Solution:
    """
    @param A: a list of integer
    @return: The root of Segment Tree
    """
    def build(self, A):
        # write your code here
        if len(A) == 0:
            return None
        return self.buildTree(A, 0, len(A)-1)

    def buildTree(self, A, start, end):
        root = SegmentTreeNode(start, end, None)
        if start == end:
            root.max = A[start]
            return root
        root.left = self.buildTree(A, start, (start+end)//2)
        root.right = self.buildTree(A, (start+end)//2+1, end)
        root.max = max(root.left.max, root.right.max)
        return root