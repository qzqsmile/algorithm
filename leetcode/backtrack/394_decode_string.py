class Solution:
    def decodeString(self, s: str) -> str:
        i, res = 0, ""

        while i < len(s):
            if s[i].isdigit():
                b = i
                while s[i].isdigit():
                    i += 1
                times = int(s[b:i])
                j, stk = i, []
                while j < len(s):
                    if s[j] == '[':
                        stk.append(j)
                    elif s[j] == ']':
                        stk.pop()
                    if not stk:
                        break
                    j += 1
                tmp = self.decodeString(s[i + 1:j])
                res = res + tmp * times
                i = j
            else:
                res += s[i]
            i += 1
        return res