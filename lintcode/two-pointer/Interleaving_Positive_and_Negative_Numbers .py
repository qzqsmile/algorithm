class Solution:
    """
    @param: A: An integer array.
    @return: nothing
    """
    def rerange(self, A):
        # write your code here
        i, j = 0, len(A)-1
        while i < j:
            while i < j and A[i] < 0:
                i += 1
            while i < j and A[j] > 0:
                j -= 1
            if  i < j:
                A[i], A[j] = A[j], A[i]
                i += 1
                j -= 1
        i, p, n = 0, 0, 0
        while i < len(A):
            if A[i] > 0:
                p += 1
            else:
                n += 1
            i += 1

        if p > n:
            i, j = 0, len(A)-2
        elif n > p:
            i, j = 1, len(A)-1
        else:
            i, j = 0, len(A)-1
        while i < j:
            A[i], A[j] = A[j], A[i]
            i += 2
            j -= 2