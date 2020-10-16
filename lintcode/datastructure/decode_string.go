package main

import (
	"strconv"
)

func decodeString(s string) string {
	strStk := []string{}
	numStk := []int{}
	tmpStr := ""
	tmpNum := 0
	for i := 0; i < len(s); i++{
		if '0' <= s[i] && s[i] <= '9'{
			t1, _ :=  strconv.Atoi(string(s[i]))
			tmpNum = 10 * tmpNum + t1
		}else if 'a' <= s[i] && s[i] <= 'z'{
			tmpStr += string(s[i])
		}else if s[i] == '['{
			strStk = append(strStk, tmpStr)
			strStk = append(strStk, string(s[i]))
			numStk = append(numStk, tmpNum)
			tmpNum = 0
			tmpStr = ""
		}else if s[i] == ']'{
			if tmpStr != ""{
				strStk = append(strStk, tmpStr)
				tmpStr = ""
			}
			c := 1
			tStr := ""
			if len(numStk) > 0{
				c = numStk[len(numStk)-1]
				numStk = numStk[0:len(numStk)-1]
			}
			for len(strStk) > 0{
				t1 := strStk[len(strStk)-1]
				strStk = strStk[0:len(strStk)-1]
				if t1 == "["{
					break
				}
				tStr = t1 + tStr
			}
			if tStr != ""{
				t2 := tStr
				for c > 1{
					tStr += t2
					c--
				}
				strStk = append(strStk, tStr)
			}
		}
	}
	res := ""
	if tmpStr != ""{
		strStk = append(strStk, tmpStr)
	}
	for len(strStk) > 0{
		res = strStk[len(strStk)-1] + res
		strStk = strStk[0:len(strStk)-1]
	}

	return res
}
