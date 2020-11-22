package main

import "fmt"

func minWindow(s string, t string) string {
	if t == "" || s == ""{
		return ""
	}
	mp := make(map[uint8] int)
	count := 0
	res :=  s+"T"
	for i := 0; i < len(t); i++{
		if _, ok := mp[t[i]]; ok{
			mp[t[i]]++
		}else{
			mp[t[i]]=1
			count++
		}
	}

	for i, j := 0, 0; i < len(s); {
		for ;j < len(s);j++{
			if _, ok := mp[s[j]]; ok{
				mp[s[j]]--
				if mp[s[j]] == 0{
					count--
					if count == 0{
						j++
						break
					}
				}
			}
		}
		for ;i < len(s); i++{
			if _, ok := mp[s[i]]; ok{
				mp[s[i]]++
				if mp[s[i]] == 1 && count == 0{
					count++
					if j-i < len(res){
						res = s[i:j]
					}
					i++
					break
				}
			}
		}
	}
	if res == s+"T"{
		return ""
	}
	return res
}


func main(){
	s := "ADOBECODEBANC"
	t1 := "ABC"
	//wanted := "BANC"
	r := minWindow(s, t1)
	fmt.Print(r)
}