func removeKdigits(num string, k int) string {
    stk := []int{}
    for i := 0; i < len(num); i++{
        for;k > 0 && len(stk) > 0 && num[i] < num[stk[len(stk)-1]];{
            stk = stk[0:len(stk)-1]
            k--
        }
        stk = append(stk, i)
    }

    ans := ""
    for i := 0; i < len(stk); i++{
        ans += string(num[stk[i]])
    }
    if k > 0 && len(ans) >= k{
        ans = ans[0:len(ans)-k]
    }
    i := 0
    for ;i < len(ans); i++{
        if ans[i] != '0'{
            break
        }
    }
    
    ans = ans[i:]
    if ans == ""{
        ans = "0"
    }

    return ans
}