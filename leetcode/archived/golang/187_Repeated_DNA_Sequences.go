func findRepeatedDnaSequences(s string) []string {
    record := make(map[string]bool)
    repeated := make(map[string]bool)
    for i := 0; i < len(s)-9; i++{
        sub := s[i:i+10]
        if _, ok := record[sub]; ok{
            repeated[sub] = true
        }else{
            record[sub] = true
        }
    }
    res := make([]string, 0)
    for k := range repeated{
        res = append(res, k)
    }
    return res
}