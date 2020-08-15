"""
Definition of TreeNode:
class TreeNode:
    def __init__(self, val):
        this.val = val
        this.left, this.right = None, None
Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
"""
class Solution:
    # @param {TreeNode} root the root of binary tree
    # @return {ListNode[]} a lists of linked list
    def binaryTreeToLists(self, root):
        # Write your code here
        from queue import Queue
        res = []
        q1 = Queue()
        if root is None:
            return res
        q1.put(root)
        while not q1.empty():
            r = []
            q2 = Queue()
            while not q1.empty():
                n = q1.get()
                r.append(n.val)
                if n.left is not None:
                    q2.put(n.left)
                if n.right is not None:
                    q2.put(n.right)
            if r:
                t = self.convert(r)
                res.append(t)
            q1 = q2
        return res

    def convert(self, l):
        head, pre = None, None
        for v in l:
            n = ListNode(v)
            if pre is not None:
                pre.next = n
                pre = n
            if head is None:
                head = pre = n
        return head
