class Solution:
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        word_list, next_words, step = set(wordList), [], 1
        next_words.append(beginWord)

        while len(next_words) > 0:
            new_to_do = []
            for word in next_words:
                if word == endWord:
                    return step
                for i in range(len(word)):
                    for s1 in "abcdefghijklmnopqrstuvwxyz":
                        if s1 == word[i]:
                            continue
                        else:
                            tmp = word[0:i] + s1 + word[i + 1:]
                            if tmp in word_list:
                                word_list.remove(tmp)
                                new_to_do.append(tmp)
            next_words = new_to_do
            step = step + 1
        return 0

