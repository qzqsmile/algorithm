package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	q := []string{}
	visited := make(map[string] bool)
	step := 0
	q = append(q, beginWord)

	for len(q) > 0{
		nq := q
		q = []string{}
		for len(nq) > 0{
			t := nq[0]
			nq = nq[1:]
			if t == endWord{
				return step +1
			}
			for _, v := range wordList{
				if match(t, v){
					if _, ok := visited[v]; !ok{
						visited[v] = true
						q = append(q, v)
					}
				}
			}
		}
		if len(q) > 0{
			step++
		}
	}
	return 0
}

func match(str1 string, str2 string) bool{
	if len(str1) != len(str2){
		return false
	}
	d := 0
	for i := 0; i < len(str1); i++{
		if str1[i] != str2[i]{
			d++
		}
	}
	if d == 1{
		return true
	}
	return false
}
