class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def recoverTree(self, root) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        cur = root
        pre = first = second = None
        s = []
        while cur:
            if not cur.left:
                if not pre:
                    pre = cur
                if cur.val < pre.val and not first:
                    first = pre
                if first and cur.val > first.val:
                    second = pre
                    break
                pre = cur
                s.append(cur.val)
                cur = cur.right
            else:
                prev = cur.left
                while prev.right and prev.right != cur:
                    prev = prev.right
                if not prev.right:
                    prev.right = cur
                    cur = cur.left
                else:
                    prev.right = None
                    if not pre:
                        pre = cur
                    if cur.val < pre.val and not first:
                        first = pre
                    if first and cur.val > first.val:
                        second = pre
                        break
                    pre = cur
                    s.append(cur.val)
                    cur = cur.right
        if not second:
            second = pre
        first.val, second.val = second.val, first.val


root = TreeNode(68)
root.left = TreeNode(41)
root.left.left = TreeNode(-85)
root.left.left.left = TreeNode(-73)
root.left.left.right = TreeNode(-49)
root.left.left.left.left = TreeNode(-98)
root.left.left.left.left.left = TreeNode(-124)

s = Solution()
res = s.recoverTree(root)

