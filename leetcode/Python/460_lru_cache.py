class LinkNode:
    def __init__(self, key, val, pre=None, nxt=None, ref=1):
        self.key = key
        self.val = val
        self.pre = pre
        self.next = nxt
        self.ref = ref


class T:
    def __init__(self):
        self.head = LinkNode(None, None)
        self.tail = self.head


from collections import defaultdict


class LFUCache:
    def __init__(self, capacity: int):
        self.r = defaultdict(T)
        self.m = {}
        self.c = capacity
        self.m_r = -1

    def get(self, key: int) -> int:
        if key not in self.m or self.m_r == -1:
            return -1
        if self.m_r == self.m[key].ref and self.r[self.m_r].tail.pre == self.r[self.m_r].head:
            self.m_r += 1
        self.m[key].ref += 1
        if self.m[key].next:
            self.m[key].next.pre = self.m[key].pre
            self.m[key].pre.next = self.m[key].next
        else:
            self.r[self.m[key].ref - 1].tail = self.r[self.m[key].ref - 1].tail.pre
            self.m[key].pre.next = None
            self.m[key].pre = None
        self.m[key].next = None
        self.r[self.m[key].ref].tail.next = self.m[key]
        self.m[key].pre = self.r[self.m[key].ref].tail
        self.r[self.m[key].ref].tail = self.r[self.m[key].ref].tail.next
        return self.m[key].val

    def put(self, key: int, value: int) -> None:
        if key in self.m:
            self.get(key)
            self.m[key].val = value
            return
        if self.c > 0:
            self.c -= 1
            self.m[key] = LinkNode(key, value)
            self.r[1].tail.next = self.m[key]
            self.m[key].pre = self.r[1].tail
            self.r[1].tail = self.r[1].tail.next
            self.m_r = 1
        else:
            if self.m_r == -1:
                return
            del self.m[self.r[self.m_r].head.next.key]
            self.m[key] = LinkNode(key, value)
            self.r[1].tail.next = self.m[key]
            self.m[key].pre = self.r[1].tail
            self.r[1].tail = self.r[1].tail.next
            self.r[self.m_r].head = self.r[self.m_r].head.next
            self.m_r = 1

# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)