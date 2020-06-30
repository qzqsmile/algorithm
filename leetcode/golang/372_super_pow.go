func power(a int, b int) int {
	res := 1
	for i := 1; i <= b; i++ {
		res *= a
		if res > 1337 {
			res %= 1337
		}
	}
	return res
}

func superPow(a int, b []int) int {
	res := 0
	for index, value := range b {
		if index == 0 {
			res = power(a, value) % 1337
		} else {
			res = (power(res, 10) % 1337) * (power(a, value) % 1337) % 1337
		}
	}
	return res
}

