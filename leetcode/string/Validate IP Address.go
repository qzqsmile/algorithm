func validIPAddress(queryIP string) string {
    if containStr('.', queryIP) && IsIPv4(queryIP){
        return "IPv4"
    }else if containStr(':', queryIP) && IsIPv6(queryIP){
        return "IPv6"
    }
    return "Neither"
}
// "20EE:FGb8:85a3:0:0:8A2E:0370:7334"
func containStr(str1 byte, str2 string) bool{
    for i := 0; i < len(str2); i++{
        if str2[i] == str1{
            return true
        }
    }
    return false
}

func IsIPv4(str string) bool{
    strs := strings.Split(str, ".")
    if len(strs) != 4{
        return false
    }
    for i := 0; i < len(strs); i++{
        if !IsIPv4Part(strs[i]){
            return false
        }
    }
    return true
}

func IsIPv4Part(str string) bool{
    if len(str) > 1 && str[0] == '0'{
        return false
    }
    num, err := strconv.Atoi(str)
    if err != nil {
        return false
    }
    if num < 0 || num > 255{
        return false
    }
    return true
}
func IsIPv6Part(str string) bool{
    if len(str) < 1 || len(str) > 4{
        return false
    }
    for i := 0; i < len(str); i++{
        if ('a'<= str[i] && str[i] <= 'f') || 
        ('A' <= str[i] && str[i] <= 'F') ||
        ('0' <= str[i] && str[i] <= '9'){
            continue
        }else{
            return false
        }
    }
    return true
}

func IsIPv6(str string) bool{
    strs := strings.Split(str, ":")
    if len(strs) != 8{
        return false
    }
    for i := 0; i < len(strs); i++{
        if !IsIPv6Part(strs[i]){
            return false
        }
    }
    return true
}