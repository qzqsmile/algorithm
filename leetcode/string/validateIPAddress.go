package string


func validIPAddress(IP string) string {
	res := []int{-1}
	for i := 0; i < len(IP); i++{
		if IP[i] == '.' || IP[i] == ':'{
			res = append(res, i)
		}
	}
	res = append(res, len(IP))

	if len(res) == 5{
		for i := 0; i < len(res)-1; i++{
			if !isIPv4(IP[res[i]+1:res[i+1]]){
				return "Neither"
			}
		}
		return "IPv4"
	}else if len(res) == 9{
		for i := 0; i < len(res)-1; i++{
			if !isIPv6(IP[res[i]+1:res[i+1]]){
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}

func isIPv4(str string) bool{
	if len(str) < 1 || len(str) > 3 || (len(str) > 1 && str[0] == '0'){
		return false
	}
	num := 0
	for i := 0; i < len(str); i++{
		if '0' <= str[i] && str[i] <= '9'{
			num = 10 * num + int(str[i]-'0')
		}else{
			return false
		}
	}
	if num <= 255{
		return true
	}
	return false
}

func isIPv6(str string) bool{
	if len(str) <= 0 || len(str) > 4{
		return false
	}
	for i := 0; i < len(str); i++{
		if ('a' <= str[i] && str[i] <= 'f') ||
			('A' <= str[i] && str[i] <= 'F') ||
			('0' <= str[i] && str[i] <= '9'){
			continue
		}else{
			return false
		}
	}
	return true
}


