class Solution:
    """
    @param A: A list of integers
    @return: An integer
    """
    def jump(self, A):
        # write your code here
        t1 = len(A)-1
        c = 0
        b, e = 0, 0
        while e < t1:
            mb, me = b+1, b+A[b]
            while b <= e:
                if me < b+A[b]:
                    me = b+A[b]
                    mb = b+1
                b += 1
            b = mb
            e = me
            c += 1
        return c