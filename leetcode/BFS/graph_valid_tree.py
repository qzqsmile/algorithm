# 判断是否为Tree即满足以下两个条件
# 1. 全部可访问 2. 无环
# 这里的思路是在n-1条边的情形下，如果全部可访问必定无环。并且是随便选一个结点
# 深度遍历下去，如果全部可达必定无环。这里有两个技巧值的关注
# 1. 存图需要存双向的关系，因为最终结果只是观察是否全部可达，存双向也可以
# 2. 随机选一个结点即可，不用非要选入度为0的root结点。因为理论上一个节点都要消耗一条边
# 如果选择入度为0的结点，就算接下来全部可达。也仍然可能有circle。当然也可以没有circle
# 在某个子节点被两个父节点引用的情形下。

class Solution:
    """
    @param n: An integer
    @param edges: a list of undirected edges
    @return: true if it's a valid tree, or false
    """
    def validTree(self, n, edges):
        # write your code here
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
        # if 0 not in mp and :
        #     return False
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