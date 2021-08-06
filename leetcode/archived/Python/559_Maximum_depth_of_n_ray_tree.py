"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""
class Solution:
    def maxDepth(self, root: 'Node') -> int:
        if not root:
            return 0
        max_depth = 1
        for each in root.children:
            max_depth = max(max_depth, 1+self.maxDepth(each))
        return max_depth

        