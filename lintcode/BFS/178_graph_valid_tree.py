class Solution:
    """
    @param n: An integer
    @param edges: a list of undirected edges
    @return: true if it's a valid tree, or false
    """
    def validTree(self, n, edges):
        # write your code here
        from collections import defaultdict
        from queue import Queue

        if len(edges) != n-1:
            return False
        mp = defaultdict(set)
        for n1 in edges:
            mp[n1[0]].add(n1[1])
            mp[n1[1]].add(n1[0])

        visited = 0
        q1 = Queue()
        h = set()
        q1.put(0)
        h.add(0)
        while not q1.empty():
            visited += 1
            n1 = q1.get()
            for e in mp[n1]:
                if e in h:
                    continue
                q1.put(e)
                h.add(e)
        return n == visited
