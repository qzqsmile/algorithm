# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        carry = 0
        new_nodes = []
        while (l1 is not None and l2 is not None):
            res = (l1.val + l2.val) + carry
            carry = res // 10
            new_nodes.append(ListNode(res % 10))
            l1 = l1.next
            l2 = l2.next

        while (l1 is not None):
            res = l1.val + carry
            carry = res // 10
            l1 = l1.next
            new_nodes.append(ListNode(res % 10))

        while (l2 is not None):
            res = l2.val + carry
            carry = res // 10
            l2 = l2.next
            new_nodes.append(ListNode(res % 10))

        if carry != 0:
            new_nodes.append(ListNode(carry))

        for i in range(len(new_nodes)):
            if i < (len(new_nodes) - 1):
                new_nodes[i].next = new_nodes[i + 1]
        if len(new_nodes) > 0:
            return new_nodes[0]
        return None


#anther solution
 # def addTwoNumbers(self, l1, l2):
 #        def toint(node):
 #            return node.val + 10 * toint(node.next) if node else 0
 #        def tolist(n):
 #            node = ListNode(n % 10)
 #            if n > 9:
 #                node.next = tolist(n / 10)
 #            return node
 #        return tolist(toint(l1) + toint(l2))