class Solution:
    """
    @param: k: An integer
    """
    def __init__(self, k):
        # do intialization if necessary
        self.k = k
        self.A = []

    """
    @param: num: Number to be added
    @return: nothing
    """
    def add(self, num):
        # write your code here
        if len(self.A) == self.k:
            if num > self.A[0]:
                self.A[0] = num
                self.heapSort(self.A, 0)
        else:
            self.A.append(num)
            if len(self.A) == self.k:
                self.heapify(self.A)

    """
    @return: Top k element
    """
    def topk(self):
        # write your code here
        res = []
        copy = self.A[:]
        if len(copy) < self.k:
            return sorted(copy)[::-1]
        while len(copy) > 0:
            res.append(copy[0])
            copy[0] = copy[-1]
            copy = copy[0:len(copy)-1]
            self.heapSort(copy, 0)
        return res[::-1]


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
