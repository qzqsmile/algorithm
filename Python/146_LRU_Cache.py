#key point: make sure every pre and next point right node

class LinkNode:
    def __init__(self, key, value):
        self.key = key
        self.val = value
        self.pre = None
        self.next = None


class LRUCache:
    def __init__(self, capacity: int):
        self.l = capacity
        self.m = {}
        self.b = LinkNode(None, None)
        self.e = self.b

    def get(self, key: int) -> int:
        if key not in self.m or not self.m[key]:
            return -1
        if self.m[key].next:
            self.m[key].pre.next = self.m[key].next
            self.m[key].next.pre = self.m[key].pre
            self.e.next = self.m[key]
            self.m[key].pre = self.e
            self.e = self.e.next
            self.e.next = None
        return self.m[key].val

    def put(self, key: int, value: int) -> None:
        if key not in self.m or not self.m[key]:
            if self.l > 0:
                self.e.next = LinkNode(key, value)
                self.e.next.pre = self.e
                self.e = self.e.next
                self.m[key] = self.e
                self.l -= 1
            else:
                self.b = self.b.next
                self.m[self.b.key] = None
                self.e.next = LinkNode(key, value)
                self.e.next.pre = self.e
                self.e = self.e.next
                self.m[key] = self.e
        else:
            if self.m[key].next:
                self.m[key].next.pre = self.m[key].pre
                self.m[key].pre.next = self.m[key].next
                self.e.next = self.m[key]
                self.e.next.pre = self.e
                self.e = self.e.next
                self.e.next = None
            self.m[key].val = value


