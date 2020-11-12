#class SVNRepo:
#    @classmethod
#    def isBadVersion(cls, id)
#        # Run unit tests to check whether verison `id` is a bad version
#        # return true if unit tests passed else false.
# You can use SVNRepo.isBadVersion(10) to check whether version 10 is a 
# bad version.
class Solution:
    """
    @param n: An integer
    @return: An integer which is the first bad version.
    """
    def findFirstBadVersion(self, n):
        # write your code here
        b, e = 1, n
        while b + 1 < e:
            mid = b + (e-b) // 2
            if SVNRepo.isBadVersion(mid):
                e = mid
            else:
                b = mid
        if SVNRepo.isBadVersion(b):
            return b
        return e