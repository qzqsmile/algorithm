class Solution:
    """
    @param: A: Given an integer array
    @return: nothing
    """
    def heapify(self, A):
        # write your code here
        for i in range(len(A)//2, -1, -1):
            self.heapSort(A, i)


    def heapSort(self, A, i):
        l = 2 * i + 1
        r = 2 * i + 2
        min = i
        if l < len(A) and A[l] < A[min]:
            min = l
        if r < len(A) and A[r] < A[min]:
            min = r
        if min != i:
            A[i], A[min] = A[min], A[i]
            self.heapSort(A, min)
