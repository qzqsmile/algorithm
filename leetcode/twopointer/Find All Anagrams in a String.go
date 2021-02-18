package main

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

func isSub(m1 map[byte]int, m2 map[byte] int) bool{
	if len(m1) != len(m2){
		return false
	}
	for k, _ := range m1{
		if m1[byte(k)] != m2[byte(k)]{
			return false
		}
	}
	return true
}
