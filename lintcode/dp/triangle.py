# traversal time limited O(n^2)
class Solution:
    """
    @param triangle: a list of lists of integers
    @return: An integer, minimum path sum
    """


    def minimumTotal(self, triangle):
        # write your code here
        import sys
        self.triangle = triangle
        self.best = sys.maxsize
        if not triangle:
            return 0
        self.helper(0, 0, 0)
        return self.best

    def helper(self, x, y, sum):
        if x == len(self.triangle):
            if sum < self.best:
                self.best = sum
            return
        if y < len(self.triangle[x]):
            self.helper(x+1,y, sum+self.triangle[x][y])
            self.helper(x+1,y+1, sum+self.triangle[x][y])


# divide and conquer

class Solution:
    """
    @param triangle: a list of lists of integers
    @return: An integer, minimum path sum
    """
    def minimumTotal(self, triangle):
        # write your code here
        self.triangle = triangle
        if not triangle:
            return 0
        return self.helper(0, 0)

    def helper(self, x, y):
        if x == len(self.triangle)-1:
            return self.triangle[x][y]
        return self.triangle[x][y] + min(self.helper(x+1,y), self.helper(x+1,y+1))
