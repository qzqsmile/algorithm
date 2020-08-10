class Solution:
    """
    @param A: An integers array.
    @return: return any of peek positions.
    """
def findPeak(self, A):
    # write your code here
    b, e = 1, len(A) - 2
    while b + 1 < e:
        m = (e - b) // 2 + b
    if A[m] > A[m - 1] and A[m] > A[m + 1]:
        return m
    elif A[m] > A[m - 1]:
        b = m
    else:
        e = m
    if A[b] > A[b - 1] and A[b] > A[b + 1]:
        return b
    else:
        return e
