package string

import "strconv"

func addStrings(num1 string, num2 string) string {
	res := ""
	i, j := len(num1)-1, len(num2)-1
	num := 0
	for;i >=0 && j >= 0;{
		num += int(num1[i]-'0' + num2[j]-'0')
		res = strconv.Itoa(num % 10) + res
		num /= 10
		i--
		j--
	}

	for;i >= 0; i--{
		num += int(num1[i]-'0')
		res = strconv.Itoa(num % 10) + res
		num /= 10
	}

	for;j >= 0; j--{
		num += int(num2[j]-'0')
		res = strconv.Itoa(num % 10) + res
		num /= 10
	}
	if num > 0{
		res = strconv.Itoa(num) + res
	}
	return res
}
