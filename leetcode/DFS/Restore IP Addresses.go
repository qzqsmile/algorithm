import "fmt"

func restoreIpAddresses(s string) []string {
    res := [][]string{}
    tmp := []string{}
    dfs(s, tmp, &res)

    ans := []string{}
    for _, each := range res{
        tmp := ""
        for i := 0; i < len(each); i++{
            tmp += each[i]
            if i != len(each)-1{
               tmp += "."
            }
        }
        ans = append(ans, tmp)
    }
    return ans
}

func dfs(s string, tmp []string, res *[][]string){
    if len(tmp) > 4{
        return 
    }
    fmt.Println(tmp)
    if len(tmp) == 4 && len(s) == 0{
        *res = append(*res, append([]string{}, tmp...))
    }
    for i := 1; i <= len(s) && i <= 3; i++{
        if isValidIp(s[0:i]) && i <= len(s){
            tmp = append(tmp, s[0:i])
            dfs(s[i:], tmp, res)
            tmp = tmp[0:len(tmp)-1]
        }
    }
}

func isValidIp(s string) bool{
    if len(s) > 3 || len(s) == 0 || (s[0] == '0' && len(s) > 1){
        return false
    }

    num := 0
    for i := 0; i < len(s); i++{
        num = 10 * num + int(s[i]-'0')
    }

    if num >= 0 && num <= 255{
        return true
    }
    return false
}