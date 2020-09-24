class Solution:

    """
    @param: chars: The letter array you should sort by Case
    @return: nothing
    """

def sortLetters(self, chars):
    # write your code here
    i, j = 0, len(chars) - 1
    while i < j:
        while i < j and chars[i].islower():
        i += 1
    while i < j and not chars[j].islower():
        j -= 1

    if i < j:
        chars[i], chars[j] = chars[j], chars[i]
i += 1
j -= 1
