func titleToNumber(columnTitle string) int {
    ans := 0
    for i := 0; i < len(columnTitle); i++{
        ans = 26 * ans + int(columnTitle[i]-'A')+1
    }
    return ans
	sort.Slice()
}