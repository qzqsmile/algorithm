package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	carry, res, i := 0, "", len(a) - 1
	for  j := len(b)-1; (i >= 0 || j >= 0 || carry > 0); j = j-1{
		if i >= 0{
			carry += int(a[i] - '0')
		}
		if j >= 0{
			carry += int(b[j] - '0')
		}
		res, carry = strconv.Itoa(int(carry % 2)) + res, carry / 2
		i--
	}
	return res
}

