class Solution:
    def isPalindrome(self, s: str) -> bool:
        i, j = 0, len(s)-1
        while i < len(s) and j >= 0:
            if s[i].isalpha() or s[i].isnumeric():
                if s[j].isalpha() or s[j].isnumeric():
                    if s[i].lower() == s[j].lower():
                        i += 1
                        j -= 1
                    else:
                        return False
                else:
                    j -= 1
            else:
                i += 1
        return True
