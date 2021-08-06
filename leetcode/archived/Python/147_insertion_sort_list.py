# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def insertionSortList(self, head: ListNode) -> ListNode:
        if head is None or head.next is None:
            return head
        new_head = self.insertionSortList(head.next)
        if new_head.val > head.val:
            head.next = new_head
            return head
        cur = new_head
        while cur is not None and cur.next is not None:
            if cur.next.val >= head.val:
                head.next = cur.next
                cur.next = head
                break
            cur = cur.next
        if cur.next is None:
            cur.next = head
            head.next = None
        return new_head
        
        
    