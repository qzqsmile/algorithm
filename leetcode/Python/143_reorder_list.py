# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reorderList(self, head: ListNode) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        if head is None or head.next is None:
            return head
        count = 0
        cur = head
        while cur is not None:
            count += 1
            cur = cur.next
        new_count = count // 2 + count % 2
        cur = head
        while cur is not None and new_count > 1:
            cur = cur.next
            new_count -= 1
        h2 = cur.next
        cur.next = None

        n_h2, n_t2 = self.reverse(h2)
        new_head = self.merge(head, n_h2)
        return new_head
        
    
    def reverse(self, head):
        if head is None or head.next is None:
            return head, head
        h, t = self.reverse(head.next)
        t.next = head
        head.next = None
        return h, head

    def merge(self, l1, l2):
        if l1.next is None and l2.next is None:
            l1.next = l2
            return l1
        elif l2.next is None:
            l2.next = l1.next
            l1.next = l2
            return l1
        n_l = self.merge(l1.next, l2.next)
        l1.next = l2
        l2.next = n_l
        return l1
                
    
    
        