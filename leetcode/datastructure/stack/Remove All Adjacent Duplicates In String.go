func removeDuplicates(s string) string {
    stk := []byte{}
    for i := 0; i < len(s); i++{
        if len(stk) == 0 || s[i] != stk[len(stk)-1]{
            stk = append(stk, s[i])
        }else{
            stk = stk[0:len(stk)-1]
        }
    }

    res := ""
    for i := 0; i < len(stk); i++{
        res += string(stk[i])
    }
    return res
}