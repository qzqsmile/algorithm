class Vector2D(object):

    # @param vec2d {List[List[int]]}
    def __init__(self, vec2d):
        # Initialize your data structure here
        self.x = 0
        self.y = 0
        self.vec = vec2d


    # @return {int} a next element
    def next(self):
        # Write your code here
        v = self.vec[self.x][self.y]
        self.y += 1
        if self.y >= len(self.vec[self.x]):
            self.x += 1
            self.y = 0
        return v


    # @return {boolean} true if it has next element
    # or false
    def hasNext(self):
        # Write your code here
        while self.x < len(self.vec) and len(self.vec[self.x]) == 0:
            self.x += 1
        if self.x < len(self.vec) and self.y < len(self.vec[self.x]):
            return True
        return False


# Your Vector2D object will be instantiated and called as such:
# i, v = Vector2D(vec2d), []
# while i.hasNext(): v.append(i.next())