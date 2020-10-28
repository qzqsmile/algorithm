"""
Definition for singly-linked list with a random pointer.
class RandomListNode:
    def __init__(self, x):
        self.label = x
        self.next = None
        self.random = None
"""


class Solution:
    # @param head: A RandomListNode
    # @return: A RandomListNode
    def copyRandomList(self, head):
        # write your code here
        if not head:
            return head
        cur, nxt = head, head.next
        while cur:
            copy = RandomListNode(cur.label)
            copy.next = nxt
            cur.next = copy
            cur = nxt
            if nxt:
                nxt = nxt.next

        cur, copy = head, head.next
        while cur:
            if cur.random:
                copy.random = cur.random.next
            cur = copy.next
            if cur:
                copy = cur.next

        cur, copy = head, head.next
        copy_head = copy
        while cur:
            cur.next = cur.next.next
            cur = cur.next
            if copy.next:
                copy.next = copy.next.next
                copy = copy.next

        return copy_head