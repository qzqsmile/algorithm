package slidingwindows

//注意minlength的初始化条件
func minWindow(s string, t string) string {
    left, right := 0, 0
    lookup := make(map[uint8] int)
    meetcount := 0
    minLength := ""
    

    for i := 0; i < len(t); i++{
        lookup[t[i]]++
    }

    for;right < len(s);{
        if _, ok := lookup[s[right]]; ok{
            lookup[s[right]]--
            if lookup[s[right]] == 0{
                meetcount++
            }
        }
        right++
        for;meetcount == len(lookup);{
            if _, ok := lookup[s[left]]; ok{
                lookup[s[left]]++
                if lookup[s[left]] > 0{
                    meetcount--
                }
            }
            if right-left < len(minLength) || minLength == ""{
                minLength = s[left: right]
            }
            left++
        }
    }

    return minLength
}