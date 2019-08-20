func countAndSay(n int) string {
	res := "1"
	for n > 1 {
		var i, count int = 0, 1
		var tmp string = ""
		for i < len(res){
			j := i + 1
			count = 1
			for j  <  len(res){
				if res[i] == res[j] {
					count += 1
				}else {
					break
				}
				j += 1
			}
			tmp += strconv.Itoa(count) +string(res[i])
			i = j
		}
		res = tmp
		n -= 1
	}
	return res
}