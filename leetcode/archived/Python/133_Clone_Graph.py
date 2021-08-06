"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = []):
        self.val = val
        self.neighbors = neighbors
"""

# BFS

class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        from queue import Queue
        if not node:
            return node
        q = Queue()
        mp, visited = {}, set()
        mp[node.val] = Node(node.val)
        q.put(node)
        while not q.empty():
            cur = q.get()
            if cur.val in visited:
                continue
            for each in cur.neighbors:
                if each.val not in visited:
                    q.put(each)
                if each.val not in mp:
                    mp[each.val] = Node(each.val)
                mp[cur.val].neighbors.append(mp[each.val])
            visited.add(cur.val)
        return mp[node.val]


