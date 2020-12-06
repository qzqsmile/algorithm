package BFS

func ladderLength(beginWord string, endWord string, wordList []string) int {
	q := []string{}
	q = append(q, beginWord)
	step := 1
	mp := make(map[string]bool)
	str := "qwertyuiopasdfghjklzxcvbnm"
	for _, v := range wordList{
		mp[v] = true
	}

	for len(q)>0 {
		q1 := []string{}
		for len(q) >0 {
			n := q[0]
			q = q[1:]
			if n == endWord{
				return step
			}
			for i := 0; i < len(n); i ++{
				for j := 0; j < len(str); j++{
					nn := n[0:i]+string(str[j])+n[i+1:]
					if _, ok := mp[nn]; ok{
						q1 = append(q1, nn)
						delete(mp, nn)
					}
				}
			}
		}
		step++
		q = q1
	}
	return 0
}
