class Solution:
    def sumNumbers(self, root: TreeNode) -> int:
        return self.helper(0, root)

    def helper(self, num, root):
        if not root:
            return num
        left = self.helper(num * 10 + root.val, root.left)
        right = self.helper(num * 10 + root.val, root.right)
        if not root.right:
            return left
        elif not root.left:
            return right
        return left + right