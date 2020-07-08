"""
Definition of TreeNode:
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left, self.right = None, None
"""


class Solution:
    def serialize(self, root):
        # write your code here
        from queue import Queue
        res = []
        q1 = Queue()
        if not root:
            return "[]"
        q1.put(root)
        while not q1.empty():
            n = q1.get()
            if n is None:
                res.append('#')
            else:
                res.append(n.val)
                q1.put(n.left)
                q1.put(n.right)
        return str(res)

    def deserialize(self, data):
        # write your code here
        res = eval(data)
        if not res:
            return None
        from queue import Queue
        q = Queue()
        root = TreeNode(res[0])
        q.put(root)
        i = 1
        while not q.empty():
            n = q.get()
            if i < len(res):
                if res[i] != '#':
                    n.left = TreeNode(res[i])
                    q.put(n.left)
                if res[i+1] != '#':
                    n.right = TreeNode(res[i+1])
                    q.put(n.right)
                i += 2
        return root
