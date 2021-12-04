func reverseWords(s string) string {
    str := strip(s)
    if len(str) == 0{
        return str
    }
    res := []byte{}
    for i := 0; i < len(str); i++{
        if i != 0 && str[i] == ' ' && str[i-1] == ' '{
            continue
        }
        res = append(res, str[i])
    }
    reverse(res)

    for i := 0; i < len(res); i++{
        for j := i; j <= len(res); j++{
            if j == len(res) || res[j] == ' '{
                reverse(res[i:j])
                i = j
                break
            }
        }
    }

    return string(res)
}

func reverse(str []byte){
    i, j := 0, len(str)-1
    for;i < j;{
        str[i], str[j] = str[j], str[i]
        i++
        j--
    }
}

func strip(s string) string{
    i := 0
    for; i < len(s); i++{
        if s[i] != ' '{
            break
        }
    }
    j := len(s)-1
    for; j >= 0; j--{
        if s[j] != ' '{
            break
        }
    }
    if i >= len(s) || j <0{
        return ""
    }
    return s[i:j+1]
}