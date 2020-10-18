class Solution:
    """
    @param strs: A list of strings
    @return: A list of strings
    """
    def anagrams(self, strs):
        # write your code here
        from collections import defaultdict
        res = []
        mp = defaultdict(lambda: 0)
        for each in strs:
            mp[''.join(sorted(each))] += 1
        for each in strs:
            if mp[''.join(sorted(each))] > 1:
                res.append(each)
        return res