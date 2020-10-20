package main

// 这题的key points在于堆栈的比较，不仅仅是出现的此处，还设有字符串的长度

func topKFrequent(words []string, k int) []string {
	mp := make(map[string] int)
	for _, v := range words{
		if _, ok := mp[v]; ok{
			mp[v]++
		}else{
			mp[v] = 1
		}
	}
	wordsL := []wordFrequent{}
	for key, v := range mp{
		if k > 0 {
			k--
			word := wordFrequent{key, v}
			wordsL = append(wordsL, word)
		}else{
			word := wordFrequent{key, v}
			if cmp(word, wordsL[0]){
				wordsL[0] = word
				heapSort(wordsL, 0)
			}
		}
		if k == 0 {
			k--
			makeHeap(wordsL)
		}
	}
	res := []string{}
	for i := len(wordsL)-1; i >= 0; i--{
		res = append(res, wordsL[0].word)
		wordsL[0] = wordsL[i]
		wordsL = wordsL[0:i]
		heapSort(wordsL, 0)
	}
	reverse(res)
	return res
}

func reverse(s []string){
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

type wordFrequent struct{
	word string
	frequent int
}

func makeHeap(words []wordFrequent){
	for i := len(words) / 2; i >= 0; i--{
		heapSort(words, i)
	}
}

func heapSort(words []wordFrequent,  i int){
	l := 2 * i + 1
	r := 2 * i + 2
	min := i
	if l < len(words) && cmp(words[min], words[l]){
		min = l
	}
	if r < len(words) && cmp(words[min], words[r]){
		min = r
	}
	if i != min{
		words[min], words[i] = words[i], words[min]
		heapSort(words, min)
	}
}

func cmp(s1 wordFrequent, s2 wordFrequent) bool{
	if s1.frequent == s2.frequent{
		return s1.word < s2.word
	}
	return s1.frequent > s2.frequent
}