# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if not head:
            return head
        l, cur = 0, head
        while cur:
            cur = cur.next
            l += 1
        k %= l
        if k == 0:
            return head
        h = l - k
        newhead, last, cur, h = None, None, head, l-k
        while cur and cur.next:
            if h == 1:
                last = cur
                newhead = cur.next
            cur = cur.next
            h -= 1
        last.next = None
        cur.next = head
        return newhead