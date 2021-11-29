func longestPalindrome(s string) string {
    res := ""
    for i := 0; i < len(s); i++{
        tmp1 := palindrome1(s, i)
        tmp2 := palindrome2(s, i)
        res = max(res, max(tmp1, tmp2))
    }
    return res
}

func max(a string, b string) string{
    if len(a) > len(b){
        return a
    }
    return b
}

func palindrome1(s string, mid int) string{
    i, j := mid, mid 
    for ;i >= 0 && j < len(s);{
        if s[i] != s[j]{
            break
        }
        i--
        j++
    }
    return s[i+1:j]
}

func palindrome2(s string, mid int) string{
    i, j := mid, mid+1
    for ;i >= 0 && j < len(s);{
        if s[i] != s[j]{
            break
        }
        i--
        j++
    }
    return s[i+1:j]
}