# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseBetween(self, head, m: int, n: int):
        if m == 1:
            new_head, new_tail = self.helper(head, n - m)
            return new_head
        cur, m1 = head, m
        while m1 > 2:
            cur = cur.next
            m1 -= 1
        new_head, new_tail = self.helper(cur.next, n - m)
        cur.next = new_head
        return head

    def helper(self, head, c):
        if c == 0:
            return head, head
        c -= 1
        sub_head, sub_tail = self.helper(head.next, c)
        head.next = sub_tail.next
        sub_tail.next = head
        return sub_head, sub_tail.next