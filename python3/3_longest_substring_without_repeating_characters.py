class Solution:
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        first = second = 0
        longest = 0
        store = set()

        while (second < len(s)):
            if s[second] not in store:
                store.add(s[second])
            else:
                longest = max(second - first, longest)
                while (s[first] != s[second]):
                    store.remove(s[first])
                    first += 1
                first += 1
            second += 1
        longest = max(second - first, longest)

        return longest
