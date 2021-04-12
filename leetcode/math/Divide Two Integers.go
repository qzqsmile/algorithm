package math

func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1{
		return 2147483647
	}

	flag := 1
	if divisor ^ dividend < 0{
		flag = -1
	}
	up := abs(dividend)
	down := abs(divisor)
	res := 0
	for i := 31; i >= 0; i--{
		if ((up >> i) >= down){
			res += 1 << i
			up -= down << i
		}
	}

	if flag > 0{
		return res
	}
	return -res
}

func abs(num int) int{
	if num < 0{
		return -num
	}
	return num
}
