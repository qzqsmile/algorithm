class Solution:
    def plusOne(self, digits):
        carry = 1
        for i in range(len(digits)-1, -1, -1):
            t = digits[i] + carry
            digits[i] = t % 10
            carry = t // 10
        return [1] + digits if carry == 1 else digits


