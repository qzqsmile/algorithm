# import collections
#
# class Solution:
#     @staticmethod
#     def findLadders(beginWord, endWord, wordList):
#         wordList = set(wordList)
#         if endWord not in wordList:
#             return []
#         # BFS visit
#         curr_level = {beginWord}
#         parents = collections.defaultdict(list)
#         while curr_level:
#             wordList -= curr_level
#             next_level = set()
#             for word in curr_level:
#                 for i in range(len(word)):
#                     for c in 'abcdefghijklmnopqrstuvwxyz':
#                         new_word = word[:i] + c + word[i + 1:]
#                         if new_word in wordList:
#                             next_level.add(new_word)
#                             parents[new_word].append(word)
#             if endWord in next_level:
#                 break
#             curr_level = next_level
#         # DFS reconstruction
#         import pdb; pdb.set_trace()
#         res = []
#         def dfs(word, path):
#             if word == beginWord:
#                 path.append(word)
#                 res.append(path[::-1])
#             else:
#                 for p_word in parents[word]:
#                     dfs(p_word, path + [word])
#
#         dfs(endWord, [])
#         return res
#
#
# s = Solution()
# res = s.findLadders("hit", "cog", ["hot", "dot", "dog", "lot", "log", "cog"])
# print(res)
#
#
count = 1
def aka(res, path):
    pass
    if len(path) == 9:
        return
    else:
        global count
        path.append(count)
        count += 1
        res.append(path)
        aka(res, path+[1])

res = []
aka(res, [])
print(res)


