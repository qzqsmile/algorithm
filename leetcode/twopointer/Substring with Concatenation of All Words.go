package twopointer

func findSubstring(s string, words []string) []int {
	res := []int{}
	total_len := len(words)*len(words[0])
	left, right := 0, 0
	m2 :=  make(map[string]int)
	for _, v := range words{
		m2[v]++
	}
	for right < len(s){
		right++

		if right-left == total_len{
			m1 := make(map[string]int)
			for i := 0; i < len(words); i++{
				m1[s[left+len(words[0])*i:left+len(words[0])*(i+1)]]++
			}
			if isSub1(m1, m2){
				res = append(res, left)
			}
		}

		for right-left > total_len{
			left++
			if right-left == total_len{
				m1 := make(map[string]int)
				for i := 0; i < len(words); i++{
					m1[s[left+len(words[0])*i:left+len(words[0])*(i+1)]]++
				}
				if isSub1(m1, m2){
					res = append(res, left)
				}
			}
		}
	}
	return res
}

func isSub1(m1 map[string]int, m2 map[string] int) bool{
	if len(m1) != len(m2){
		return false
	}
	for k, _ := range m1{
		if m1[k] != m2[k]{
			return false
		}
	}
	return true
}