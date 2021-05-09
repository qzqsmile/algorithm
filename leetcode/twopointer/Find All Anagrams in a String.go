package twopointer

func findAnagrams(s string, p string) []int {
	res := []int{}
	left, right := 0, 0
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for _, v := range p{
		m1[byte(v)]++
	}
	for right < len(s){
		c := s[right]
		right++
		m2[c]++
		if right-left == len(p) && isSub(m1, m2){
			res = append(res, left)
		}
		for right-left > len(p){
			c := s[left]
			left++
			m2[c]--
			if m2[c] == 0{
				delete(m2, c)
			}
			if right-left == len(p) && isSub(m1, m2){
				res = append(res, left)
			}
		}
	}
	return res
}

