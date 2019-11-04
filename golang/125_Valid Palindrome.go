package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for ;(i < len(s) && j >= 0); {
		if (isNumberOrAlpha(s[i])){
			if  (isNumberOrAlpha(s[j])){
				if (strings.ToLower(string(s[i])) == strings.ToLower(string(s[j]))){
					i++; j--
				}else{
					return false
				}
			}else{
				j--
			}
		}else{
			i++
		}
	}
	return true
}

func isNumberOrAlpha(s uint8) bool{
	if ('a' <= s && s <= 'z') || ('A' <= s && s <= 'Z') || ('0' <= s && s <='9'){
		return true
	}else{
		return false
	}
}

func main()  {
	fmt.Println(strings.ToLower(string('1')))
	fmt.Println(isPalindrome("0P"))
}