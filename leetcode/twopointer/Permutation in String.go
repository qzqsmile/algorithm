package twopointer

func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	check := false
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for _, v := range s1{
		m1[byte(v)]++
	}
	for right < len(s2){
		c := s2[right]
		right++
		m2[c]++
		if right-left == len(s1) && isSub(m1, m2){
			return true
		}
		for right-left > len(s1){
			c := s2[left]
			left++
			m2[c]--
			if m2[c] == 0{
				delete(m2, c)
			}
			if right-left == len(s1) && isSub(m1, m2){
				return true
			}
		}
	}
	return check
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
