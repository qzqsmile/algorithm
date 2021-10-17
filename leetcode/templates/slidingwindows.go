package templates

func lengthOfLongestSubstring(s string) int {
    left, right := 0, 0 
    windows := make(map[uint8]int)

    for;right < len(s);{
        windows[s[right]]++
        right++
        for;shrink condition;{
            windows[s[left]]--
            left++
        }
    }
}