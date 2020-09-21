class MyQueue:

    def __init__(self):
        # do intialization if necessary
        self.stk1 = []
        self.stk2 = []
    """
    @param: element: An integer
    @return: nothing
    """
    def push(self, element):
        # write your code here
        self.stk1.append(element)

    """
    @return: An integer
    """
    def pop(self):
        # write your code here
        while self.stk1:
            self.stk2.append(self.stk1[-1])
            self.stk1.pop()

        n = self.stk2[-1]
        self.stk2.pop()
        while self.stk2:
            self.stk1.append(self.stk2[-1])
            self.stk2.pop()
        return n
    """
    @return: An integer
    """
    def top(self):
        # write your code here
        while self.stk1:
            self.stk2.append(self.stk1[-1])
            self.stk1.pop()

        n = self.stk2[-1]
        while self.stk2:
            self.stk1.append(self.stk2[-1])
            self.stk2.pop()

        return n