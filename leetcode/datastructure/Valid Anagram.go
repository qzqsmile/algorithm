func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    mp1 := make(map[byte]int)
    mp2 := make(map[byte]int)
    for i := 0; i < len(s); i++{
        mp1[s[i]]++
    }

    for i := 0; i < len(t); i++{
        mp2[t[i]]++
    }

    for k, _ := range mp1{
        if _, ok := mp2[k];ok{
            if mp1[k] != mp2[k]{
                return false
            }
        }else{
            return false
        }
    }

    return true
}