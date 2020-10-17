class ZigzagIterator:
    """
    @param: v1: A 1d vector
    @param: v2: A 1d vector
    """
    def __init__(self, v1, v2):
        # do intialization if necessary
        self.v1 = v1
        self.v2 = v2
        self.c = 1

    """
    @return: An integer
    """
    def next(self):
        # write your code here
        h = self.c // 2
        v = 0
        if len(self.v1) >= h  and len(self.v2) >= h:
            l = self.c - 2 * h
            if l == 1:
                v = self.v2[h] if len(self.v1) == h else self.v1[h]
            else:
                v = self.v2[h-1]
        elif len(self.v1) < h:
            v = self.v2[self.c-len(self.v1)-1]
        else:
            v = self.v1[self.c-len(self.v2)-1]
        self.c += 1
        return v

    """
    @return: True if has next
    """
    def hasNext(self):
        # write your code here
        if len(self.v1) + len(self.v2) >= self.c:
            return True
        return False