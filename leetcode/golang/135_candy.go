func candy(ratings []int) int {
	var res, pre, t, i, j  = 1, 1, 0, 0, 1
	for j < len(ratings){
		if ratings[j] > ratings[j-1]{
				pre += 1
				res += pre
			}else if ratings[j] == ratings[j-1]{
				t = pre
				i = j
				pre = 1
				res += 1
			}else{
				if pre > 1{
				t = pre
				i = j
				pre = 1
				res += 1
			}else if pre == 1{
				res += j - i + 1
				if i > 0 && ratings[i-1] != ratings[i] && j-i+1 >= t{
				res += 1
				}
			}
		}
		j += 1
	}
	return res
}
