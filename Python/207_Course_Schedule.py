class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        in_count = [0] * numCourses
        out_count = [[] for i in range(numCourses)]
        for each in prerequisites:
            out_count[each[1]].append(each[0])
            in_count[each[0]] += 1

        for i in range(0, numCourses):
            for node in range(len(in_count)):
                if in_count[node] == 0:
                    for v in out_count[node]:
                        in_count[v] -= 1
                    in_count[node] = -1
        for i in in_count:
            if i != 0 and i != -1:
                return False
        return True


#dfs solutions
# class Solution:
#     def canFinish(self, numCourses, prerequisites):
#         s1, s2 = set(), set()
#         out_count = [[] for i in range(numCourses)]
#         for each in prerequisites:
#             s1.add(each[1])
#             s2.add(each[0])
#             out_count[each[1]].append(each[0])
#         start = s1 - s2
#         if len(start) == 0 and len(prerequisites) != 0:
#             return False
#         for each in start:
#             visited = set()
#             if not self.dfs(visited, each, out_count):
#                 return False
#         return True
#
#     def dfs(self, visited, node, out_count):
#         if node in visited:
#             return False
#         visited.add(node)
#         if len(out_count[node]) > 0:
#             for each in out_count[node]:
#                 if not self.dfs(visited, each, out_count):
#                     return False
#         visited.remove(node)
#         return True



