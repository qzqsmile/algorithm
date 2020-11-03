class Solution:
    """
    @param n: An integer
    @return: An integer
    """
    def __init__(self):
        self.mp = {}
        self.mp[1] = 1
        self.mp[2] = 2
        self.mp[0] = 0

    def climbStairs(self, n):
        # write your code here
        if n not in self.mp:
            s = self.climbStairs(n-1)+self.climbStairs(n-2)
            self.mp[n] = s
            return s
        else:
            return self.mp[n]
