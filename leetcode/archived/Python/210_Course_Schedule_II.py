# class Solution:
#     def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
#         in_count = [0] * numCourses
#         out_count = [[] for i in range(numCourses)]
#         for each in prerequisites:
#             out_count[each[1]].append(each[0])
#             in_count[each[0]] += 1
#         order = []
#         for i in range(0, numCourses):
#             for node in range(len(in_count)):
#                 if in_count[node] == 0:
#                     order.append(node)
#                     for v in out_count[node]:
#                         in_count[v] -= 1
#                     in_count[node] = -1
#         for i in in_count:
#             if i != 0 and i != -1:
#                 return []
#         return order
#


class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        out_count = [[] for i in range(numCourses)]
        for each in prerequisites:
            out_count[each[1]].append(each[0])
        order = []
        visited = set()
        for each in range(0, numCourses):
            if (each not in order) and (not self.dfs(visited, each, out_count, order)):
                return []
        order = order[::-1]
        return order

    def dfs(self, visited, node, out_count, order):
        if node in visited:
            return False
        if node in order:
            return True
        visited.add(node)
        if len(out_count[node]) > 0:
            for each in out_count[node]:
                if not self.dfs(visited, each, out_count, order):
                    return False
        order.append(node)
        visited.remove(node)
        return True

pdb.set_trace()

from queue import Queue

q = Queue()

visited = set()
visited.add()