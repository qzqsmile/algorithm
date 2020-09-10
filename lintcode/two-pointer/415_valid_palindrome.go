package main

import (
	"strings"
)

func checkStringAlphabet(str uint8) bool {
	if  '0' <= str && str <= '9' ||
		'A' <= str && str <= 'Z'{
		return true
	}
	return false
}

func isPalindrome (s string) bool {
	// write your code here
	i, j := 0, len(s)-1
	s = strings.ToUpper(s)
	for; i < j; {
		for ; i< j && !checkStringAlphabet(s[i]); i++{}
		for ; i<j && !checkStringAlphabet(s[j]); j--{}
		if i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}