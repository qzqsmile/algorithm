func simplifyPath(path string) string {
    res := ""
    stk := []string{}
    pathStr := ""
    for i := 0; i < len(path)+1; i++{
        if i == len(path) || path[i] == '/'{
            if len(pathStr) == 1 && pathStr[0] == '.'{
                pathStr = ""
                continue
            }else if len(pathStr) == 2 && pathStr[0] == '.' && pathStr[1] == '.'{
                if len(stk) > 0{
                    stk = stk[0:len(stk)-1]
                }
            }else if len(pathStr) > 0{
                stk = append(stk, pathStr)
            }
            pathStr = ""
        }else if path[i] == '.'{
            pathStr += "."
        }else{
            pathStr += string(path[i])
        }
    }

    for i := 0; i < len(stk); i++{
        res +=  "/" + stk[i] 
    }
    if len(res) == 0{
        res = "/"
    }
    return res
}