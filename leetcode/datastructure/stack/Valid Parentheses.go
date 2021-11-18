func isValid(s string) bool {
    stk := []byte{}
    for i := 0; i < len(s); i++{
        if s[i] == '(' || s[i] == '[' ||
        s[i] == '{'{
            stk = append(stk, s[i])
        }else{
            if len(stk) == 0{
                return false
            }
            if s[i] == ')' && stk[len(stk)-1] == '('{
                stk = stk[0:len(stk)-1]
            }else if s[i] == ']' && stk[len(stk)-1] == '['{
                stk = stk[0:len(stk)-1]
            }else if s[i] == '}' && stk[len(stk)-1] == '{'{
                stk = stk[0:len(stk)-1]
            }else{
                return false
            }
        }
    }
    if len(stk) > 0{
        return false
    }
    return true
}