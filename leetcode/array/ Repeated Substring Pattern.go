func repeatedSubstringPattern(s string) bool {
    count := len(s)
    for i := 0; i < len(s)/2; i++{
        tmp := s[0:i+1]
        wordlen := i+1
        if count % wordlen == 0{
            meet := true
            for j := wordlen; j + wordlen <= len(s); {
                if tmp != s[j:j+wordlen]{
                    meet = false
                    break
                }
                j += wordlen
            }
            if meet{
                return true
            }
        }
    }
    return false
}