package main

import "fmt"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var mp map[string][]string
	var words, next_step_words []string
	var words_mp map[string]bool
	var res [][]string

	words_mp = make(map[string]bool)
	mp = make(map[string][]string)
	for i := 0; i < len(wordList); i++{
		words_mp[wordList[i]] = true
	}

	words = append(words, beginWord)

	for ; len(words) > 0; {
		var flag bool
		flag = false
		for i := 0; i < len(words); i++{
			var tmp string
			w := words[i]
			for j := 0; j < len(w); j++{
				for k := 'a'; k <= 'z'; k++{
					tmp = w[0:j]+string(k)+w[j+1:]
					if tmp == w{
						continue
					}
					_, ok := words_mp[tmp]
					if ok{
						if tmp == endWord{
							flag = true
						}
						mp[w] = append(mp[w], tmp)
						next_step_words = append(next_step_words, tmp)
					}
				}
			}
		}
		if flag{
			break
		}
		words = next_step_words
		next_step_words = next_step_words[:0]
		for i := 0; i < len(words); i++ {
			delete(words_mp, words[i])
		}
	}
	fmt.Println(mp)
	return res
}

func dfs(word string, path []string, b string, e string){
	if word == b{
		path = append(path, word)
	}
}

func main(){
	//c := 'a'+1
	//fmt.Printf(string(c))
	beginWord := "hit"
	endWord := "cog"
	wordList := []string {"hot","dot","dog","lot","log","cog"}
	findLadders(beginWord, endWord, wordList)
}

