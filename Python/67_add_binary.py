#only works on int, will exceed this limit on other progmramming language

class Solution:
    def addBinary(self, a: str, b: str) -> str:
        def converttonumber(a):
            res, i = 0, 0
            while i < len(a):
                res = res * 2 + int(a[i])
                i += 1
            return res

        def convertnumbertostr(a):
            if a == 0:
                return "0"
            res = ""
            while a > 0:
                res += str(a % 2)
                a = a // 2
            return res[::-1]

        n1, n2 = converttonumber(a), converttonumber(b)
        return convertnumbertostr(n1 + n2)
