func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
        return ""
    }
    res := ""
    for pre := 0; pre < len(strs[0]); pre++{
        for w := 0; w < len(strs); w++{
            if pre >= len(strs[w]) || (w != 0 && strs[w][pre] != strs[w-1][pre]){
                return res
            }
        }
        res += string(strs[0][pre])
    }
    return res
}