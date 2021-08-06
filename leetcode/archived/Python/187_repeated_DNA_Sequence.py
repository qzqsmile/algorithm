class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        repeated, record = set(), set()
        for i in range(0, len(s) - 9):
            sub = s[i:i + 10]
            if sub in record:
                repeated.add(sub)
            else:
                record.add(sub)
        return list(repeated)
