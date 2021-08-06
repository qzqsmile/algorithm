class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        s1, s2 = version1.split('.'), version2.split('.')
        l, i = min(len(s1), len(s2)), 0
        while i < l:
            if int(s1[i]) < int(s2[i]):
                return -1
            elif int(s1[i]) > int(s2[i]):
                return 1
            i += 1
            
        if len(s1) > len(s2):
            while i < len(s1):
                if int(s1[i]) != 0:
                    return 1
                i += 1
        elif len(s1) < len(s2):
            while i < len(s2):
                if int(s2[i]) != 0:
                    return -1
                i += 1
        return 0