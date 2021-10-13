package twopointer

// 同向指针
func reverseWords(s string) string {
    b, e := 0, 0
    res := ""
    for;e < len(s);{
        //locate e
        for ;e < len(s) && !isspace(s[e]); e++{}
        tmp := reverse(s, b, e)
        res += tmp
        if e == len(s){
            break
        }else{
            res += " "
        }
        //locate b
        b = e + 1
        e++
    }
    return res
}

func reverse(s string, b int, e int)string{
    res := ""
    for i := e-1; i >= b; i--{
        res += string(s[i])
    }
    return res
}

func isspace(s uint8) bool{
    if s == ' '{
        return true
    }
    return false
}