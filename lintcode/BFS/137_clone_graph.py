"""
class UndirectedGraphNode:
     def __init__(self, x):
         self.label = x
         self.neighbors = []
"""

class Solution:
    """
    @param node: A undirected graph node
    @return: A undirected graph node
    """
    def cloneGraph(self, node):
        # write your code here
        new = None
        from queue import Queue

        q = Queue()
        h = set()
        mp = {}
        if not node:
            return new

        q.put(node)
        # h.add(node)
        mp[node.label] = UndirectedGraphNode(node.label)
        new = mp[node.label]
        while not q.empty():
            n = q.get()
            for n1 in n.neighbors:
                if n1.label in mp:
                    continue
                mp[n1.label] = UndirectedGraphNode(n1.label)
                q.put(n1)

        q.put(node)
        h.add(node.label)
        while not q.empty():
            n = q.get()
            c = mp[n.label]
            for n1 in n.neighbors:
                c.neighbors.append(mp[n1.label])
                if n1.label not in h:
                    q.put(n1)
                    h.add(n1.label)

        return new



