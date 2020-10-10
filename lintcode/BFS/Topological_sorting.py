"""
Definition for a Directed graph node
class DirectedGraphNode:
    def __init__(self, x):
        self.label = x
        self.neighbors = []
"""


class Solution:
    """
    @param: graph: A list of Directed graph node
    @return: Any topological order for the given graph.
    """
    def topSort(self, graph):
        # write your code here
        # def dfs(indegree,g,ans):
        #     ans.append(g)
        #     indegree[g] = -1 
        #     for c in g.neighbors:
        #         indegree[c] -= 1
        #         if indegree[c] == 0:
        #             dfs(indegree,c,ans)

        # indegree = {}
        # ans = []
        # for g in graph:
        #     for n in g.neighbors:
        #         if n not in indegree:
        #             indegree[n] = 1
        #         else:
        #             indegree[n] += 1
        # for g in graph:
        #     if g not in indegree or indegree[g] == 0:
        #         dfs(indegree,g,ans)
        # return ans

        indegree = {}
        ans = []
        for g in graph:
            for n in g.neighbors:
                if n not in indegree:
                    indegree[n] = 1
                else:
                    indegree[n] += 1

        q = []
        for g in graph:
            if g not in indegree:
                q.append(g)

        while q:
            temp = q.pop(0)
            ans.append(temp)
            for n in temp.neighbors:
                indegree[n] -= 1
                if indegree[n] == 0:
                    q.append(n)
        return ans
        